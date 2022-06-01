package main

import "fmt"

func main() {
	arr := []int{20, 30, 40, 50}
	isSorted(arr)
	arr1 := []int{20, 3, 40, 50}
	isSorted(arr1)
	arr2 := []int{20}
	isSorted(arr2)
}

func isSorted(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
