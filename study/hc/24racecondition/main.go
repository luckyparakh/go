package main

import (
	"fmt"
	"sync"
)

func main(){
	fmt.Println("Race Condition")
	var wg sync.WaitGroup
	mut := &sync.Mutex{}
	var data = []int{0}
	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex){
		fmt.Println("One")
		mut.Lock()
		data = append(data,1)
		mut.Unlock()
		wg.Done()
	}(&wg,mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex){
		fmt.Println("Two")
		mut.Lock()
		data = append(data,2)
		mut.Unlock()
		wg.Done()
	}(&wg,mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex){
		fmt.Println("Three")
		mut.Lock()
		data = append(data,3)
		mut.Unlock()
		wg.Done()
	}(&wg,mut)
	wg.Wait()

}