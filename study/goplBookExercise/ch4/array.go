package main

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	INR
	EUR
	GBP
)

func main() {
	b := [3]int{20, 30, 40}
	fmt.Println(b)
	a := [...]int{2, 3, 4}
	fmt.Println(a)

	// var aa [3]int=[3]int{1,2,3}
	// var bb [2]int=[2]int{1,2}
	// fmt.Println(aa==bb) //invalid
	currArr:=[...]Currency{USD:1,INR:73,EUR:2,GBP:1}
	fmt.Println(currArr,currArr[USD])
	zeroArr:=[...]int{9:-1}
	fmt.Println(zeroArr)
	shaArr:= sha256.Sum256([]byte{'X'})
	fmt.Printf("%x\n%[1]T",shaArr)
}
