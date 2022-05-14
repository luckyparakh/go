//Find Duplicate lines

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var count = make(map[string]int)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		//dup2(args)
		dup3(args)
	} else {
		dup()
	}
}

func dup() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { //press ctl+d to stop while reading from os.stdin
		count[input.Text()]++
	}
	countAndPrint()
}
func dup2(args []string) {
	for _, arg := range args {
		fle, err := os.Open(arg)
		defer fle.Close()
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error while reading:%v", arg)
			continue
		}
		input := bufio.NewScanner(fle)
		for input.Scan() { // streaming data
			count[input.Text()]++
		}
		countAndPrint()
	}
}
func dup3(args []string) {
	for _, arg := range args {
		var tmpcount = make(map[string]int)
		// Read file once, for big file it need more memory
		bytedata, err := os.ReadFile(arg) // can also use ioutil.ReadFile
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error while reading:%v", arg)
			continue
		}
		for _, line := range strings.Split(string(bytedata), "\n") {
			tmpcount[line]++
		}
		fmt.Fprintf(os.Stdout,"Filename:%s\n",arg)
		countAndPrint1(tmpcount)
	}
}
func countAndPrint() {
	fmt.Println("Count and Print")
	for index, val := range count {
		if val > 1 {
			fmt.Printf("%v:%d\n", index, val)
		}
	}
}

func countAndPrint1(tmpcount map[string]int) {
	fmt.Println("Count and Print")
	for index, val := range tmpcount {
		if val > 1 {
			fmt.Printf("%v:%d\n", index, val)
		}
	}
}
