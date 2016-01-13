package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/context"

	"github.com/bitly/go-nsq"

	"gopkg.in/mgo.v2"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	go func() {
		<-signalChan
		cancel()
		log.Println("停止します...")
	}()
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	if err := dialdb(); err != nil {
		log.Fatalln("MongoDBへのダイヤルに失敗しました:", err)
	}
	defer closedb()

	// 処理を開始します
	votes := make(chan string) // 投票結果のためのチャネル
	go twitterStream(ctx, votes)
	publishVotes(votes)
}

var db *mgo.Session

func dialdb() error {
	var err error
	log.Println("MongoDBにダイヤル中: localhost")
	db, err = mgo.Dial("localhost")
	return err
}
func closedb() {
	db.Close()
	log.Println("データベース接続が閉じられました")
}

type poll struct {
	Options []string
}

func loadOptions() ([]string, error) {
	var options []string
	iter := db.DB("ballots").C("polls").Find(nil).Iter()
	var p poll
	for iter.Next(&p) {
		options = append(options, p.Options...)
	}
	iter.Close()
	return options, iter.Err()
}

func publishVotes(votes <-chan string) {
	pub, _ := nsq.NewProducer("localhost:4150", nsq.NewConfig())
	for vote := range votes {
		pub.Publish("votes", []byte(vote)) // 投票内容をパブリッシュします
	}
	log.Println("Publisher: 停止中です")
	pub.Stop()
	log.Println("Publisher: 停止しました")
}
