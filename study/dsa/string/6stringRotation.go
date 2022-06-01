package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strRotate("abcd", "cdab"))
	fmt.Println(strRotate("abcd", "abc"))
	fmt.Println(strRotate("aaaba", "aaaab"))
	fmt.Println(strRotate("aaaba", "aaaaa"))
	fmt.Println("-----------------Better")
	fmt.Println(strRotateBetter("abcd", "cdab"))
	fmt.Println(strRotateBetter("abcd", "abc"))
	fmt.Println(strRotateBetter("aaaba", "aaaab"))
	fmt.Println(strRotateBetter("aaaba", "aaaaa"))
}

func strRotate(str, rotateStr string) bool {
	//TC: o(n2) because of tr comparision
	//Sc: o(1)
	if len(str) != len(rotateStr) {
		return false
	}
	strCopy := str

	for i := 0; i < len(strCopy); i++ {
		var sb strings.Builder
		sb.WriteString(strCopy[1:])
		sb.WriteString(string(strCopy[0]))
		strCopy = sb.String()
		//strCopy=strCopy[1:]+string(strCopy[0])
		if strCopy == rotateStr {
			return true
		}
	}
	return false
}

func strRotateBetter(str, rotateStr string) bool {
	//TC:o(n) SC:o(n)
	if len(str) != len(rotateStr) {
		return false
	}
	str = str + str
	if op := strings.LastIndex(str, rotateStr); op == -1 {
		return false
	}
	return true
}
