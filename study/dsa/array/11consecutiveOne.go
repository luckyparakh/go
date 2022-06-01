package main

import (
	"fmt"
	"math"
)

func main() {
	consOne([]int{0, 1, 1, 0, 1, 0})
	consOne([]int{0, 0})
	consOne([]int{1, 1, 1, 1})
	fmt.Println("------------Another")
	consOneOther([]int{0, 1, 1, 0, 1, 0})
	consOneOther([]int{0, 0})
	consOneOther([]int{1, 1, 1, 1})
}

func consOne(arr []int) {
	s, e, maxLen := 0, 0, -1
	for _, i := range arr {
		if i == 1 {
			e++
		} else {
			maxLen = int(math.Max(float64(maxLen), float64(e-s)))
			e++
			s = e
		}
	}
	fmt.Println(int(math.Max(float64(maxLen), float64(e-s))))
}

func consOneOther(arr []int) {
	curr,res:=0,0
	for _, i := range arr {
		if i==0{
			curr=0
		}else{
			curr++
			res=int(math.Max(float64(res),float64(curr)))
		}
	}
	fmt.Println(res)
}
