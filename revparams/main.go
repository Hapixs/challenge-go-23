package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	displaySlice(os.Args[1:])
}

func displaySlice(s []string) {
	if len(s) <= 0 {
		return
	}
	displaySlice(s[1:])
	for _, c := range s[0] {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
