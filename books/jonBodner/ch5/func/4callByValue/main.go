package main

import "fmt"

type Person struct {
	age int
}

func main() {
	// GO is call by value i.e. if a variable is passed to param of func then its value is copied
	i, j := 2, 3

	p := Person{10}
	wrongValue(i, j, p)
	fmt.Println(i, j, p)

	// Every type in GO is pass by value, some time a pointer can be passed
	// Like map/slice both are implemented as pointer hence values changed in func is also changed in main.
	// For slice it bit complicated, as if increase size of slice a new slice is created. Hence lengthen of slice in func should be avoided.
	// For example values in s1 are changed but values in s2 changed
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(m)
	s := []int{1, 2}
	s1 := []int{1, 2}
	fmt.Println(s, s1)
	printComposite(m, s, s1)

	fmt.Println(m)
	fmt.Println(s)
	fmt.Println(s1)
}

func printComposite(m map[string]int, s, s1 []int) {
	m["b"] = 20
	delete(m, "a")
	m["c"] = 30

	for k := range s {
		s[k] += 2
	}

	// 
	s1 = append(s1, 10)
	for k := range s1 {
		s1[k] += 2
	}
	
}
func wrongValue(i, j int, p Person) {
	i, j = 20, 30
	p.age = 100
	fmt.Println(i, j, p)
}
