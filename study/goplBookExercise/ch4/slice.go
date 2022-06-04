package main

import "fmt"

func main() {
	months := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "Jun"}
	fmt.Println(months)
	a := months[:2]
	fmt.Println(len(a), cap(a))
	b := months[2:5]
	fmt.Println(len(b), cap(b))
	x := make([]int, 0)
	x = append(x, 1) //Push
	top := x[len(x)-1]
	fmt.Println(x, top)
	x = append(x, 11, 12, 13)
	top = x[len(x)-1]
	fmt.Println(x, top)
	x = x[:len(x)-1]
	top = x[len(x)-1]
	fmt.Println(x, top)

	//remove from middle
	i:=1
	x = append(x[:i], x[i+1:]...)
	fmt.Println(x)
}
