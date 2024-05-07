package main

import "fmt"

func main() {

	// Declaration
	var arr [3]int // array of capacity 3 with all 0 values
	fmt.Println(arr)

	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1)

	var arr2 = [...]int{10, 20, 30}
	fmt.Println(arr2)

	// Two arrays are equal if len and value of each element is same
	var arr3 = [3]int{1, 2, 3}
	fmt.Println(arr3 == arr1)

	fmt.Println(arr3 == arr2)
	/*
	Go considers the size of the array to be part of the type
	of the array. This makes an array that’s declared to be [3]int a different type from
	an array that’s declared to be [4]int. This also means that you cannot use a variable
	to specify the size of an array, because types must be resolved at compile time, not at
	runtime.
	CAn't use a type conversion to directly convert arrays of different sizes to identical types
	*/

	/*
	Due to this array are rarely used, expect where len of array is already known, 
	for e.g. some crypto func in std lib 
	*/
}
