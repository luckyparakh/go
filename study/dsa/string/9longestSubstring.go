//find distinct longest substring
package main

import "fmt"

func main() {
	longestDistinctSubstr("abcdabc")
	longestDistinctSubstr("abcdabcde")
	longestDistinctSubstr("aaa")
	longestDistinctSubstr("")
	longestDistinctSubstr("abaacdbab")
}

func longestDistinctSubstr(s string) {
	// arr := [128]int{}
	hm:=make(map[byte]int)
	start := 0
	maxLen := 0
	i:=0
	for  ;i < len(s); i++ {
		if _,ok:=hm[s[i]];!ok{
			hm[s[i]]=i
		}else{
			if i-start > maxLen {
				maxLen = i - start
			}
			start=i
			hm=make(map[byte]int)
			hm[s[i]]=i
		}
		// if arr[s[i]] == 0 {
		// 	arr[s[i]] = 1
		// }else{
		// 	if i-start > maxLen {
		// 		maxLen = i - start
		// 	}
		// 	start = i
		// 	for _, v := range s {
		// 		arr[v] = 0
		// 	}
		// 	arr[s[i]] = 1
		// }
	}
	if i-start>maxLen{
		maxLen=i-start
	}
	fmt.Println(maxLen)
}
