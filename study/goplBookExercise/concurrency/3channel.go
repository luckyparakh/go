package main

import (
	"fmt"
)

func printMsg(done chan bool) {
	fmt.Println("Called via routine")
	done <- true
}
func main() {
	ch := make(chan bool)
	go printMsg(ch)
	<-ch
	fmt.Println("In Main")
}
