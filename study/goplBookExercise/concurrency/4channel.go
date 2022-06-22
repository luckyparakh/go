package main

import "fmt"

func square(num int, done chan int) {
	done <- num * num
}

func cube(num int, done chan int) {
	done <- num * num * num
}

func main() {
	num := 4
	sqStatus := make(chan int)
	cbStatus := make(chan int)
	go square(num, sqStatus)
	go cube(num, cbStatus)
	sq, cb := <-sqStatus, <-cbStatus
	op := sq + cb
	fmt.Println(op)
}
