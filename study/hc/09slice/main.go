package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slice")
	var fruitList []string
	// var fruitList = []string{"Mango"}
	fruitList = append(fruitList, "Peach", "Mango")
	fmt.Println(fruitList)
	fruitList = append(fruitList, "Apple", "Cocunut")
	fmt.Println(fruitList[1:3])

	scores := make([]int, 4)
	scores[0] = 234
	scores[1] = 321
	scores[2] = 100
	scores[3] = 2
	fmt.Println(scores)
	// will give error
	// scores[4]=20
	// fmt.Println(scores)
	scores = append(scores, 12, 700)
	fmt.Println(scores)
	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println(sort.IntsAreSorted(scores))

	//how to remove value from slice
	var courses = []string{"js","java","php","go","py"}
	// ... is varidic implies multiple params can be passed
	// https://www.geeksforgeeks.org/variadic-functions-in-go/
	fmt.Println(append(courses[:2],courses[2+1:]...))
}
