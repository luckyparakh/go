package main

import "fmt"

func main() {
	y := []string{"Hello", "Hi", "Namaskar"}
	for _, v := range y {
		switch size := len(v); size { // Variable is init and used. It scope in available in all branches of switch
		case 1, 2: // use comma if few values has same logic
			fmt.Println(v, "is too short")
		case 4, 5: // No need of break
			a := 2
			fmt.Println(a) // a is scoped only in this case block
			fmt.Println(v, "is perfect sized")
		case 6, 7: // Can have blank case
		default: // is called when none of above matches
			fmt.Println(v, "is long")
			// fmt.Println(a) // undefined
		}
	}

	// Blank switch
	// A switch is called blank, if it  doesn’t specify the value that need to be compared
	// A regular switch only allow equality but blank switch allows to use any boolean comparison
	for _, v := range y {
		switch size := len(v); { // Variable is init
		case size < 3:
			fmt.Println(v, "is too short")
		case size > 7: // No need of break
			fmt.Println(v, "is long")
		default: // is called when none of above matches
			fmt.Println(v, "is perfect sized")
			// fmt.Println(a) // undefined
		}
	}

	// Switch is more cleaner as compared to if/else
}
