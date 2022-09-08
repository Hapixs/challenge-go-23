package piscine

import "github.com/01-edu/z01"

// Hard Mod
func PrintCombN(n int) {
	PrintCombNRec(n - 1) // sinon c'est la d !
}

// Fonction avec boucle principal qui contrôle le nombre total de groupe de chiffres à afficher
func PrintCombNRec(n int) {
	for ref := 0; ref <= 9-n; ref++ {
		lastLoop := ref == 9-n
		runes := []rune{rune(ref + 48)} // Création d'un tableau avec le premier chiffre à afficher
		if n > 0 {
			loop(runes, ref, n-1, lastLoop) // Appelle de récursion si necessaire
		} else {
			displayRunes(runes, lastLoop)
		}
	}
	z01.PrintRune('\n')
}

// Fonction récursive
// Contrôle le nombre de chiffres à afficher par groupe moin le premier
// pref pour previus_ref
// intAmount pour le nombre récursions restante à faire
// lastLoop pour is this the f**king last loop man ? (utiliser uniquement pour print ou non le ', ')
func loop(runes []rune, pref int, intAmount int, lastLoop bool) {
	for ref := pref + 1; ref <= 9-intAmount; ref++ {
		tmpRune := runes // Copy du tableau de rune d'entrer pour ne pas le modifier (tmpRune = temporaire rune)
		tmpRune = append(tmpRune, rune(ref+48))
		if intAmount > 0 { // Si il reste des loop à faire
			loop(tmpRune, ref, intAmount-1, lastLoop && (ref == 9-intAmount)) // Début de récursion
		} else { // Sinon on affiche la puré
			displayRunes(tmpRune, lastLoop)
		}
	}
}

// Fonction qui print les runes depuis un array
func displayRunes(runes []rune, lastLoop bool) {
	for _, r := range runes { // Et la on passe trquillou sur l'array final
		z01.PrintRune(r)
	}
	if !(lastLoop) {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}

// bingo ca marche
