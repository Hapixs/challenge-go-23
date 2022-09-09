package piscine

func BasicAtoi(s string) int {
	// Algo de base pour ce cas précis
	/*total := 0
	for _, c := range s {
		if c >= 48 && c <= 57 {
			total = total*10 + int(c-48)
		}
	}
	return total*/
	return Atoi(s) // Appelle la fonction Atoi(s) dans le fichier atoi.go
}
