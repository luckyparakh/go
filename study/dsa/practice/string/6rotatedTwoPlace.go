//https://practice.geeksforgeeks.org/batch-problems/check-if-string-is-rotated-by-two-places-1587115620/0/?track=DSASP-Strings&batchId=154

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isRotated("amazon", "azonam"))
	fmt.Println(isRotated("amazon", "onamaz"))
}
func isRotated(actual, rotated string) bool {
	rotatedBy := 2
	var acw strings.Builder
	var cw strings.Builder
	lenA := len(actual)
	
	//ACW Check
	acw.WriteString(actual[rotatedBy:])
	for i := 0; i < rotatedBy; i++ {
		acw.WriteString(string(actual[i]))
	}

	//CW
	for i := rotatedBy; i > 0; i-- {
		cw.WriteString(string(actual[lenA-i]))
	}
	cw.WriteString(actual[:lenA-rotatedBy])
	
	//Compare
	if acw.String() == rotated {
		return true
	}else if cw.String() == rotated {
		return true
	}
	return false
}
