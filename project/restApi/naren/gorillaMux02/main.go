package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/articles/{category}/{id:[0-9]+}", articleFunc)
	r.HandleFunc("/articles/", queryHandler)
	r.Queries("id", "category")

	//Chaining
	//r.Path("/articles/{category}/{id:[0-9]+}").HandlerFunc(articleFunc)

	//Sub routing
	// s:=r.PathPrefix("/articles").Subrouter()
	// s.HandleFunc("/id/setting",settings)
	// s.HandleFunc("/id/details",details)

	//Reverse URL
	r.HandleFunc("/reverse/{category}/{id:[0-9]+}", articleFunc).Name("rUrl")
	url, _ := r.Get("rUrl").URL("category", "ac", "id", "2")
	fmt.Println("Reverse URL:", url)
	s := &http.Server{
		Addr:         "127.0.0.1:8070",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.ListenAndServe()
}

func articleFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Category is %v\n", vars["category"])
	fmt.Fprintf(w, "Id is %v\n", vars["id"])
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()
	fmt.Fprintf(w, "Category is %v\n", qp["category"][0])
	fmt.Fprintf(w, "ID is %v\n", qp.Get("id"))
}
