package main

import "fmt"

type person struct {
	age  int
	name string
}

func main() {
	// 1st style
	// Need to give values to each field and the same order as fields are defined
	p1 := person{23, "ram"}

	// 2ns style;Map like style
	// Need not to assign value to each field, if any field is not given then its zero value is assigned
	p2 := person{
		name: "hanu",
	}

	// Both the style can be mixed
	fmt.Println(p1, p2)

	// Anonymous Struct
	p := struct {
		age  int
		name string
	}{
		age:  3,
		name: "a",
	}
	fmt.Println(p)

	var animal struct {
		size int
	}
	animal.size = 10

	/*
		Go doesn’t allow comparisons between variables of different primitive types,
		Go doesn’t allow comparisons between variables that represent structs of different
		types.
		Go does allow you to perform a type conversion from one struct type to
		another if the fields of both structs have the same names, order, and types. Let’s
	*/

	type PersonOne struct {
		Age  int
		Name string
	}
	type PersonTwo struct {
		Age  int
		Name string
	}
	p11 := PersonOne{34, "P11"}
	p12 := PersonTwo{34, "P12"}
	// fmt.Println(p11 == p12) // (mismatched types PersonOne and PersonTwo)

	var p121 PersonTwo
	p121 = PersonTwo(p11) // Conversion is allowed
	fmt.Println(p121)

	// Can't convert PersonOne or PersonTwo into PersonThree as order of Field is not same
	type PersonThree struct {
		Name string
		Age  int
	}
	// p13 := PersonThree{"P13", 23}
	var anonymousPerson struct {
		Age  int
		Name string
	}

	/*
		In case of comparing structs, if one of them is anonymous struct then can convert one to another also
		can compare one to another provided both structs have same names, order and type
	*/
	anonymousPerson = p11               // valid
	fmt.Println(anonymousPerson == p12) // false
	// fmt.Println(anonymousPerson == p13) // Type Mismatch
}
