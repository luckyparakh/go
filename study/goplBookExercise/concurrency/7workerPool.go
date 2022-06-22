package main

import "time"

type Job struct {
	id        int
	randomNum int
}
type Result struct {
	job Job
	sum int
}

var job = make(chan int, 10)
var result = make(chan int, 10)

func main() {

}

func sumDigit(num int) int{
	sum := 0
	for num != 0 {
		sum += num % 10
		num = num / 10
	}
	time.Sleep(2000 * time.Millisecond)
	return sum
}

func allocate(numOfJob int){
	for i:=0;i<numOfJob;i++{
		job<-
	}
}
