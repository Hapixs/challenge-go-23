package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, c := range os.Args {
		for _, r := range c[2:] {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
