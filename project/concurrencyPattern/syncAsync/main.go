//https://levelup.gitconnected.com/basic-parallel-computing-in-go-fda50894241c
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	aSyncCall()
	//syncCall()
}

func syncCall() {
	c, err := content("https://www.google.com")
	if err != nil {
		fmt.Printf("error = %v", err)
		os.Exit(1)
	}
	fmt.Printf("%v", countWords(c))
}

func aSyncCall() {
	startTime := time.Now().Unix()
	size := 10
	in := make(chan string, size)
	op := make(chan Counts, size)
	for i := 0; i < size*2; i++ {
		go contentAsync("https://www.google.com", in)
	}
	//go countAsync(in, op)
	for k := 0; k < 5; k++ {
		go countAsync(in, op)
	}

	// Only for reading purpose
	for j := 0; j < size*2; j++ {
		c := <-op
		fmt.Println(j, len(c))
	}
	fmt.Println(time.Now().Unix() - startTime)
}
