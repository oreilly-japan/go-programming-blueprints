package main

type room struct {
	// forwardは他のクライアントに転送するためのメッセージを保持するチャネルです。
	forward chan []byte
}
