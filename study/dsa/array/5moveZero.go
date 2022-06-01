package main

import "fmt"

func main() {
	moveZero([]int{8, 5, 0, 10, 0, 20})
	moveZero([]int{0, 0, 0, 10, 0, 0})
	moveZero([]int{0})
	fmt.Println("----------------------Better")
	moveZeroBetter([]int{8, 5, 0, 10, 0, 20})
	moveZeroBetter([]int{0, 0, 0, 10, 0, 0})
	moveZeroBetter([]int{0})
}
func moveZeroBetter(arr []int) {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
	fmt.Println(arr)
}
func moveZero(arr []int) {
	for i := 0; i < len(arr); {
		if arr[i] != 0 {
			i++
		} else {
			j := i + 1
			for ; j < len(arr); j++ {
				if arr[j] != 0 {
					arr[i], arr[j] = arr[j], arr[i]
					break
				}
			}
			if j == len(arr) {
				i = len(arr)
			}
		}
	}
	fmt.Println(arr)
}
