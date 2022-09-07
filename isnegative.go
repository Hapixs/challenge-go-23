package piscine // Package piscine, qui peut être importer depuis les autres packages ce qui leurs donnerais accès aux fonctions de 'piscine'

import "github.com/01-edu/z01"

//Fonction isNegative
// Cette fonction affiche si un nombre 'nb' pris en entrer est négatif ou non
// Le format est tel que:
// T = Oui donc négatif (T pour true)
// F = Non donc positif (F pour false)
func IsNegative(nb int) {
	// Ici,
	// map[bool]rune{...} permet de créer un 'dictionnaire' qui prend en compte un boolean en clé et une rune en valeur
	// On récupere la clé attendu grâce à [nb < 0] cette expresion est un boolean
	// Example si nb=10;  [10 < 0] et donc égale à 'false' on demande donc au 'dictionnaire' de nous donner la valeur de 'false')
	z01.PrintRune(map[bool]rune{true: 'T', false: 'F'}[nb < 0])
	z01.PrintRune('\n') // Petit retoure à la ligne des familles
}
