package main

import "fmt"

func main() {
	fmt.Println("Method")
	me := User{"Rishi", "riship@go.dev", 34, true}
	fmt.Println(me)
	me.getStatus()
	me.setEmail()
	fmt.Println(me)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

//Methods
func (u User) getStatus() {
	fmt.Println(u.Status)
}

//Pass by value (copy of u has been passed)
func (u User) setEmail() {
	u.Email = "dummy@go.dev"
	fmt.Println("Email now is", u.Email)
}
