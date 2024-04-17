package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	i := 0
	var ii atomic.Int32
	for j := 0; j < 5000; j++ {
		wg.Add(1)
		go func() {
			i++
			ii.Add(1)
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(i)
	fmt.Println(ii.Load())
}
