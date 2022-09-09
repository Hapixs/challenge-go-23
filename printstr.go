package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, c := range s { // Boucle pour passer sur chaque rune (c) du string (s)
		z01.PrintRune(c) // Affichage de la rune actuelle
	}
}
