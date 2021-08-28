package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	client := http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				// Skip TLS dial
				return net.Dial(network, addr)
			},
		},
	}

	r, err := client.Get("http://localhost:8080")
	perr(err)
	b, err := ioutil.ReadAll(r.Body)
	perr(err)
	log.Printf("body: %s", b)
}

func perr(err error) {
	if err != nil {
		log.Panicf("err: %s", err)
	}
}
