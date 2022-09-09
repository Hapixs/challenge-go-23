package piscine

func StrRev(s string) string {
	str := ""             // String qui se verra être retourner
	for _, c := range s { // Boucle qui passe sur chaque caractères du string d'entrer (s)
		str = string(c) + str // Ajout du string (str) au caractère actuel, à l'inverse ajouter le caractère au string ne ferais qu'ajouter dans le bon sens les caractères
	}
	return str // reversed ingenering ;D
}
