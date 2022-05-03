package main

import (
	"fmt"
	"sync"
)

var counter int = 0
var wg sync.WaitGroup = sync.WaitGroup{}
var mu sync.Mutex =sync.Mutex{}
func main(){

	for i:=0;i<10;i++{
		wg.Add(1)
		go read()
	}
	wg.Wait()
}

func read(){
	mu.Lock()
	increment()
	mu.Unlock()
	fmt.Println("Hello", counter)
	wg.Done()
}

func increment(){
	counter++
}