package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)
var wg sync.WaitGroup
var mut sync.Mutex
var signals = []string{"test"}
func main(){
	//caller1()
	caller2()
}

func caller2(){
	websiteList:=[]string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
	}

	for _ , website := range websiteList{
		
		go getStatusCode(website)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(website string){
	defer wg.Done()
	res,err := http.Get(website)
	if err!=nil{
		log.Fatal("Error")
	}else{
		mut.Lock()
		//Resource used by many go routines
		signals = append(signals, website)
		mut.Unlock()
		fmt.Printf("%d status code of %s\n", res.StatusCode,website)
	}
}
func caller1(){
	go greeting("hello") // Go routine
	greeting("world")
}

func greeting(s string){
	for i:=0;i<6;i++{
		time.Sleep(100 * time.Millisecond) //This is necessay else main will not wait for gorouitne to complete
		fmt.Println(s)
	}
}