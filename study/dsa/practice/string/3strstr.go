// https://practice.geeksforgeeks.org/problems/implement-strstr/0/?track=DSASP-Strings&batchId=154

package main

import "fmt"

func main() {
	fmt.Println(strtsr("GeeksforGeeks", "fr"))
	fmt.Println(strtsr("GeeksforGeeks", "for"))
}

func strtsr(text, pat string) int {
	lenPat := len(pat)
	for i := 0; i < len(text)-lenPat; i++ {
		if text[i:i+lenPat] == pat {
			return i
		}
	}
	return -1
}
