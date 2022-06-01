package main

import (
	"fmt"
	"strings"
)

func main() {
	reverseWord("I   love   coding   ")
}

func reverseWord(ss string) {
	ss = strings.Trim(ss, " ")
	s := []byte(ss) // as string as immutable
	start := 0
	//for end, v := range s {
	for end := 0; end < len(s); {
		if s[end] == ' ' {
			reverse(s, start, end-1)
			// to handle spaces
			for s[end+1] == ' ' {
				end++
			}
			start = end + 1
		}
		end++
	}
	// fmt.Println(string(s))
	reverse(s, start, len(s)-1) // this is needed as last word will not reverse due to absense of space
	// fmt.Println(string(s))
	reverse(s, 0, len(s)-1)
	fmt.Println(string(s))
}

func reverse(s []byte, start, end int) {

	for start <= end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
