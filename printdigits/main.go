package main

import "github.com/01-edu/z01"

// Fonction principal
func main() {
	for i := 48; i <= 57; i++ { // Même principe que 'printalphabet' mais avec le code ASCII des chiffres.
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
}
