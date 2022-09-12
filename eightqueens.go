package piscine

import (
	"github.com/01-edu/z01"
)

type Point struct {
	x int
	y int
}

func EightQueens() {
	Solve(8)
}

var results = make([][]Point, 0)

func Solve(n int) {
	for col := 0; col < n; col++ {
		start := Point{x: col, y: 0}
		current := make([]Point, 0)
		Recurse(start, current, n)
	}
	for _, result := range results {
		for i, rs := range result {
			if i == 0 && rs.x == 0 {
				continue
			}
			z01.PrintRune(rune(rs.x + 48))
		}
		z01.PrintRune('\n')
	}
}

func Recurse(point Point, current []Point, n int) {
	if CanPlace(point, current) {
		current = append(current, point)
		if len(current) == n {
			c := make([]Point, n)
			for i, point := range current {
				c[i] = point
			}
			results = append(results, c)
		} else {
			for col := 0; col < n; col++ {
				for row := point.y; row < n; row++ {
					nextStart := Point{x: col, y: row}
					Recurse(nextStart, current, n)

				}

			}
		}
	}
}
func CanPlace(target Point, board []Point) bool {
	for _, point := range board {
		if CanAttack(point, target) {
			return false
		}
	}
	return true
}

func CanAttack(a, b Point) bool {
	//fmt.Print(a, b)
	answer := a.x == b.x || a.y == b.y || float64(a.y-b.y) == float64(a.x-b.x)
	//fmt.Print(answer)
	return answer
}
