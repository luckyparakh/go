package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	result := add(3, 2)
	fmt.Println(result)

	pResult,_ := proAdd(3, 2,4,5,1)
	fmt.Println(pResult)

}

func add(i1, i2 int) int {
	return i1 + i2
}

func proAdd(values ...int) (int,string) {
	total:=0
	for _,val := range values{
		total+=val
	}
	return total,"ProAdd"
}
