package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go printMsg(i, &wg)
	}
	wg.Wait()
}

func printMsg(i int, waitGroup *sync.WaitGroup) {
	fmt.Printf("Printing %v\n", i)
	waitGroup.Done()
}
