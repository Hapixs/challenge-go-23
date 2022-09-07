package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	PrintCombNRec(n - 1)
}

func PrintCombNRec(remain int) {
	for ref := 0; ref <= 9-remain; ref++ {
		lastLoop := ref == 9-remain
		runes := []rune{rune(ref + 48)}
		if remain > 0 {
			loop(runes, ref, remain-1, lastLoop)
		} else {
			displayRunes(runes, lastLoop)
		}
	}
	z01.PrintRune('\n')
}

func loop(runes []rune, pref int, intAmount int, lastLoop bool) {
	for ref := pref + 1; ref <= 9-intAmount; ref++ {
		tmpRune := runes
		tmpRune = append(tmpRune, rune(ref+48))
		if intAmount > 0 {
			loop(tmpRune, ref, intAmount-1, lastLoop && (ref == 9-intAmount))
		} else {
			displayRunes(tmpRune, lastLoop)
		}
	}
}

func displayRunes(runes []rune, lastLoop bool) {
	for _, r := range runes {
		z01.PrintRune(r)
	}
	if !(lastLoop) {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}
