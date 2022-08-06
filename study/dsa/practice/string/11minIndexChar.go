//https://practice.geeksforgeeks.org/batch-problems/minimum-indexed-character-1587115620/0/?track=DSASP-Strings&batchId=154

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minIndexChar("geeksforgeeks", "set"))
	fmt.Println(minIndexChar("adcffaet", "onkl"))
}

func minIndexChar(text, pat string) int {
	letterIndex := [128]int{}
	minIndex := math.MaxInt
	for k, v := range text {
		if letterIndex[v] == 0 {
			//To avoid confusion if min Index is zero from initailzed value
			letterIndex[v] = k + 1
		}
	}
	// fmt.Println(letterIndex)
	for _, v := range pat {
		if letterIndex[v] != 0 {
			minIndex = int(math.Min(float64(minIndex), float64(letterIndex[v])))
		}
	}
	if minIndex == math.MaxInt {
		return -1
	}
	return minIndex - 1
}
