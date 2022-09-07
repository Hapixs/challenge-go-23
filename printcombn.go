package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	PrintCombNRec(n - 1)
}

func PrintCombNRec(remain int) {
	for ref := 0; ref <= 9-remain; ref++ {
		loop([]rune{rune(ref + 48)}, ref, remain-1)
	}
	z01.PrintRune('\n')
}

func loop(runes []rune, pref int, intAmount int) {
	for ref := pref + 1; ref <= 9-intAmount; ref++ {
		tmpRune := runes
		tmpRune = append(tmpRune, rune(ref+48))
		if intAmount > 0 {
			loop(tmpRune, ref, intAmount-1)
		} else {
			for _, r := range tmpRune {
				z01.PrintRune(r)
			}
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
}
