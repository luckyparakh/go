package main

import "fmt"

func main() {
	// defer fmt.Println("World")
	// defer fmt.Println("One")
	// fmt.Println("Hello")
	// printNum()
	//fmt.Println(panicker())

	fmt.Println("Panicker2==============")
	fmt.Println("Start")
	panicker2()
	fmt.Println("End")
}

func printNum() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func panicker() int{
	//Defer will be called before exit (even in case of panic)
	fmt.Println("Start")
	defer fmt.Println("Defer")
	panic("Panic")
}


func panicker2(){
	//Defer will be called before exit (even in case of panic) and if defer has recover than system will recover
	fmt.Println("About to Panic")
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println(err)
		}
	}()
	panic("Paniced")
	fmt.Println("Will not print")
}