package piscine

import "github.com/01-edu/z01"

func PrintWordStables(s []string) {
	for _, c := range s {
		for _, st := range c {
			z01.PrintRune(st)
		}
		z01.PrintRune('\n')
	}
}
