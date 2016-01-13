package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var addr = flag.String("addr", ":8081", "Webサイトのアドレス")
	flag.Parse()
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/",
		http.FileServer(http.Dir("public"))))
	log.Println("Webサイトのアドレス:", *addr)
	http.ListenAndServe(*addr, mux)
}
