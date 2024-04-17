package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("close")
			return
		default:
			fmt.Println("Do work")
		}
	}
}

// func main() {
// 	done := make(chan struct{})
// 	go doWork(done)
// 	time.Sleep(1 * time.Second)
// 	close(done)
// }
