package main

import "fmt"

func main() {
	leftMostRepating("abcccbcc")
	leftMostRepating("abc")
	fmt.Println("--------------------Better")
	leftMostRepatingBetter("abcccbcc")
	leftMostRepatingBetter("abc")
	fmt.Println("--------------------Non Repeating")
	leftMostNonRepatingBetter("abcccbcc")
	leftMostNonRepatingBetter("geeksforgeeks")
}
func leftMostRepatingBetter(s string) {
	var arr [256]int
	for _, v := range s {
		arr[v]++
	}
	for i, v := range s {
		if arr[v] > 1 {
			fmt.Printf("%d\n", i)
			return
		}
	}
	fmt.Printf("%d\n", -1)
}
func leftMostRepating(s string) {
	//TC : O(n)
	//Sc : O(n)
	freqMap := make(map[rune][]int)
	for i, v := range s {
		val, ok := freqMap[v]
		if !ok {
			freqMap[v] = []int{1, i}
		} else {
			freqMap[v] = []int{val[0] + 1, val[1]}
		}
	}
	fmt.Println(freqMap)
	min := 1000000
	for _, v := range freqMap {
		if v[0] > 1 && v[1] < min {
			min = v[1]
		}
	}
	if min==1000000{
		fmt.Printf("%d\n", -1)
	}else{
		fmt.Printf("%d\n", min)
	}
	

}

func leftMostNonRepatingBetter(s string) {
	var arr [256]int
	for _, v := range s {
		arr[v]++
	}
	for i, v := range s {
		if arr[v] == 1 {
			fmt.Printf("%d\n", i)
			return
		}
	}
	fmt.Printf("%d\n", -1)
}
