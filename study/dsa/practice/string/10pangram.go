//https://practice.geeksforgeeks.org/batch-problems/pangram-checking-1587115620/0/?track=DSASP-Strings&batchId=154

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(checkPangram("Bawds jog, flick quartz, vex nymph"))
	fmt.Println(checkPangram("Bawds jog, flick quartz, vex nymaa"))
}

func checkPangram(text string) bool {
	if len(text) < 26 {
		return false
	}
	text = strings.ToLower(text)
	letterCount := [128]byte{}
	for _, v := range text {
		letterCount[v]++
	}
	// fmt.Println(letterCount)
	for i := 97; i <= 122; i++ {
		if letterCount[i] == 0 {
			return false
		}
	}
	return true
}
