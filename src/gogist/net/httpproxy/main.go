package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type handle struct {
	reverseProxy string
}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse(h.reverseProxy)
	if err != nil {
		log.Fatalln(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	r.Host = remote.Host
	proxy.ServeHTTP(w, r)
	log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String() + " " + r.Proto + " " + r.UserAgent())
}

func main() {
	bind := flag.String("l", "0.0.0.0:5555", "listen on ip:port")
	remote := flag.String("r", "http://bing.com:80", "reverse proxy addr")
	flag.Parse()
	log.Printf("Listening on %s, forwarding to %s", *bind, *remote)
	h := &handle{reverseProxy: *remote}
	err := http.ListenAndServe(*bind, h)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
