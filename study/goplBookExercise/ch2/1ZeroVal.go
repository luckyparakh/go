package main

import "fmt"

type ZV struct {
	dummy string
	num   int
}

func main() {
	// The zero value of an aggregate type like an array or a struct has the zero value of all of its elements or fields.
	zv := ZV{}
	fmt.Println(zv)
}
