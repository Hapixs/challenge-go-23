package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	a1 := 0
	a2 := 0

	if len(args) != 3 {
		return
	}

	if !(IsNumeric(args[0]) && IsNumeric(args[2])) {
		return
	}

	a1 = Atoi(args[0])
	a2 = Atoi(args[2])

	switch args[1] {
	case "*":
		Display(a1 * a2)
	case "/":
		if a2 == 0 {
			os.Stdout.WriteString("No division by 0" + "\n")
			return
		}
		Display(a1 / a2)
	case "%":
		if a2 == 0 {
			os.Stdout.WriteString("No modulo by 0" + "\n")
			return
		}
		Display(a1 % a2)
	case "+":
		Display(a1 + a2)
	case "-":
		Display(a1 - a2)
	}
}

func Display(i int) {
	str := Itoa(i)
	if str == "-" || str == "" || str == " " {
		return
	}
	os.Stdout.WriteString(str + "\n")
}

func IsNumeric(s string) bool {
	for _, c := range s {
		if !(c >= 48 && c <= 57) && c != '-' {
			return false
		}
	}
	return true
}

func Atoi(s string) int {
	negative := false     // Boolean permettant de savoir si le chiffre est negatif (modifier plus tard)
	total := 0            // chiffre total à retourner
	for _, c := range s { // Boucle qui passe sur chaques rune du string d'entrer
		if c >= 48 && c <= 57 { // Si le caractère actuel est bien un chiffre (ref ASCII)
			total = total*10 + int(c-48) // On multipli le chiffre par 10 pour pour ajouter la nouvelle unité trouvée
			continue                     // On passe le reste et on continue la boucle à l'iteration suivante
		} else if c == '-' || c == '+' && total == 0 { // si la rune est un - ou + et qui le total n'as pas encore été modifier (sinon cela veux dire qu'on est au milieu du string)
			negative = c == '-' // si la rune et - alors le chiffre et negatif sinon il est positif
			continue            // on continue à l'iteration suivante de la boucle
		}
		return 0 // Tout autre cas, on retourne 0, le string n'est pas un chiffre valide pour l'algorithme
	}
	//     dictionaire [  <clé> : <valeur>, ...       ] la valeur que l'on cherche en fonction de la clé (negative donc soit true soit false dans ce cas)
	return map[bool]int{true: total * -1, false: total}[negative] // Ici on retourn le chiffre en positif si (negative == false) ou en negatif si (negative == true)
}

func Itoa(i int) string {
	negative := i < 0
	if negative {
		i *= -1
	}
	str := ""
	for i > 0 {
		if i <= 9 && i > 0 {
			if negative {
				str = "-" + str
			}
			return string(rune(i+48)) + str
		}
		for y := 0; y <= 9; y++ {
			if (i-y)%10 == 0 {
				i = (i - y) / 10
				str = string(rune(y+48)) + str
			}
		}
	}
	if negative {
		str = "-" + str
	}
	return str
}
