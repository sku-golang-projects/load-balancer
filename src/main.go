package main

import (
	"LoadBalancer/model"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting load balancer")

	server := []model.Server{
		model.CreateServer("https://www.google.com"),
		model.CreateServer("https://www.facebook.com"),
		model.CreateServer("https://www.yahoo.com"),
	}

	lb := model.NewLoadBalancer("8000", server)

	handlerFunction := func(rw http.ResponseWriter, r *http.Request) {
		lb.Serve(rw, r)
	}

	http.HandleFunc("/", handlerFunction)

	http.ListenAndServe(":8000", nil)
}
