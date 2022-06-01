package main

import "fmt"

func main() {
	freqSortedArr([]int{10, 10, 50, 60})
	fmt.Println("---------------------Better")
	freqSortedArrBetter([]int{10, 10, 50, 60})
	freqSortedArrBetter([]int{25})
}

func freqSortedArr(a []int) {
	//TC o(n) SC o(n)
	hm := make(map[int]int)
	for i := 0; i < len(a); i++ {
		hm[a[i]]++
	}
	fmt.Println(hm)
}

func freqSortedArrBetter(a []int) {
	//TC o(n) SC o(1)
	freq := 1
	i:=1
	for ; i < len(a); i++ {
		if a[i] == a[i-1] {
			freq++
		} else {
			fmt.Printf("%d:%d\n", a[i-1], freq)
			freq = 1
		}
	}
	if i==len(a){
		fmt.Printf("%d:%d\n", a[i-1], freq)
	}
}
