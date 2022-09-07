package main // Package principale du dossier '/printalphabet'

import "github.com/01-edu/z01" // import

// Fonction principal
func main() {
	for i := 97; i <= 122; i++ { // Boucle allant de 97 à 122
		z01.PrintRune(rune(i)) // Dans la table ASCII (google is your friend), les caractère de 97 à 122 corresponde à ceux de l'alphabet latin en minuscule et dans l'ordre
	} // il suffit ensuite pour chaqu'un de le 'print' sans espace ni retour à la ligne

	z01.PrintRune('\n') // Ajout d'un retour à la ligne
}
