package main

import "github.com/01-edu/z01"

// simple function
// print alaphabet using ASCII ref to optimize
func main() {
	for i := 97; i <= 122; i++ {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
}
