//Go interfaces are implemented implicitly
//if a type contains all the methods declared in the interface.
package main

import "fmt"

type vowelFinder interface {
	FindVowel() []rune
}

type MyString string

func (m MyString) FindVowel() []rune {
	var vowel []rune
	for _, r := range m {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			vowel = append(vowel, r)
		}
	}
	return vowel
}

func main() {
	name := MyString("Rishi Parakh")
	var v vowelFinder
	v = name
	fmt.Printf("%c\n", v.FindVowel())
}
