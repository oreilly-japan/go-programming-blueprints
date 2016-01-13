package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/matryer/filedb"
)

type path struct {
	Path string
	Hash string
}

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			flag.PrintDefaults()
			log.Fatalln(fatalErr)
		}
	}()
	var (
		dbpath = flag.String("db", "./backupdata", "データベースのディレクトリへのパス")
	)
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fatalErr = errors.New("エラー; コマンドを指定してください")
		return
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

	switch strings.ToLower(args[0]) {
	case "list":
		var path path
		col.ForEach(func(i int, data []byte) bool {
			err := json.Unmarshal(data, &path)
			if err != nil {
				fatalErr = err
				return true
			}
			fmt.Printf("= %s\n", path)
			return false
		})

	case "add":
	case "remove":
	}

}
