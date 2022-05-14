package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	echo1(args)
}

func echo1(args []string) {
	//args[0] prints file name
	fmt.Println("Echo1")
	for _, arg := range args[:] {
		fmt.Println(arg)
	}
}
