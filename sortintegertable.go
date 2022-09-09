package piscine

// Methode: tri par incersion
func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ { // Boucle qui passe sur chaque valeur du tableau
		for j := i; j > 0 && table[j-1] > table[j]; j-- { // Boucle qui repasse sur chaque valeur du tableau dans le sens inverse de i, tant que la valeur précedente de la liste (table[j-1]) est inferieur à celle actuel (table[j])
			table[j], table[j-1] = table[j-1], table[j] // Echange des deux valeur si table[j-1] > table[j]
		}
	}
}
