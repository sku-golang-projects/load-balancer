package model

import (
	"log"
	"net/http"
)

type LoadBalancer struct {
	Port            string
	Server          []Server
	RoundRobinCount int
}

func NewLoadBalancer(port string, server []Server) *LoadBalancer {
	lb := LoadBalancer{
		Port:            port,
		Server:          server,
		RoundRobinCount: 0,
	}

	return &lb
}

func (lb *LoadBalancer) GetNextAvailableServer() Server {
	server := lb.Server[lb.RoundRobinCount%len(lb.Server)]
	for !server.IsAlive() {
		lb.RoundRobinCount++
		server = lb.Server[lb.RoundRobinCount%len(lb.Server)]
	}
	lb.RoundRobinCount++

	return server
}

func (lb *LoadBalancer) Serve(rw http.ResponseWriter, r *http.Request) {
	targetServer := lb.GetNextAvailableServer()
	log.Printf("Running the server %s", targetServer.Address())
	targetServer.Serve(rw, r)
}
