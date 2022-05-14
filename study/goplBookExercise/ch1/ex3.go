package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args
	echo1(args)
	echo2(args)
	echo3(args)

	fmt.Println(benchmark(echo2,args)) // to see diff. check for bigger args
	fmt.Println(benchmark(echo3,args))
}

func benchmark(f func([]string), args []string) int {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		f(args)
	}
	end := time.Now()
	return int(end.Sub(start).Nanoseconds())
}
func echo1(args []string) {
	//args[0] prints file name
	for _, arg := range args[:] {
		fmt.Println(arg)
	}
}

func echo2(args []string) {
	// fmt.Println("Echo2")
	var s, sep string
	for _, arg := range args[1:] {
		s = s + sep + arg
		sep = " "
	}
	// fmt.Println(s)
}

func echo3(args []string) {
	// fmt.Println("Echo3")
	strings.Join(args, " ")
	// fmt.Println(s)
}
