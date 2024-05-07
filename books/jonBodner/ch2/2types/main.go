package main

import "fmt"

func main() {
	var a byte = 2
	var b uint8 = 2
	if a == b {
		fmt.Println("byte and uint8 are same")
	}
	fmt.Printf("%T %T\n", a, b)

	/*
		On most 64-bit CPUs, int is a 64-bit signed integer, just like an int64.
	Because int isn’t consistent from platform to platform, it is a compile-time error to
	assign, compare, or perform mathematical operations between an int and an int32
	or int64 without a type conversion
	*/
	var a1 int = 2
	var b1 int64 = 2
	// if a1 == b1 { // not comparable
	// 	fmt.Println("byte and uint8 are same")
	// }
	fmt.Printf("%T %T\n", a1, b1)

}
