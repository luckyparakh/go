package main

import "fmt"

func main() {
	rotateArr([]int{1, 2, 3, 4, 5}, 1)
	rotateArr([]int{1, 2, 3, 4, 5}, 2)
	fmt.Println("--------------------better")
	rotateArrBetter([]int{1, 2, 3, 4, 5}, 1)
	rotateArrBetter([]int{1, 2, 3, 4, 5}, 2)
	fmt.Println("--------------------best")
	rotateArrBest([]int{1, 2, 3, 4, 5}, 1)
	rotateArrBest([]int{1, 2, 3, 4, 5}, 2)
}

func rotateArr(arr []int, rotateBy int) {
	//TC o(n*rotateBy)
	//Sc:o(1)
	for j := 0; j < rotateBy; j++ {
		for i := 0; i < len(arr)-1; i++ {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	fmt.Println(arr)
}

func rotateArrBetter(arr []int, rotateBy int) {
	//TC o(n)
	//Sc o(rotateBy)
	tmp := make([]int, rotateBy)
	for j := 0; j < rotateBy; j++ {
		tmp[j] = arr[j]
	}

	for i := 0; i < len(arr)-rotateBy; i++ {
		arr[i] = arr[rotateBy+i]
	}

	j := 0
	//for i, j := len(arr)-rotateBy, 0; i < len(arr); i, j = i+1, j+1 {
	for i := len(arr) - rotateBy; i < len(arr); i++ {
		arr[i] = tmp[j]
		j++
	}
	fmt.Println(arr)
}
func rotateArrBest(arr []int, rotateBy int) {
	//TC o(n)
	//Sc o(1)
	rev(arr, 0, rotateBy-1)
	rev(arr, rotateBy, len(arr)-1)
	rev(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
func rev(arr []int, low, high int) {
	for low < high {
		arr[low], arr[high] = arr[high], arr[low]
		low++
		high--
	}
}
