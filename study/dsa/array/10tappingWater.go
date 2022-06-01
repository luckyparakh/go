package main

import (
	"fmt"
	"math"
)

func main() {
	tappingWater([]int{5, 0, 6, 2, 3})
	tappingWater([]int{2})
	tappingWater([]int{3, 0, 1, 2, 5})
	fmt.Println("-----------------------Better")
	tappingWaterBetter([]int{5, 0, 6, 2, 3})
	tappingWaterBetter([]int{2})
	tappingWaterBetter([]int{3, 0, 1, 2, 5})
}

func tappingWater(arr []int) {
	//TC O(n2);SC o(1)
	quantity := 0
	for i := 1; i < len(arr)-1; i++ {
		lMax := arr[i]
		rMax := arr[i]
		for j := i - 1; j >= 0; j-- {
			lMax = int(math.Max(float64(lMax), float64(arr[j])))
		}
		for j := i + 1; j < len(arr); j++ {
			rMax = int(math.Max(float64(rMax), float64(arr[j])))
		}
		quantity += int(math.Min(float64(lMax), float64(rMax))) - arr[i]
	}
	fmt.Println(quantity)
}

func tappingWaterBetter(arr []int) {
	//TC O(n2);SC O(1)
	lArr := len(arr)
	lMax := make([]int,lArr)
	rMax := make([]int,lArr)
	quantity:=0
	lMax[0]=arr[0]
	rMax[lArr-1] = arr[lArr-1]
	// fmt.Println(lMax)
	// fmt.Println(rMax)
	for i := 1; i < lArr; i++ {
		lMax[i] = int(math.Max(float64(arr[i]), float64(lMax[i-1])))
	}
	// fmt.Println(lMax)
	for i := lArr - 2; i >= 0; i-- {
		rMax[i] = int(math.Max(float64(arr[i]), float64(rMax[i+1])))
	}
	// fmt.Println(rMax)
	for i:=0;i<lArr;i++{
		quantity += int(math.Min(float64(lMax[i]), float64(rMax[i]))) - arr[i]
	}
	fmt.Println(quantity)
}
