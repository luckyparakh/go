package main

import "fmt"

func main() {
	patternSearch("abcd", "abcdabc")
	patternSearch("aa", "aaaaa")
	patternSearch("bc", "aaaaa")
}

func patternSearch(pattern, search string) {
	//TC: O(n-m+1)*m where m is length of pattern
	start := 0
	end := len(pattern)
	patternFound:=false
	for end <= len(search) {
		if search[start:end] == pattern { // considering this comparision will take m time to compare
			patternFound=true
			fmt.Println(start)
		}
		start++
		end++
	}
	if !patternFound{
		fmt.Println("-1")
	}
}
