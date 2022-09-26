package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	printStr(map[bool]string{true: "I have an even number of arguments", false: "I have an odd number of arguments"}[len(os.Args[1:])%2 == 0] + "\n")
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
