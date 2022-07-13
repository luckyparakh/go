package ma

import (
	"fmt"
	"unicode"
)

func main() {

}

func IsPalin(str string) bool {
	var letters []rune
	for _, r := range str {
		if unicode.IsLetter(r) {
			letters = append(letters, r)
		}
	}
	for i, _ := range letters {
		fmt.Println(i)
	}
}
