package main

import (
	"time"
)

// messageは1つのメッセージを表します。
type message struct {
	Name    string
	Message string
	When    time.Time
}
