package main

import (
	"fmt"
	"math"
)

func main() {
	maxDiff([]int{2, 3, 10, 6, 4, 8, 1})
	maxDiffBetter([]int{2, 3, 10, 6, 4, 8, 1})
}
func maxDiffBetter(arr []int) {
	min := arr[0]
	var diff float64 = float64(arr[1] - min)
	for i := 0; i < len(arr); i++ {
		diff = math.Max(diff, float64(arr[i]-min))
		min = int(math.Min(float64(min), float64(arr[i])))
	}
	fmt.Println(diff)
}
func maxDiff(arr []int) {
	//TC o(n2)
	max := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j]-arr[i] > max {
				max = arr[j] - arr[i]
			}
		}
	}
	fmt.Println(max)
}
