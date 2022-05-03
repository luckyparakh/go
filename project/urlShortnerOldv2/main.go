package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luckyparakh/urlShortner/constants"
	"github.com/luckyparakh/urlShortner/handlers"
)

func main() {
	fmt.Println("URL Shortner")
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.ListAllShortLinkHandler).Methods("GET")
	r.HandleFunc("/encode", handlers.CreateShortLinkHandler).Methods("POST")
	r.HandleFunc("/file", handlers.ListAllShortLinkFileHandler).Methods("GET")
	r.HandleFunc("/file/encode", handlers.CreateShortLinkFileHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+constants.Port, r))

}
