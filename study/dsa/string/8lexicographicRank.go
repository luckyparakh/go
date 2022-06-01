package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lexicographRank("bca"))

	fmt.Println(lexicographRank("string"))
	fmt.Println(lexicographRankBetter("bca"))
	fmt.Println(lexicographRankBetter("string"))
	fmt.Println(lexicographRankBest("bca"))
	fmt.Println(lexicographRankBest("string"))

}
func lexicographRankBest(s string) int {
	arr := [256]int{}
	fact := factorial(len(s))
	for _, v := range s {
		arr[v]++
	}
	//fmt.Println(arr)
	for i := 1; i < len(arr); i++ {
		arr[i] += arr[i-1]
	}
	//fmt.Println(arr)
	rank := 1
	for i := 0; i < len(s)-1; i++ {
		fact = fact / (len(s) - i)
		rank += arr[s[i]-1] * (fact)
		for k := int(s[i]); k < len(arr); k++ {
			arr[k]--
		}
	}
	return rank
}

func lexicographRankBetter(s string) int {
	rank := 1
	for i := 0; i < len(s)-1; i++ {
		smallChar := findSmallerChar(s[i], s[i+1:])
		//fmt.Printf("%v %v %v \n",s[i],s[i+1:],smallChar)
		rank = rank + smallChar*factorial(len(s)-1-i)
	}
	return rank
}

func factorial(i int) int {
	f := 1
	for j := 2; j <= i; j++ {
		f *= j
	}
	return f
}

func findSmallerChar(b byte, s string) int {
	c := 0
	for i := 0; i < len(s); i++ {

		if s[i] < b {
			c++
		}
	}
	return c
}

func lexicographRank(s string) int {
	//TC(n! + n + n)
	permSlice := findPerm("", s)
	sort.Slice(permSlice, func(i, j int) bool {
		return permSlice[i] < permSlice[j]
	})
	//fmt.Println(permSlice)
	for i, v := range permSlice {
		if v == s {
			return i + 1
		}
	}
	return -1
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
