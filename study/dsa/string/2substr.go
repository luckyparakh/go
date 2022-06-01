package main

import "fmt"

func main() {
	subStrIsPresent("bc", "abcd")
	subStrIsPresent("dbc", "abcd")
	fmt.Println("---------------Better Version")
	subStrIsPresentBetter("bc", "abcd")
	subStrIsPresentBetter("dbc", "abcd")
}

func subStrIsPresent(subStr, str string) {
	//TC: O(2^n) + n
	subStrSlice := substr("", str)
	//fmt.Println(subStrSlice)
	for _, v := range subStrSlice {
		if subStr == v {
			fmt.Println("Present")
			return
		}
	}
	fmt.Println("Not Present")
}

func subStrIsPresentBetter(subStr, str string) {
	//Todo write recursive
	//TC:O(n)
	i := 0
	j := 0
	for i < len(subStr) && j < len(str) {
		if subStr[i] == str[j] {
			i++
			j++
		} else {
			j++
		}
	}
	if i == len(subStr) {
		fmt.Println("Present")
		return
	}
	fmt.Println("Not Present")
}

func substr(p, up string) []string {
	//TC : O(2^n)
	//SC:O(n)
	if up == "" {
		return []string{p}
	}
	tmp1 := substr(p+up[:1], up[1:])
	tmp2 := substr(p, up[1:])
	for _, v := range tmp2 {
		tmp1 = append(tmp1, v)
	}
	return tmp1
}
