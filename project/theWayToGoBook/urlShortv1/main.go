package main

import (
	"fmt"
	"net/http"
)
const PORT = "5678"
var store *UrlStore = NewUrlStore()
func main(){
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(":"+PORT, nil)
}

func Redirect(w http.ResponseWriter, r *http.Request){
	key:= r.URL.Path[1:]
	if url := store.Get(key);url == ""{
		http.NotFound(w,r)
		return
	}else{
		http.Redirect(w,r,url,http.StatusFound)
	}
}

func Add (w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "GET":
		http.ServeFile(w,r,"urlForm.html")
	case "POST":
		url:=r.FormValue("url")
		//fmt.Println("Url",url)
		if url==""{
			//fmt.Fprintf(w, "URL is empty.")
			http.ServeFile(w,r,"urlForm.html")
			return
		}
		key := store.Put(url)
		// fmt.Println("http://localhost:%s/%s",PORT,key)
		fmt.Fprintf(w, "http://localhost:%s/%s",PORT,key)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}