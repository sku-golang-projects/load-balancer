package model

import "net/http"

type Server interface {
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
	Address() string
}

func (aserver *ApplicationServer) Address() string {
	return aserver.Addr
}

func (aserver *ApplicationServer) IsAlive() bool {
	return true
}

func (aserver *ApplicationServer) Serve(rw http.ResponseWriter, r *http.Request) {
	aserver.Proxy.ServeHTTP(rw, r)
}
