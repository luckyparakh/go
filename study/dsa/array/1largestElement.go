package main

import (
	"fmt"
)

func main() {
	arr := [5]int{20, 20, 20, 20, 20}
	largestEle(arr)
	secondLargestEle(arr)
	arr1 := [5]int{20, 2, -2, 2, 2}
	secondLargestEle(arr1)
}

func largestEle(arr [5]int) {
	max := -100000
	maxIndex := -1
	for k, v := range arr {
		if v > max {
			max = v
			maxIndex = k
		}
	}
	fmt.Println(maxIndex)
}

func secondLargestEle(arr [5]int) {
	maxIndex1 := -1
	// maxIndex := -1
	val := -1000
	val1 := -1000
	for i := 0; i < len(arr); i++ {
		if arr[i] > val {
			// val1 = val
			val = arr[i]
			// maxIndex1 = maxIndex
			// maxIndex=i
		} else if arr[i] != val && arr[i] > val1 {
			val1 = arr[i]
			maxIndex1 = i
		}
	}
	fmt.Println(maxIndex1)
}
