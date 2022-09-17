package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	table := os.Args[1:]
	for i := 0; i < len(table); i++ {
		for j := i; j > 0 && table[j-1] > table[j]; j-- {
			table[j], table[j-1] = table[j-1], table[j]
		}
	}
	for _, c := range table {
		for _, r := range c {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
