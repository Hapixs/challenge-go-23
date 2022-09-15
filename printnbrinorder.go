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
	for i := 0; i < 9; i++ {
		if n <= 9 {
			tab = append(tab, i)
			return tab
		}
		if (n-i)%10 == 0 {
			tab = append(tab, i)
			return defineTab((n - i) / 10)
		}
	}
	return tab
}
