package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	table := os.Args[1:]
	for i := 0; i < len(table); i++ { // Boucle qui passe sur chaque valeur du tableau
		for j := i; j > 0 && table[j-1] > table[j]; j-- { // Boucle qui repasse sur chaque valeur du tableau dans le sens inverse de i, tant que la valeur précedente de la liste (table[j-1]) est inferieur à celle actuel (table[j])
			table[j], table[j-1] = table[j-1], table[j] // Echange des deux valeur si table[j-1] > table[j]
		}
	}
	for _, c := range table {
		for _, r := range c {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
