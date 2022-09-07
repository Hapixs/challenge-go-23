package main

import "github.com/01-edu/z01"

// simple function
// reverse print alaphabet using ASCII ref to optimize
func main() {
	for i := 48; i <= 57; i++ {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
}
