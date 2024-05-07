package main

import "fmt"

// Although it legal but it should be avoided as it is hard to find who changed this value
// declare variables in the package block that are effectively immutable , use const
// Also GO complier will not complain even package level variable is unused.
var rrr = 10

// s:=10 // can't use := for package level variable

func main() {
	var x int = 10           // sets x to 10
	var y int                // sets y to zero value of int which is 0
	var a, b int = 100, 1000 // define multiple in same line
	fmt.Println(x, y)
	x = 103
	fmt.Println(a, b)

	var c, d = 10, "hello" // define multiple in same line with diff types, as each literal has it default type like for numbers it is int and for string it is string
	fmt.Println(c, d)

	e := 34.4           // It will define and assign value and auto get the type by its literal
	e, f := 34.4, "abc" // e is already defined
	fmt.Printf("%T %T \n", e, f)

	// Can be defined in func
	const w rune = 'a'
	// q = 12 // Compile error
	// w = 'd' // Compile Error
}

// Constants in Go are a way to give names to literals (bool, numeric, rune, string literal etc). There is no way
// in Go to declare that a variable is immutable.

const q = 10 // Immutable

/*
Constants can be typed or untyped. An untyped constant works exactly like a literal;
it has no type of its own but does have a default type that is used when no other type
can be inferred.
*/
const u = 100

var x = u
var x1 byte = u
var xy float32 = u

const u1 int = 100

// var u2 byte = u1 // Error as u1 is a type constant
