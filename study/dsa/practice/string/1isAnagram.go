package main

import "fmt"

// https://practice.geeksforgeeks.org/problems/anagram-1587115620/0/?track=DSASP-Strings&batchId=154
func main() {
	fmt.Println(isAnagram("abcd", "bcda"))
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	text := [128]byte{}

	for i := 0; i < len(a); i++ {
		text[a[i]]++
		text[b[i]]--
	}
	for j := 0; j < len(text); j++ {
		if text[j] != 0 {
			return false
		}
	}
	return true
}
