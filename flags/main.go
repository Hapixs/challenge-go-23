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

	toOrder := []rune{}
	order := false
	for index := 0; index < len(args); index++ {
		s := args[index]
		if !order {
			order = s == "-o" || s == "--order"
		}

		i := StrIndex(s, "--insert=") + StrIndex(s, "-i=")
		if i+2 > 0 {
			toInsert := ""
			vIndex := StrIndex(s, "=")
			for _, c := range s[vIndex+1:] {
				toInsert += string(c)
			}
			toInsert = args[index+1] + toInsert
			toOrder = append(toOrder, []rune(toInsert)...)
			index++
			continue
		}
		if s != "-o" && s != "--order" {
			toOrder = append(toOrder, []rune(s)...)
			continue
		}
	}

	if order {
		SortRuneTable(toOrder)
	}
	for _, s := range toOrder {
		if s != ' ' {
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
	for i := 0; i < len(table); i++ { // Boucle qui passe sur chaque valeur du tableau
		for j := i; j > 0 && table[j-1] > table[j]; j-- { // Boucle qui repasse sur chaque valeur du tableau dans le sens inverse de i, tant que la valeur précedente de la liste (table[j-1]) est inferieur à celle actuel (table[j])
			table[j], table[j-1] = table[j-1], table[j] // Echange des deux valeur si table[j-1] > table[j]
		}
	}
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

func Index(s []string, find string) int {
	SizeS := len(s)
	SizeF := StrLen(find)
	for index, str := range s {
		for i := 0; i < SizeS-SizeF; i++ {
			if str[i:i+SizeF] == find {
				return index
			}
		}
	}
	return -1
}

func StrLen(s string) int {
	return len([]rune(s))
}
