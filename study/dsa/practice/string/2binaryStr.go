// https://practice.geeksforgeeks.org/problems/binary-string-1587115620/0/?track=DSASP-Strings&batchId=154

package main

import "fmt"

func main() {
	fmt.Println(binarySubstring("10011"))
	fmt.Println(binarySubstring("1111"))
	fmt.Println(binarySubstring("1000"))
	fmt.Println("Better")
	fmt.Println(binarySubstringBetter("10011"))
	fmt.Println(binarySubstringBetter("1111"))
	fmt.Println(binarySubstringBetter("1000"))
}

func binarySubstring(text string) int {
	lenT := len(text)
	count := 0
	for i := 0; i < lenT; i++ {
		if string(text[i]) == "1" {
			for j := i + 1; j < lenT; j++ {
				if string(text[j]) == "1" {
					count++
				}
			}
		}
	}
	return count
}

func binarySubstringBetter(text string) int {
	// Count number of ones
	lenT := len(text)
	countOne := 0
	for i := 0; i < lenT; i++ {
		if string(text[i]) == "1" {
			countOne++
		}
	}
	return (countOne * (countOne - 1)) / 2
}
