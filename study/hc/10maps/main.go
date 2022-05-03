package main

import "fmt"

func main() {
	fmt.Print("Maps")
	var courses = make(map[string]string)
	courses["JS"] = "Javascript"
	courses["Py"] = "Python"
	courses["PHP"] = "Php"

	fmt.Println(courses)
	fmt.Println(courses["JS"])
	delete(courses, "PHP")
	fmt.Println(courses)

	for key, val := range courses {
		fmt.Println(key, ":", val)
	}

	state := make(map[string]int)
	state["Raj"] = 1000
	state["UP"] = 2000
	fmt.Println(state)
	delete(state,"UP")
	fmt.Println(state)
	fmt.Println(state["UP"]) // It will print zero, how to know if it is error or actaul value is zero.
	v, err := state["UP"]
	fmt.Println(v,err) //0 false (false means key is not there)

	//Pass by value
	states := make(map[string]int)
	states = map[string]int{
		"Raj":1000,
		"UP":2000,
		"HP":3000,
	}
	fmt.Println(states)
	statesCp:=states
	fmt.Println(statesCp)
	delete(states, "HP")
	fmt.Println(states)
	fmt.Println(statesCp)
}
