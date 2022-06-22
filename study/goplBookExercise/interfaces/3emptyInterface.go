package main

import "fmt"

func printType(v interface{}) {
	fmt.Printf("Type %T value %[1]v\n",v)
}

func main() {
	printType(2)
	printType("str")
	type user struct{
		name string
		place string
	}
	u:=user{
		name: "RP",
		place: "Pali",
	}
	printType(u)
}
