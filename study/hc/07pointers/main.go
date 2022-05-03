package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	mystr:="rp"
	var ptr *string
	fmt.Println("Default Value of ptr is ", ptr)
	ptr = &mystr
	fmt.Println("Value of ptr is ", ptr)

	myvar := 23
	var ptr1 = &myvar
	fmt.Println("Print ptr1", ptr1)
	fmt.Println("Print what ptr1 contains", *ptr1)
	*ptr1+=2
	fmt.Println("Value of myvar", myvar)
}
