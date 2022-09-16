package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	final := []rune{}
	toInsert := []rune{}
	order := false

	for _, s := range args {
		order = order || s == "-o" || s == "--order"
		if StrIndex(s, "--insert=")+StrIndex(s, "-i=") > 0 {
			toInsert = append(toInsert, []rune(s[StrIndex(s, "=")+1:])...)
			continue
		}
		if s != "-o" && s != "--order" {
			final = append(final, []rune(s)...)
		}
	}

	final = append(final, []rune(toInsert)...)

	for _, s := range map[bool][]rune{true: SortRuneTable(final), false: final}[order] {
		if s != ' ' || order {
			z01.PrintRune(s)
		}
	}
	if len(final) > 0 {
		z01.PrintRune(' ')
	}
	z01.PrintRune('\n')
}

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("	 This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")
}

func SortRuneTable(t []rune) []rune {
	table := []rune(t)
	for i := 0; i < len(table); i++ {
		for j := i; j > 0 && table[j-1] > table[j]; j-- {
			table[j], table[j-1] = table[j-1], table[j]
		}
	}
	return t
}

func StrIndex(s, find string) int {
	SizeS := StrLen(s)
	SizeF := StrLen(find)
	for i := 0; i < SizeS-SizeF; i++ {
		if s[i:i+SizeF] == find {
			return i
		}
	}
	return -1
}

func StrLen(s string) int {
	return len([]rune(s))
}
