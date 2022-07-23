// https://practice.geeksforgeeks.org/batch-problems/keypad-typing0119/0?track=DSASP-Strings&batchId=154
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(printNumber("adg"))
	fmt.Println(printNumber("geeksforgeeks"))
}

func printNumber(s string) string {
	var sb strings.Builder

	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'b' || s[i] == 'c' {
			sb.WriteString("2")
		} else if s[i] == 'd' || s[i] == 'e' || s[i] == 'f' {
			sb.WriteString("3")
		} else if s[i] == 'g' || s[i] == 'h' || s[i] == 'i' {
			sb.WriteString("4")
		} else if s[i] == 'j' || s[i] == 'k' || s[i] == 'l' {
			sb.WriteString("5")
		} else if s[i] == 'm' || s[i] == 'n' || s[i] == 'o' {
			sb.WriteString("6")
		} else if s[i] == 'p' || s[i] == 'q' || s[i] == 'r' || s[i] == 's' {
			sb.WriteString("7")
		} else if s[i] == 't' || s[i] == 'u' || s[i] == 'v' {
			sb.WriteString("8")
		} else if s[i] == 'w' || s[i] == 'x' || s[i] == 'y' || s[i] == 'z' {
			sb.WriteString("9")
		}
	}
	return sb.String()
}
