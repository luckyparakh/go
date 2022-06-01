package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(anagramSearch("nos", "worldisone"))
	fmt.Println(anagramSearch("nose", "worldisone"))
	fmt.Println(anagramSearch("ide", "worldisone"))

	fmt.Println("---------------------Better")
	fmt.Println(anagramSearchSlidingWindow("nos", "worldisone"))
	fmt.Println(anagramSearchSlidingWindow("nose", "worldisone"))
	fmt.Println(anagramSearchSlidingWindow("ide", "worldisone"))
}
func anagramSearchSlidingWindow(pattern, text string) bool {
	arrText := [128]int{}
	arrPat := [128]int{}
	lenPat := len(pattern)
	for i := 0; i < lenPat; i++ {
		arrText[text[i]]++
		arrPat[pattern[i]]++
	}
	if arrCheck(arrText, arrPat) {
		return true
	}
	for i := lenPat; i < len(text); i++ {
		//fmt.Println(arrText)

		arrText[text[i]]++
		arrText[text[i-lenPat]]--
		if arrCheck(arrText, arrPat) {
			return true
		}
	}

	return false
}
func arrCheck(arr1, arr2 [128]int) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
func anagramSearch(pattern, text string) bool {
	//TC (n! + n)
	permSlice := findPerm("", pattern)
	fmt.Println(permSlice)
	for _, v := range permSlice {
		if index := strings.Index(text, v); index != -1 {
			return true
		}
	}
	return false
}

func findPerm(p, up string) []string {
	//TC(n!)
	if up == "" {
		return []string{p}
	}
	opSlice := []string{}
	for i := 0; i <= len(p); i++ {
		op := findPerm(p[0:i]+string(up[0])+p[i:], up[1:])
		for _, v := range op {
			opSlice = append(opSlice, v)
		}
	}
	return opSlice
}
