package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(baseName("abc/dcg.go"))
	fmt.Println(baseName("abc"))
	fmt.Println(baseName("abc.def/ghi.go"))
	fmt.Println("BaseName1")
	fmt.Println(baseName1("abc/dcg.go"))
	fmt.Println(baseName1("abc"))
	fmt.Println(baseName1("abc.def/ghi.go"))
}

func baseName(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s=s[:i]
			break
		}
	}
	return s
}

func baseName1(s string) string {
	slash:=strings.LastIndex(s,"/")
	s=s[slash+1:]
	if dot:=strings.LastIndex(s,".");dot>=0{
		s=s[:dot]
	}
	return s
}