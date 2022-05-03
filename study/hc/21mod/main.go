package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello Mod")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serve).Methods("GET")
	http.ListenAndServe(":40000", r)
}

func greeter() {
	fmt.Println("Hello from mod")
}

func serve(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang</h1>"))
}
