// https://practice.geeksforgeeks.org/batch-problems/sum-of-numbers-in-string-1587115621/0/?track=DSASP-Strings&batchId=154

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(findSum("012abc5a7abc09"))
	fmt.Println(findSumOther("012abc5a7abc09"))
}

func findSum(text string) int {
	var sb strings.Builder
	sum := 0
	for _, v := range text {
		if v >= 48 && v <= 57 {
			sb.WriteString(string(v))
		} else {
			if num, err := strconv.Atoi(sb.String()); err == nil {
				sum += num
			}
			sb.Reset()
		}
	}
	if num, err := strconv.Atoi(sb.String()); err == nil {
		sum += num
	}
	return sum
}

func findSumOther(text string) int {
	sum := 0
	num := 0
	for _, v := range text {
		if v >= 48 && v <= 57 {
			if op, err := strconv.Atoi(string(v)); err == nil {
				num = num*10 + op
			}

		} else {
			sum += num
			num = 0
		}
	}
	return sum + num
}
