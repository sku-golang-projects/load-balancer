package model

import (
	"log"
	"net/http/httputil"
	"net/url"
	"os"
)

type ApplicationServer struct {
	Addr  string
	Proxy *httputil.ReverseProxy
}

func CreateServer(addr string) *ApplicationServer {

	target, err := url.Parse(addr)

	handleError(err)

	server := ApplicationServer{
		Addr:  addr,
		Proxy: httputil.NewSingleHostReverseProxy(target),
	}

	return &server
}

func handleError(err error) {
	if err != nil {
		log.Printf("Error is:%v\n", err)
		os.Exit(1)
	}
}
