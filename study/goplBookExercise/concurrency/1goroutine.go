package main

import "fmt"

func printMsg() {
	fmt.Println("Called via routine")
}
func main() {
	go printMsg()
	fmt.Println("In Main")
}
