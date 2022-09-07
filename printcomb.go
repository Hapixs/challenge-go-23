package piscine

import "github.com/01-edu/z01"

// Fonction PrintComb
// Niveau recrue
func PrintComb() {
	// Ici on utilise 3 boucle différentes mais liés
	// En imagine aussi que
	// - la variable fr (pour first) correspond au premier de nos chiffres
	// - la variable sc (pour second) correspond au deuxième de nos chiffres
	// - la variable th (pour third) correspond au troisième de nos chiffres
	// On sait que nous devons afficher trois chiffres, ce pour quoi les trois variables.
	// On sait aussi qu'un chiffre ne peux pas être supperieur ni égale à sont prochain
	// A l'inverse un chiffre ne peux pas être inferrieur nis égale à sont précédent
	for fi := 0; fi <= 7; fi++ { // On part de zéro pour aller jusqu'as 7 car dans la logique le chiffre max est 9 et ne peux donc pas être le première (inferieur au suivant)
		for sc := fi + 1; sc <= 8; sc++ { // On par de la valleur actuelle + 1 du chiffre précédent, ce qui assure d'être suppérieur non égale
			for th := sc + 1; th <= 9; th++ { // tout pareil
				z01.PrintRune(rune(fi + 48)) // On affiche les trois valeurs
				z01.PrintRune(rune(sc + 48))
				z01.PrintRune(rune(th + 48))
				// Ici on verifie que ca ne sera pas le dérnier print en prenant en compte la valeur max possible de chaqu'un des chiffres
				// Si chaqu'un on leurs valeurs alors on ne print pas
				// (Remarque, le '!' au début de la condition inverse sont résultat)
				if !(fi == 7 && sc == 8 && th == 9) {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n') // Retoure à la ligne ;) On oubli pas
}
