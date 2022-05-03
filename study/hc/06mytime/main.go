package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time")

	currentTime := time.Now();
	fmt.Println(currentTime)
	// Layout in format is constant, remember it
	// https://pkg.go.dev/time#Time.Format
	fmt.Println(currentTime.Format("02-01-2006 15:04:05 MST"))
	myTime := time.Date(2021,time.April,28,23,34,3,3,time.UTC)
	fmt.Println(myTime)
	fmt.Println(myTime.Format("2006-01-02 Monday MST"))

	// GOOS=windows go build . (it will create Window exe)
	// GOOS=linux go build . (it will create Linux installation file)
}