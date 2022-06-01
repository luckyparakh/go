package main

import "fmt"

func main() {
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d)

	const (
		aa = 1 * iota
		bb
		cc
		dd
	)
	fmt.Println(aa, bb, cc, dd)

	const (
		_ = 1 << (10 * iota)
		KiB
		MiB
		GiB
	)
	fmt.Println(KiB, MiB, GiB)

	const (
		flag1 = 1 << iota
		flag2
		flag3
	)
	fmt.Printf("%b %b %b\n", flag1, flag2, flag3)

	const (
		// Iota has some limits like it’s not possible to generate the more familiar powers of 1000 (KB, MB, and so on) 
		// because there is no exponentiation operator
		KB = 1000
		MB = KB * 1000
		GB = MB * 1000
	)
	fmt.Printf("%T %[1]d\n",KB)
	fmt.Printf("%T %[1]d\n",MB)
}
