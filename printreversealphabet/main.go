package main

import "github.com/01-edu/z01"

// Fonction principal
func main() {
	for i := 122; i >= 97; i-- { // Même principe que 'printalphabet' mais cette fois dans le sens inverse;
		z01.PrintRune(rune(i)) // On part de 122 et on dessend jusqu'à 97
	}
	z01.PrintRune('\n')
}
