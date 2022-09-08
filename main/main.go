package main

import (
	"fmt"
	"piscine"
)

func main() {
	// piscine.PrintComb()
	// piscine.PrintComb2()

	s := []int{10, -10, 100, 4, 0, -4}
	piscine.SortIntegerTable(s)
	fmt.Println(s)

}
