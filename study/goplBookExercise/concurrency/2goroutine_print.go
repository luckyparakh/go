package main

import (
	"fmt"
	"time"
)

func printMsg() {
	fmt.Println("Called via routine")
}
func main() {
	go printMsg()
	time.Sleep(1 * time.Second)
	fmt.Println("In Main")
}
