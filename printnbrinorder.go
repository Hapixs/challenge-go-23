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
	for n > 0 {
		for i := 0; i < 10; i++ {
			if n < 10 {
				tab = append(tab, n)
				n = 0
			} else if (n-i)%10 == 0 {
				tab = append(tab, i)
				n = (n - i) / 10
			}
		}
	}
	return tab
}
