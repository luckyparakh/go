package main

import "fmt"

func main() {
	fmt.Println("Welcome to array")

	var myarr [4]string
	myarr[0] = "apple"
	myarr[1] = "banana"
	myarr[3] = "melon"

	fmt.Println("My arr contains:", myarr)
	var veg = [3]string{"Tomato", "Beans", "Spinach"}
	fmt.Println("Veg list:", veg, len(veg))

}
