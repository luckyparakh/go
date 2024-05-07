package main

import "fmt"

func main() {

	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}
	// Multiple variable
	for i, j := 0, 0; i < 3 || j < 2; i, j = i+1, j+1 {
		fmt.Println(i, j) // 0,0 1,1 2,2
	}
	// fmt.Println(i,j) // only scoped till for

	// Condition only loop
	i, j := 0, 0
	for i < 3 || j < 2 {
		fmt.Println(i, j) // 0,0 1,1 2,2
		i, j = i+1, j+1
	}

	// for { Infinite for loop

	// }

	// use a for-range loop only to iterate over the built-in compound types and user-defined types that are based on them.
	m := make(map[string]int, 3)
	m["a"] = 1
	m["b"] = 2
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println(k)
	}
	for _, v := range m {
		fmt.Println(v)
	}

	// In maps while iterating k,v are not displayed or published in the order. Map is unordered.

	// in for-range loop value is copy
	// From GO1.22, for each iteration new value variable was created. Before GO1.22, same variable is used for each iteration.
	// Due to this, there is bug while creating closure in GO
	y := []int{2, 4, 6, 8}
	for _, v := range y {
		v *= 2
	}
	fmt.Println(y)

	// ҙ has unicode 1177 and same can't be stored in bytes, hence when loop encounter multi byte rune it converts UTF-8 to 32bit sized
	// and offset is increased by that number of bytes.
	// ﻼ takes 3 bytes & ҙ takes 2 bytes
	s := "aҙﻼd"
	for i, v := range s {
		fmt.Println(i, v, string(v))
	}
	/*
		0 97 a
		1 1177 ҙ
		3 100 d
	*/
	for i := 0; i < len(s); i++ {
		fmt.Println(i, string(s[i]))
	}
	/*
		0 a
		1 Ò
		2
		3 ï
		4 »
		5 ¼
		6 d
	*/

}
