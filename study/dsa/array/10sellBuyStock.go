//https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

package main

import "fmt"

func main() {
	sellBuyStock([]int{7, 1, 5, 3, 6, 4})
}

func sellBuyStock(arr []int) {
	profit := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			profit += arr[i] - arr[i-1]
		}
	}
	fmt.Println(profit)
}
