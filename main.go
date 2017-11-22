package main

import (
	"net/http/httputil"
	"net/url"
	"flag"
	"net/http"
	"fmt"
	"log"
)

func main() {
	port := flag.Int("port", 3000, "local port to run on")
	host := flag.String("host", "localhost:8000", "target host")
	scheme := flag.String("scheme", "http", "scheme (http/https)")
	flag.Parse()
	log.Fatal(run(*host, *scheme, *port))
}

func run(host, scheme string, port int) error {
	u := &url.URL{}
	u.Host = host
	u.Scheme = scheme
	p := httputil.NewSingleHostReverseProxy(u)
	log.Printf("listening on :%v", port)
	return http.ListenAndServe(fmt.Sprintf(":%v", port), p)
}
