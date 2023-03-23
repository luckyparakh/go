package main

import "fmt"

func main() {
	a := "èlite"
	// length is 6 because string is byte array
	// ASCII character uses on byte but non ascii character uses more than 1 byte to store
	fmt.Printf("%v len is %d\n",a, len(a))
	b := []rune(a)
	fmt.Printf("%T %[1]v\n", b)
	c := []byte(a)
	fmt.Printf("%T %[1]v\n", c)
}
