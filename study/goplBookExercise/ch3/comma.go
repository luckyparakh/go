//The task is to take a str ing represent ation of an integer, such as "12345", and insert commas every three places, as in "12,345".
package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	
	intsToString([]int{1, 2, 3}) // Byte buffer example
	fmt.Println(comma("1023000"))
	fmt.Println(comma("1023000"))
	fmt.Println(commaFloatSign("-1234567890.1234"))
	fmt.Println(commaFloatSign("1234567890.1234"))
	fmt.Println(commaFloatSign("+1234567890.1234"))
}

func intsToString(num []int) {
	b := bytes.Buffer{}
	b.WriteByte('[')
	for i, v := range num {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(strconv.Itoa(v))
		//OR
		//fmt.Fprintf(&b,"%d",v)
	}
	b.WriteByte(']')
	fmt.Println(b.String())
}
func commaRecur(num string) string {
	s := len(num)
	if s < 3 {
		return num
	}
	return comma(num[:s-3]) + "," + num[s-3:]
}

func commaReverse(num string) string {
	b := bytes.Buffer{}
	for i := 1; i <= len(num); i++ {
		b.WriteByte(num[len(num)-i])
		if i%3 == 0 {
			b.WriteString(",")
		}
	}
	return b.String()
}

func comma(num string) string {
	b := bytes.Buffer{}
	numLen := len(num)
	pre := numLen % 3
	b.WriteString(num[:pre])
	//b.WriteString(",")

	for i := pre; i < len(num); i += 3 {
		b.WriteString(",")
		b.WriteString(num[i : i+3])
		// if i+3 != numLen {
		// 	//to avoid comma at last
		// 	b.WriteString(",")
		// }

	}
	return b.String()
}

func commaFloatSign(num string) string{
	sign:=""
	if num[:1]=="+" || num[:1]=="-"{
		sign=num[:1]
		num=num[1:]
	}
	unsigned:=strings.Split(num,".")
	if sign==""{
		return comma(unsigned[0])+"."+unsigned[1]
	}
	return sign+comma(unsigned[0])+"."+unsigned[1]
}
