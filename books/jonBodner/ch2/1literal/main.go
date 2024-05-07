package main

import "fmt"

// P19
// string literal
func main() {

	// These contain zero or more rune literals. They are called “interpreted”
	// because they interpret rune literals (both numeric and backslash escaped)
	// into single characters.

	s := "Abc\n"
	fmt.Println(s)

	// 	One rune literal backslash escape is not legal in a string literal: the
	// single quote escape. It is replaced by a backslash escape for double
	// quotes.
	// The only characters that cannot appear in an interpreted string literal are unescaped
	// backslashes, unescaped newlines, and unescaped double quotes.

	// s1 := "Abc\'" // Error unknown escape
	// s1:="ABC\" // Error

	// If you use an
	// interpreted string literal and want your greetings on a different line from your
	// salutations and want “Salutations” to appear in quotes, you need to type "Greetings
	// and\n\"Salutations\"".
	s1 := "Greetings and \n \"GM\""
	fmt.Println(s1)

	// OR use raw string literal
	s2 := `Greetings and
        "GM"`
	fmt.Println(s2)

}
