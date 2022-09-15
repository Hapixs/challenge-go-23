package piscine

func PrintNbrInOrder(n int) {
	tab := defineTab(n)
	SortIntegerTable(tab)
	for _, i := range tab {
		PrintNbr(i)
	}
}

func defineTab(n int) []int {
	tab := []int{}
	for i := int(0); i <= 9; i++ {
		if n <= 9 && n >= -9 { // Si le nombre n est plus petit que 10 (ou plus grand que -10 si il est negatif) alors pas besoin de ce prendre la tête on à juste à l'afficher en rune depuis l'ASCII
			tab = append(tab, n)
			return tab
		}
		if (n-i)%10 == 0 {
			tab = append(tab, n)
			tab = append(tab, defineTab((n-i)/10)...) // On rebalance donc notre chiffre n dans la fonction de récursion en enlevant ce qu'on vient de trouver (i) et on le divise par 10 pour atteindre la prochaine unité à trouver
			return tab
		}
	}
	return tab
}
