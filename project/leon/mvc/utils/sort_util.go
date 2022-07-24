package utils

import "sort"

func BubbleSort(ele []int) {
	for i := 0; i < len(ele); i++ {
		for j := 0; j < len(ele)-i-1; j++ {
			if ele[j] > ele[j+1] {
				ele[j], ele[j+1] = ele[j+1], ele[j]
			}
		}
	}
}

func Sort(ele []int) {
	// After benchmarking, it is found that element less than 1000 simple Bubblesort is better
	// then inbuilt sort function.
	if len(ele) < 50 {
		BubbleSort(ele)
		return
	}
	sort.Ints(ele)
}
