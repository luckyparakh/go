package main

import (
	"commitlog/internal/server"
	"log"
)

func main() {
	srv := server.NewHTTPServer(":7890")
	log.Println("Starting Server")
	log.Fatal(srv.ListenAndServe())
}
