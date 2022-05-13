package main

import (
	"encoding/json"
	"fmt"
	cls "middleware/closures"
	"net/http"
)

type city struct {
	Name string
	Area uint32
}

func main() {
	numGenerator()
	fmt.Println("Fib Generator")
	fibGen := cls.FibClosure()
	for i := 1; i < 7; i++ {
		fmt.Println(fibGen())
	}
	cls.Timer(cls.Add)(1, 2)

	//Http Closure
	http.HandleFunc("/", loggingMiddleware(homeHandler))
	//http.HandleFunc("/city", cityHandler)
	http.HandleFunc("/city", loggingMiddleware(contentCheck(cityHandler)))
	http.ListenAndServe(":9000", nil)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Hello from Home Page")
}
func loggingMiddleware(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before Func")
		f(w, r)
		fmt.Println("After Func")
	}
}
func cityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tmpCity city
		if err := json.NewDecoder(r.Body).Decode(&tmpCity); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error while decoding"))
		}
		defer r.Body.Close()
		fmt.Println("City is:", tmpCity.Name)
		fmt.Println("Area is:", tmpCity.Area)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Created"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", "Wrong method type")
	}

}
func contentCheck(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%v", "Wrong Content Type")
			return
		}
		f(w, r)
	}
}
func numGenerator() {
	numg1 := cls.NumGen()
	for i := 0; i < 5; i++ {
		fmt.Println(numg1())
	}
	fmt.Println("Another Instance")
	numg2 := cls.NumGen()
	for i := 0; i < 2; i++ {
		fmt.Println(numg2())
	}
}
