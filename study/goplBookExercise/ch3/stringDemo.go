package main

import "fmt"

func main() {
	s := "Hello World"
	fmt.Println(s[3])
	//Raw string literal where no escape sequence is processed
	a:=`Go usage \t
	-h help
	-run to run`
	fmt.Println(a)
	b:="Go usage \t-h help -run to run \n"
	fmt.Println(b)
}
