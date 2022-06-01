package main

import "fmt"

func main() {
	fmt.Println(isAnagram("aaa", "bbbddd"))
	fmt.Println(isAnagram("aba", "baa"))
	fmt.Println(isAnagram("abb", "baa"))
	
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	aFreq := make(map[rune]int)
	bFreq := make(map[rune]int)

	for _, v := range a {
		aFreq[v]++
	}

	for _, v := range b {
		bFreq[v]++
	}
	for k, v := range aFreq {
		if bFreq[k] != v {
			return false
		}
	}
	return true
}
