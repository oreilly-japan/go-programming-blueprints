package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/matryer/filedb"
	"github.com/oreilly-japan/go-programming-blueprints/appendixB/backup"
)

type path struct {
	Path string
	Hash string
}

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			log.Fatalln(fatalErr)
		}
	}()
	var (
		interval = flag.Duration("interval", 10*time.Second, "チェックの間隔")
		archive  = flag.String("archive", "archive", "アーカイブの保存先")
		dbpath   = flag.String("db", "./db", "filedbデータベースへのパス")
	)
	flag.Parse()

	m := &backup.Monitor{
		Destination: *archive,
		Archiver:    backup.ZIP,
		Paths:       make(map[string]string),
	}

	db, err := filedb.Dial(*dbpath)
	if err != nil {
		fatalErr = err
		return
	}
	defer db.Close()
	col, err := db.C("paths")
	if err != nil {
		fatalErr = err
		return
	}

	var path path
	col.ForEach(func(_ int, data []byte) bool {
		if err := json.Unmarshal(data, &path); err != nil {
			fatalErr = err
			return true
		}
		m.Paths[path.Path] = path.Hash
		return false // 処理を続行します
	})
	if fatalErr != nil {
		return
	}
	if len(m.Paths) < 1 {
		fatalErr = errors.New("パスがありません。backupツールを使って追加してください")
		return
	}
	check(m, col)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
Loop:
	for {
		select {
		case <-time.After(*interval):
			check(m, col)
		case <-signalChan:
			// 終了
			fmt.Println()
			log.Printf("終了します...")
			break Loop
		}
	}

}

func check(m *backup.Monitor, col *filedb.C) {
	log.Println("チェックします...")
	counter, err := m.Now()
	if err != nil {
		log.Panicln("バックアップに失敗しました:", err)
	}
	if counter > 0 {
		log.Printf(" %d個のディレクトリをアーカイブしました\n", counter)
		// ハッシュ値を更新します
		var path path
		col.SelectEach(func(_ int, data []byte) (bool, []byte, bool) {
			if err := json.Unmarshal(data, &path); err != nil {
				log.Println("JSONデータの読み込みに失敗しました。"+
					"次の項目に進みます:", err)
				return true, data, false
			}
			path.Hash, _ = m.Paths[path.Path]
			newdata, err := json.Marshal(&path)
			if err != nil {
				log.Println("JSONデータの書き出しに失敗しました。"+
					"次の項目に進みます:", err)
				return true, data, false
			}
			return true, newdata, false
		})
	} else {
		log.Println(" 変更はありません")
	}
}
