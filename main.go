package main

import (
	"github.com/231tr0n/graphql-example/routes"
	"log"
	"net/http"
)

func main() {
	var server = http.Server{
		Handler:        &routes.Mux,
		Addr:           ":8080",
		MaxHeaderBytes: 50 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
