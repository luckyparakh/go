package main

import "fmt"

func main() {
	// There are no classes in GO hence no inheritance/super etc.
	fmt.Println("Structs")
	rishi := User{"Rishi", "riship@go.dev", true, 35}
	fmt.Println(rishi)
	fmt.Println(rishi.Age)
	// use +v for structs
	fmt.Printf("My details: %+v", rishi)
}

// Keep 1st letter in capital so to make public
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
