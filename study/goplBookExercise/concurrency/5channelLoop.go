package main

import "fmt"

func main() {
	ch := make(chan int)
	go printMsg(ch)
	for i := range ch {
		fmt.Println(i)
	}
}

func printMsg(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
