package main

import "fmt"

func main() {
	var emptyMap map[int]int
	nilMap := map[int]string{}
	notNilMap := map[int]int{1: 10, 2: 20, 3: 30}

	// Can read from empty map
	a := emptyMap[1]
	fmt.Println(a) // would be zero as 0 is default value for int
	b := nilMap[1]
	fmt.Println(b) // would be "" as "" is default value for string
	c := notNilMap[1]
	fmt.Println(c)
	// Write
	nilMap[1] = "1"
	fmt.Println(nilMap)
	// Write to empty map will fail
	emptyMap[1] = 1
	fmt.Println(emptyMap)
}
