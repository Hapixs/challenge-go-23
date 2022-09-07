package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for ref := 0; ref <= 7; ref++ {
		for sc := ref + 1; sc <= 8; sc++ {
			for th := sc + 1; th <= 9; th++ {
				z01.PrintRune(rune(ref + 48))
				z01.PrintRune(rune(sc + 48))
				z01.PrintRune(rune(th + 48))
				if !(ref == 7 && sc == 8 && th == 9) {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
