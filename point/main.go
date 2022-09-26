package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{0, 0}

	setPoint(points)

	str := "x = " + Itoa(points.x) + ", y = " + Itoa(points.y) + "\n"
	for _, c := range str {
		z01.PrintRune(c)
	}
}

func Itoa(i int) string {
	str := ""
	for i > 0 {
		if i <= 9 && i > 0 {
			return string(rune(i+48)) + str
		}
		for y := 0; y <= 9; y++ {
			if (i-y)%10 == 0 {
				i = (i - y) / 10
				str = string(rune(y+48)) + str
			}
		}
	}
	return str
}
