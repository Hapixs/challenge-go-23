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
	for n > 9 {
		for i := 0; i <= 9; i++ {
			if (n-i)%10 == 0 {
				tab = append(tab, i)
				n = (n - i) / 10
			}
		}
	}
	return append(tab, n)
}
