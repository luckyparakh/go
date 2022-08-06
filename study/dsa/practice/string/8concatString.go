// https://practice.geeksforgeeks.org/batch-problems/remove-common-characters-and-concatenate-1587115621/0/?track=DSASP-Strings&batchId=154
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(concatenatedString("aacdb", "gafd"))
}

func concatenatedString(s1, s2 string) string {
	var sb strings.Builder
	s1Count := [128]byte{}
	for _, v := range s1 {
		if s1Count[v] == 0 {
			s1Count[v]++
		}
	}
	for _, v := range s2 {
		if s1Count[v] == 0 {
			sb.WriteString(string(v))
		} else {
			s1Count[v]++
		}
	}
	for k, v := range s1Count {
		if v == 1 {
			//fmt.Println(string('a' - 96 + k))
			sb.WriteString(string('a' - 96 + k))
		}
	}
	return sb.String()
}
