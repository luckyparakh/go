package main

import "fmt"

func main() {
	f := square()
	fmt.Println(f())
	fmt.Println(f())
	g := square()
	fmt.Println(g())
	fmt.Println(g())
}

func square() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
