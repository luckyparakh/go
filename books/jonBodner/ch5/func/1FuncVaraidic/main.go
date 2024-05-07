package main

import "fmt"

type f func(x, y int) (int, error)

func main() {

	// GOlang doesn't have option of optional parameters
	// But with help of struct as input to func same can be achieved
	user1 := User{
		Name:    "abc",
		Address: "xyz",
	}
	user(user1)
	user2 := User{
		Name: "abc",
	}
	user(user2)

	// Variadic
	variadic(3, 1, 2, 3)
	variadic(2, 1, 2)
	// In case if slice literal is used, use '...'
	variadic(2, []int{10, 20}...)
	q := []int{30, 40}
	variadic(2, q...)

	// Func as variable
	var a func(x, y int) (int, error)
	a = func(x, y int) (int, error) {
		return x + y, nil
	}
	fmt.Println(a(2, 3))

	// Func as type
	var b f
	b = func(x, y int) (int, error) {
		return x * y, nil
	}
	fmt.Println(b(2, 3))

	// Anonymous func it has func type, input and output parameters but no name
	d := func() int { return 10 }
	fmt.Println(d())
}

// Variadic can only be used as final parameter
func variadic(c int, e ...int) {
	// e is slice of []int
	fmt.Println(c)
	for _, v := range e {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

type User struct {
	Name    string
	Address string
}

func user(u User) {
	fmt.Println(u)
}
