package main

import "fmt"

func main() {
	/*
		To create a slice independent of parent slice use Copy.

		The first is the destination slice, and the
		second is the source slice. The function copies as many values as it can from source to
		destination, limited by whichever slice is smaller, and returns the number of elements
		copied. The capacity of x and y doesn’t matter; it’s the length that’s important.
	*/

	var x = []int{1, 2, 3}
	y := make([]int, 2)
	copy(y, x)
	fmt.Println(y, x)

	var xx = []int{1, 2, 3}
	yy := make([]int, 5)
	copy(yy, xx)
	fmt.Println(yy, xx)

	z := make([]int, 1)
	copy(z, xx[1:])

	a := []int{1, 2, 3, 4}
	num := copy(a[:3], a[1:])
	fmt.Println(a, num)

	// CAn also use array
	q := []int{1, 2, 3}
	p := make([]int, 1)
	r := [4]int{5, 6, 7, 8}
	copy(q, r[:]) // array can be converted to slice using [:]
	copy(r[:], p)
	fmt.Println(q, r) // [5,6,7],[0,6,7,8]

}
