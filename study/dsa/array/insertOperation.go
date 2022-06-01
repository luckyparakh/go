package main

import "fmt"

func main() {
	arr := [10]int{20, 30, 50}
	fmt.Println(arr)
	insert(&arr, 3, len(arr), 40, 3)
	fmt.Println(arr)
	insert(&arr, 4, len(arr), 10, 1)
	fmt.Println(arr)

	insert(&arr, 5, len(arr), 60, 6)
	fmt.Println(arr)
	delete(&arr, 6, 4)
	fmt.Println(arr)
}

func insert(arr *[10]int, lenArr, capacity, value, position int) {
	//TC o(n)
	//Insertion can only happen if array has capacity
	if lenArr == capacity {
		return
	}
	for i := lenArr - 1; i >= position-1; i-- {
		arr[i+1] = arr[i]
	}
	arr[position-1] = value
}

func delete(arr *[10]int, lenArr, position int) {
	i:=0
	for i = position - 1; i < lenArr-1; i++ {
		arr[i] = arr[i+1]
	}
}
