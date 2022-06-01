package main

import (
	"fmt"
	"math"
)

func main() {
	altEvenOdd([]int{7, 10, 13, 14})
	altEvenOdd([]int{10, 12, 14})
	altEvenOdd([]int{7, 10, 13, 14, 16, 18, 19, 20, 21, 22, 23})
}

func altEvenOdd(arr []int) {
	//TC O(n), SC O(1)
	c := 0
	cmax := math.MinInt
	for i := 1; i < len(arr); i++ {
		if (arr[i]%2 == 0 && arr[i-1]%2 == 1) || (arr[i]%2 == 1 && arr[i-1]%2 == 0) {
			c++
		} else {
			cmax = int(math.Max(float64(cmax), float64(c)))
			c = 0
		}
	}
	c++
	cmax = int(math.Max(float64(cmax), float64(c)))
	fmt.Println(cmax)
}

func altEvenOddOther(arr []int) {
	c := 1
	cmax := math.MinInt
	for i := 1; i < len(arr); i++ {
		if (arr[i]%2 == 0 && arr[i-1]%2 == 1) || (arr[i]%2 == 1 && arr[i-1]%2 == 0) {
			c++
			cmax = int(math.Max(float64(cmax), float64(c)))
		} else {
			c = 1
		}
	}
	fmt.Println(cmax)
}
