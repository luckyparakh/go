package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		str := strings.Split(scan.Text(), old)
		fmt.Println(strings.Join(str, new))
	}
}
