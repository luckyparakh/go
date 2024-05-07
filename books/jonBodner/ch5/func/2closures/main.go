package main

import (
	"fmt"
	"sort"
)

/*
func defined inside func are special and called as closures

Inside func can access/modify outside func's variable
*/
func main() {
	a := 20
	f := func() {
		fmt.Println(a)
		a = 30
		fmt.Println(a)
	}
	f()
	fmt.Println(a)

	shadowing()

	// If you have a piece of logic that is repeated multiple times within a function, a closure can be used to remove that repetition.
	// Also it reduce repetition in a func, also reduce number package level func
	printer()

	// Func as input
	type City struct {
		Name       string
		Population int
	}
	cities := []City{
		{"Delhi", 500000}, {"Pali", 100000}, {"Jaipur", 300000},
	}

	// Func is passed as input to slice func, often these type of funcs (slice) are called high level func
	// In func is a closure also it captures cities
	sort.Slice(cities, func(i, j int) bool {
		return cities[i].Population < cities[j].Population
	})
	fmt.Println(cities)

	// Func can be returned as 

}

func shadowing() {
	fmt.Println("Shadowing")
	a := 20
	f := func() {
		fmt.Println(a)
		// below a shadows outer a
		a := 30
		fmt.Println(a)
	}
	f()
	fmt.Println(a)
}

func printer() {
	p := func(s string) {
		fmt.Println(s)
	}

	p("doing some validation")
	p("doing some work")
	p("returning")
}
