package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, s := range os.Args[1:] {
		for _, c := range s {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
