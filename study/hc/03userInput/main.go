package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("User Input")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter pizza rating:")
	input,_ := reader.ReadString('\n')
	fmt.Printf("Rating was %v.\n", input)
	fmt.Printf("Rating type was %T.\n", input)
}
