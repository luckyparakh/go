// https://practice.geeksforgeeks.org/batch-problems/check-if-a-string-is-isogram-or-not-1587115620/0/?track=DSASP-Strings&batchId=154

package main

import "fmt"

func main() {
	fmt.Println(isIsogram("machine"))
	fmt.Println(isIsogram("amazon"))
	fmt.Println(isIsogram("geeks"))
}

func isIsogram(text string) bool {
	charSet := make([]int, 128)
	for i := 0; i < len(text); i++ {
		if charSet[text[i]] == 0 {
			charSet[text[i]] = 1
		} else if charSet[text[i]] == 1 {
			return false
		}
	}
	return true
}
