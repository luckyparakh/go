package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		defer is used with func/method/closure
		Called in LIFO
		Defer is called after return
	*/
	lifo()
	clean()
	fmt.Println(changeReturn())

	// Although defer is called after return but process input values immediately
	input()

	addHello := prefixer("Hello")
	fmt.Println(addHello("Ram"))
}

func lifo() {
	// prints 3,2,1
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
}

func clean() (err error) {
	f, err := os.CreateTemp("", "test")
	// Using Anonymous func
	defer func() {
		// Using named variable err to assign error in any func
		err = f.Close()
		if err == nil {
			err = os.Remove(f.Name())
		}
	}()
	if err != nil {
		return err
	}
	fmt.Println(f.Name())
	return nil
}

// 3 will returned
func changeReturn() (a int) {
	defer func() {
		a = 3
	}()

	a = 2
	return a
}

func input() {
	fmt.Println("Input Starts")
	i := 1
	defer func(i int) {
		fmt.Println("defer", i)
	}(i)
	fmt.Println(i)
	i = 2
	fmt.Println(i)
}

func prefixer(o string) func(string) string {
	return func(i string) string {
		r := o + " " + i
		return r
	}
}
