package main

import "fmt"

func typeInterface(v interface{}) {
	if value, ok := v.(int); ok {
		fmt.Println(value)
	}else{
		fmt.Println("Type not supported")
	}
}


func typeSwitch(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println(v.(int))
	case string:
		fmt.Println(v.(string))
	default:
		fmt.Println("Not Valid Format")
	}
}

func main() {
	typeInterface(33)
	typeInterface("str")
	typeSwitch("abc")
	typeSwitch(23.3)
}
