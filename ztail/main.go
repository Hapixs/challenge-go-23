package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		return
	}

	if args[0] != "-c" || !IsNumeric(args[1]) {
		return
	}

	c := Atoi(args[1])

	osExit := false
	for _, a := range args[2:] {
		content, err := os.ReadFile(a)
		if err != nil {
			os.Stdout.WriteString(err.Error() + "\n")
			continue
		}
		ct := string(content)
		os.Stdout.WriteString("\n==> " + a + " <==\n")
		if len([]rune(ct)) < c {
			os.Stdout.WriteString(ct)
			osExit = true
			continue
		}
		for i := c; i > 0; i-- {
			os.Stdout.WriteString(string(ct[len([]rune(ct))-i]))
		}
	}
	if osExit {
		os.Exit(1)
	}
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

func IsNumeric(s string) bool {
	for _, c := range s {
		if !(c >= 48 && c <= 57) && c != '-' {
			return false
		}
	}
	return true
}
