package piscine

import "github.com/01-edu/z01"

// Fonction principal avec le nombre à afficher 'n'
func PrintNbr(n int) {
	// On commence par verifier si il est négatif
	// Les nombres négatifs trop grands ne peuvent pas être passer en positif, en revanche certains nombre négatifs transferable en positif font buger si il ne sont pas mis en positifs
	// Donc on check si on peut le passer en positif sinon osef ca marchera quand même
	negative := n < 0
	if negative && (n*-1) != 0 { // n*-1 renvois n positif, si le resultat est 0 alors on ne peux pas passer n en positif
		n *= -1 // ici on passe n en positif t'as vu :D
	}

	// Première étape,
	// Si le chiffre est négatif, alors on affiche un '-' en premier.
	if negative {
		z01.PrintRune('-')
	}

	rec(n) // Début de la méchante récursion
}

// Fonction de récursion
// Elle fait un tas de calcules shelou
func rec(n int) { // n étant le nombre (négatif / positif)
	// Pour ce faire nous allons trouver chaques chiffres en dessous de la dizaine
	// Une boucle boucle (émércé) de 0 à 9
	for i := int(0); i <= 9; i++ {
		if n <= 9 && n >= -9 { // Si le nombre n est plus petit que 10 (ou plus grand que -10 si il est negatif) alors pas besoin de ce prendre la tête on à juste à l'afficher en rune depuis l'ASCII
			z01.PrintRune(rune(n + 48))
			return
		}
		if (n-i)%10 == 0 { // Si le chiffre 'n' moins 'i' divisé par 10 donne un reste de 0 à ca division c'est donc que le chiffre des unités est 'i'
			rec((n - i) / 10)           // On rebalance donc notre chiffre n dans la fonction de récursion en enlevant ce qu'on vient de trouver (i) et on le divise par 10 pour atteindre la prochaine unité à trouver
			z01.PrintRune(rune(i + 48)) // Ensuite on print la valeur de i (L'affichage aura lieu une fois la récursion terminer)
			return
		}
	}
}
