package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	i := 0
	j := 0

	k := 0
	l := 0
	for ref := 0; ref < 9; ref++ {
		for sref := 0; sref <= 9; sref++ {
			for tref := 0; tref < 9; tref++ {
				j = sref
				if tref+1 < sref || tref+1 == sref {
					continue
				}
				l = tref + 1
				z01.PrintRune(rune(i + 48))
				z01.PrintRune(rune(j + 48))
				z01.PrintRune(' ')
				z01.PrintRune(rune(k + 48))
				z01.PrintRune(rune(l + 48))
				if !(i == 8 && j == 9 && k == 9 && l == 9) {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
			i = ref + 1
			k = ref + 1
		}
	}
	z01.PrintRune('\n')
}
