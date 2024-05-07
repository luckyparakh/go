package main

import "fmt"

/*
Each place where declaration is done is called block
Type of Blocks
	Package Block
		Variable, const, type, func which are not declared in block is are in package block
	File Block
		Import define names for other packages that are valid for the file that contains the import statement.
		These names are in the file block.
	Func Block
		Top line of Func is one block and after that is another block
	Universal Block
		Is the block where global const are defined like true/false, built-in types like (int, string), built in func like (make,close) etc are defined in this block.
		And it contains all other blocks
*/

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x) // 10
		x := 20
		fmt.Println(x) //20; shadows x in outer block
	}
	fmt.Println(x) //10

	a := 10
	if a > 5 {
		fmt.Println(a)    // 10
		a, b := 20, 30    // new variable is defined not affect outer variable but it will shadow it
		fmt.Println(a, b) //20,30
	}
	fmt.Println(a) //10 := only reuses variable defined in the current scope

	e := 10
	if e > 5 {
		fmt.Println(e) //10
		e := 15
		fmt.Println(e)    // 15
		e, b := 20, 30    // := only defines variable in the scope
		fmt.Println(e, b) //20,30; shadows x in outer block
	}
	fmt.Println(e) //10
}
