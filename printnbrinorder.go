package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune(rune(48))
		return
	}
	tab := defineTab(n)
	SortIntegerTable(tab)
	for _, i := range tab {
		z01.PrintRune(rune(i + 48))
	}
}

func defineTab(n int) []int {
	tab := []int{}
	for n > 9 {
		for i := 0; i <= 9; i++ {
			if (n-i)%10 == 0 {
				tab = append(tab, i)
				n = (n - i) / 10
				if i >= 6 {
					break
				}
			}
		}
	}
	if n > 0 {
		return append(tab, n)
	}
	return tab
}
