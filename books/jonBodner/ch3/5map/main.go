package main

import (
	"fmt"
	"maps"
)

func main() {
	var nilMap map[string]int
	fmt.Println(nilMap) // return default value of value type
	// nilMap["abc"] = 2 // panic: assignment to entry in nil map

	var emptyMap = map[string]int{} // map literal; it not same as nil map
	fmt.Println(emptyMap)
	emptyMap["a"] = 2 // can write in empty map; Type of values in Map can be anything.
	fmt.Println(emptyMap)

	/*
		Maps automatically grow as you add key-value pairs to them.
		If you know how many key-value pairs you plan to insert into a map, you can use
		make to create a map with a specific initial size.
		Passing a map to the len function tells you the number of key-value pairs in a map.
	*/
	makeM := make(map[string][]int, 10) // 10 is size
	fmt.Println(len(makeM))             // 0; cap is not a valid thing for map

	// Map is not comparable like slice, only can be compared to NIL
	// fmt.Println(emptyMap==nilMap) // Error
	fmt.Println(makeM == nil) // false

	// Key can be anything comparable, hence map and slice can't be the key
	// Value can be of any type
	// sliceMap:=make(map[[]string][]int, 10) // invalid key type

	// Maps in Golang are implemented using Hash map and hashing also for all types are already present in GO runtime.
	// GO runtime is part of all GO binary i.e. the Go runtime is compiled into every Go program

	wins := map[string]int{
		"Bharat": 10,
		"US":     5,
	}
	fmt.Println(wins["US"])    //5
	fmt.Println(wins["India"]) // 0; if key is not present then GO prints zero value of that type

	// Do diff between actaul Zero Value and key not present use comma ok expr
	if _, ok := wins["India"]; !ok {
		fmt.Println("Key not present")
	}

	delete(wins, "US")    // Delete US from Map
	delete(wins, "India") // Even key is not present it will not give error/panic
	fmt.Println(wins)
	clear(wins)
	fmt.Println(wins, len(wins)) // clear removes all elements also reduce len to 0

	m := map[int]int{
		1: 10,
		2: 20,
	}
	n := map[int]int{
		1: 10,
		2: 20,
	}
	fmt.Println(maps.Equal(m, n)) // true

	// Sets - Golang doesn't has set but can be created using MAP
	set := map[int]bool{}
	set[5] = true
	if !set[6] {
		fmt.Println("key is not present in set")
	}

	// Using empty struct as key is good as it used zero memory where bool uses 1 byte
	// But using struct makes code clumsier , also need to use comma ok notation
	setS := map[int]struct{}{}
	setS[5] = struct{}{}
	if _, ok := setS[6]; !ok {
		fmt.Println("key is not present in set")
	}

}
