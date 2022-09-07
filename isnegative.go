package student

import "github.com/01-edu/z01"

func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}

func IsNegative(nb int) {
	z01.PrintRune(map[bool]rune{true: rune(84), false: rune(70)}[nb < 0])
	z01.PrintRune('\n')
}
