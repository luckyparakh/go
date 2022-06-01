package main

import (
	"fmt"
	"sort"
)

func main() {
	findLeader([]int{7, 10, 4, 3, 6, 5, 2})
	findLeader([]int{30, 20, 10})
	findLeader([]int{7, 10, 4, 10, 6, 5, 2})
	fmt.Println("-------------------better")
	findLeaderBetter([]int{7, 10, 4, 3, 6, 5, 2})
	findLeaderBetter([]int{30, 20, 10})
	findLeaderBetter([]int{7, 10, 4, 10, 6, 5, 2})
}

func findLeader(arr []int) {
	//TC O(n2),SC o(n)
	tmp := []int{}
	for i := 0; i < len(arr); i++ {
		biggerNumFound := false
		for j := i + 1; j < len(arr); j++ {
			if arr[j] >= arr[i] {
				biggerNumFound = true
				break
			}
		}
		if !biggerNumFound {
			tmp = append(tmp, arr[i])
		}
	}
	fmt.Println(tmp)
}

func findLeaderBetter(arr []int) {

	arrLen := len(arr)
	curr_leader := arr[arrLen-1]
	tmp := []int{curr_leader}
	for i := arrLen - 2; i >= 0; i-- {
		if arr[i] > curr_leader {
			curr_leader = arr[i]
			tmp = append(tmp, curr_leader)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(tmp)))
	fmt.Println(tmp)
}