package main

import "fmt"

func main() {
	reverse([]int{30, 20, 40, 50})
}

func reverse(arr []int) {
	high := len(arr)-1
	low := 0
	for low <= high {
		arr[low], arr[high] = arr[high], arr[low]
		low++
		high--
	}
	fmt.Println(arr)
}
