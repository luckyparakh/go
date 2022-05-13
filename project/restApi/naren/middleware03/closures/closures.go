package closures

import (
	"fmt"
	"time"
)

func NumGen() func() int {
	i := 5
	return func() int {
		i++
		return i
	}
}

func FibClosure() func() int {
	a := 0
	b := 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func Timer(f func(int, int) int) func(int, int) {
	return func(a, b int) {
		start := time.Now()
		f(a, b)
		end := time.Now()
		fmt.Println(end.Sub(start))
	}
}
func Add(a, b int) int {
	return a + b
}
