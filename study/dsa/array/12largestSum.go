package main

import (
	"fmt"
	"math"
)

func main() {
	maxSum([]int{-5, 1, -2, 3, -1, 2, -2})
}

func maxSum(arr []int) {
	//TC O(n), SC O(1)
	max := math.MinInt
	for i := 1; i < len(arr); i++ {
		arr[i] = int(math.Max(float64(arr[i]), float64(arr[i-1]+arr[i])))

		if arr[i]>max{
			max=arr[i]
		}
	}
	fmt.Println(max)
}
