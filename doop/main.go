package main

import (
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 3 || (!IsNumeric(args[0]) && IsNumeric(args[2])) || len(args[0]) > 8 || len(args[2]) > 8 {
		return
	}

	a1 := Atoi(args[0])
	a2 := Atoi(args[2])

	operator := args[1]

	if (operator == "/" || operator == "%") && a2 == 0 {
		os.Stdout.WriteString("No " + map[string]string{"/": "division", "%": "modulo"}[operator] + " by 0\n")
	} else if val, ok := map[string]int{"*": a1 * a2, "/": a1 / a2, "%": a1 % a2, "+": a1 + a2, "-": a1 - a2}[operator]; ok {
		os.Stdout.WriteString(Itoa(val) + "\n")
	}
}

func IsNumeric(s string) bool {
	for _, c := range s {
		if !(c >= 48 && c <= 57) && c != '-' {
			return false
		}
	}
	return true
}

func Atoi(s string) int {
	negative := false
	total := 0
	for _, c := range s {
		if c >= 48 && c <= 57 {
			total = total*10 + int(c-48)
			continue
		} else if c == '-' || c == '+' && total == 0 {
			negative = c == '-'
			continue
		}
		return 0
	}
	return map[bool]int{true: total * -1, false: total}[negative]
}

func Itoa(i int) string {
	if i == 0 {
		return "0"
	}
	str := ""
	negative := false
	if i < 0 {
		i *= -1
		negative = true
	}
	for i > 0 {
		y := i % 10
		i = (i - y) / 10
		str = string(rune(y+48)) + str
	}
	return map[bool]string{true: "-" + str, false: str}[negative]
}
