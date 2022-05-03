package main

import "fmt"

func main() {
	fmt.Println("If Else")
	loginCount := 23

	if loginCount > 10 {
		fmt.Println("Watch out")
	} else if loginCount < 10 {
		fmt.Println("OK")
	} else {
		fmt.Println("Ideal")
	}

	if nums := 10; nums < 100 {
		fmt.Println("Example of defining variable in if and using it.")
	}

	switch i:=2+3;i{
	case 1,2,3:
		fmt.Println("Not 5")
	default:
		fmt.Println("5")
	}

	//var i interface{} = 2.0
	var i interface{} = [3]int{}
	switch i.(type){
	case int:
		fmt.Println("Int")
	case float64:
		fmt.Println("Float64")
	case [30]int:
		fmt.Println("[30]int")
	case [3]int:
		fmt.Println("[3]int")
	default:
		fmt.Println("Some other type")
	}
}
