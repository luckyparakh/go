package main

import "fmt"

// func main() {
// 	ch := make(chan string)
// 	anotherCh := make(chan string)
// 	go func() {
// 		ch <- "From Ch"
// 	}()
// 	go func() {
// 		anotherCh <- "From Another Channel"
// 	}()

// 	select {
// 	case msg := <-ch:
// 		fmt.Println(msg)
// 	case msg1 := <-anotherCh:
// 		fmt.Println(msg1)
// 	}
// }
