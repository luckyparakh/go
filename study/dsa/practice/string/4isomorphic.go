package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("egg", "abc"))
	fmt.Println(isIsomorphic("aabbc", "xyxyz"))
	fmt.Println(isIsomorphic("baba", "badc"))
	fmt.Println(isIsomorphic("badc", "baba"))
}

func isIsomorphic(s string, t string) bool {
	sCount := [128]byte{}
	tCount := [128]byte{}

	for i := 0; i < len(s); i++ {
		sCount[s[i]] = t[i]
		
	}
	for j := 0; j < len(s); j++ {
		if sCount[s[j]] != t[j] {
			return false
		}
	}
	// Needed because of following case
	// 	fmt.Println(isIsomorphic("baba", "badc"))
	//  fmt.Println(isIsomorphic("badc", "baba"))
	for i := 0; i < len(t); i++ {
		tCount[t[i]] = s[i]
	}
	for j := 0; j < len(t); j++ {
		if tCount[t[j]] != s[j] {
			return false
		}
	}
	return true
}
