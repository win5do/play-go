package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("world"))
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// http://127.0.0.1:8000/debug/pprof/
