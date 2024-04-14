package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Proxy struct {
}

func (Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, _ := url.Parse("http://127.0.0.1:20023")
	reverseProxy := httputil.NewSingleHostReverseProxy(remote)
	reverseProxy.ServeHTTP(w, r)
}

func main() {
	addr := "127.0.0.1:8081"
	proxy := Proxy{}

	fmt.Printf("proxy server on %s\n", addr)
	http.ListenAndServe(addr, proxy)
}
