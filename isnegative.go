package student

import "github.com/01-edu/z01"

func main() {
	isNegative(1)
	isNegative(0)
	isNegative(-1)
}

func isNegative(nb int) {
	z01.PrintRune(map[bool]rune{true: rune(84), false: rune(70)}[nb < 0])
	z01.PrintRune('\n')
}
