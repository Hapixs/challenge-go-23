package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for first := 0; first < 9; first++ {
		for second := 0; second < 9; second++ {
			for second <= first && first != 9 {
				second++
			}
			for third := 0; third < 9; third++ {
				for third <= second && first != 9 && second != 9 {
					third++
				}
				z01.PrintRune(rune(first + 48))
				z01.PrintRune(rune(second + 48))
				z01.PrintRune(rune(third + 48))
				z01.PrintRune(',')
				z01.PrintRune(' ')

			}
		}
	}
}
