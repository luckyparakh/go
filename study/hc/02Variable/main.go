package main

import (
	"fmt"
	"strconv"
)

const LoginToken = "asdre" // Public scoped as it has L capital
// jwtToken:= "1234" // This style is not allowed
var jwtToken = "1234" // It is okay and these package scoped variable

func main() {
	fmt.Println("Variables")
	var username string = "Rishi"
	fmt.Printf("Variable '%v' is of type %T \n", username, username)

	var isLogged bool = true
	fmt.Printf("Variable '%v' is of type %T \n", isLogged, isLogged)

	var smallVal uint8 = 255
	fmt.Printf("Variable '%v' is of type %T \n", smallVal, smallVal)

	var f32 float32 = 255.3456787665
	fmt.Printf("Variable '%v' is of type %T \n", f32, f32)

	var f64 float64 = 255.3456787665
	fmt.Printf("Variable '%v' is of type %T \n", f64, f64)

	//default value
	var i int
	fmt.Printf("Variable '%v' is of type %T \n", i, i)

	//Implicit value
	var add = "Hello"
	fmt.Printf("Variable '%v' is of type %T \n", add, add)

	//No Var value
	greet := "Hello"
	fmt.Printf("Variable '%v' is of type %T \n", greet, greet)

	fmt.Println(jwtToken)

	conversion()
	printConstants()

}

func conversion(){
	var i int = 2
	fmt.Println(float32(i))
	fmt.Println(string(i)) // It will print unicode char represented by 2, hence strconv
	fmt.Println(strconv.Itoa(i))
}

const (
	a = iota
	b = iota
	c = iota
)
const (
	//iota reset in new block
	d = iota
	e = iota
	f = iota
)

const (
	//iota reset in new block
	x = iota + 1 // Initialize iota by 1
	y // Y will be auto assigned iota ie 2
	z
)
const (
	isAdmin = 1 << iota // 00000001 = 1
	isUser // 00000010 = 2

	canSeeAsia
	canSeeAfrica

)
const (
	_  = 1 << (iota * 10) // ignore the first value
	KB                    // decimal:       1024 -> binary 00000000000000000000010000000000
	MB                    // decimal:    1048576 -> binary 00000000000100000000000000000000
	GB                    // decimal: 1073741824 -> binary 01000000000000000000000000000000
)
func printConstants(){
	fmt.Println(a,b,c)
	fmt.Println(d,e,f)
	fmt.Println(x,y,z)
	role := isAdmin | isUser | canSeeAsia
	fmt.Println(role)
	fmt.Println(canSeeAsia)
	fmt.Println(canSeeAfrica)
	fmt.Println(isAdmin & role == isAdmin)
	fmt.Println(isUser & role == isUser)
	fmt.Println(isUser & role == canSeeAfrica)
}