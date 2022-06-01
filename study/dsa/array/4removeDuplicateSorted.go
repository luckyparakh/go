package main

import "fmt"

func main() {
	removeDup([]int{20, 20, 40, 40, 60})
}

func removeDup(arr []int) {
	for i := 0; i < len(arr)-1; {
		if arr[i] == -1 || arr[i] != arr[i+1] {
			i++
		} else {
			removeEle(arr, i+1)
		}
	}
	fmt.Println(arr)
}

func removeEle(arr []int, position int) {
	i := position
	for ; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[i] = -1
}
