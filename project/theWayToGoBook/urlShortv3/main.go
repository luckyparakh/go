package main

import (
	"fmt"
	"net/http"
	"sync"
	"flag"
	
)

var wg sync.WaitGroup
//Command line Argument
var (
	Port = flag.String("port", "5678", "Port on which app listen")
	FileName = flag.String("file", "data.gob", "File where to save data")
	HostName = flag.String("host", "localhost", "Hostname")
)
var store *UrlStore
func main(){
	flag.Parse()
	store = NewUrlStore(*FileName)
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(":"+*Port, nil)
	wg.Wait()
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
		key:=store.Put(url)
		// fmt.Println("http://localhost:%s/%s",PORT,key)
		fmt.Fprintf(w, "http://%s:%s/%s",*HostName,*Port,key)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}