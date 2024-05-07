package main

import "fmt"

func main() {
	/*
		In GO string is not made of rune but sequence of bytes
	*/

	s := "abcdef"
	s1 := s[2:]
	fmt.Println(s1)
	// s1[0] = "a" // immutable they don’t have the modification problems that slices of slices do

	s2 := "hello Ǎ"
	fmt.Println(len(s2)) // 8 Ǎ is of 2 bytes
	fmt.Println(s2[:7])  // hello � because only 1st byte of Ǎ is copied

	/*
		Even though Go allows you to use slicing and indexing syntax with
		strings, you should use it only when you know that your string contains only characters that take up one byte.
	*/

	// String can be converted into rune/byte
	q := "qwerty"
	var r = []rune(q)
	var d = []byte(q)
	fmt.Println(r, d)
	var g = string([]byte{113, 119})
	fmt.Println(g)

	c := 65
	e := string(c) // warning , conversion from int to string yields a string of one rune, not a string of digits
	fmt.Println(e)
}
