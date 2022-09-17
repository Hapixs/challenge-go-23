package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 || contains(args[0], []string{"--help", "-h"}) {
		printHelp()
		return
	}
	final := []rune{}
	toInsert := []rune{}
	order := false
	for _, s := range args {
		order = order || contains(s, []string{"-o", "--order"})
		if contains(s, []string{"--insert=", "-i="}) {
			toInsert = append(toInsert, []rune(s[StrIndex(s, "=")+1:])...)
		} else if s != "-o" && s != "--order" {
			final = append(final, []rune(s)...)
		}
	}
	final = append(final, []rune(toInsert)...)
	if order {
		SortRuneTable(final)
	}
	for _, s := range final {
		if s != ' ' || order {
			z01.PrintRune(s)
		}
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

func SortRuneTable(table []rune) {
	for i := 0; i < len(table); i++ {
		for j := i; j > 0 && table[j-1] > table[j]; j-- {
			table[j], table[j-1] = table[j-1], table[j]
		}
	}
}

func contains(reference string, find []string) bool {
	for _, c := range find {
		if StrIndex(reference, c) >= 0 {
			return true
		}
	}
	return false
}

func StrIndex(s, find string) int {
	SizeS := len([]rune(s))
	SizeF := len([]rune(find))
	for i := 0; i <= SizeS-SizeF; i++ {
		if s[i:i+SizeF] == find {
			return i
		}
	}
	return -1
}
