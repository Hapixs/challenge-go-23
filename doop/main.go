package main

import (
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 3 || (!IsNumeric(args[0]) && IsNumeric(args[2])) {
		return
	}

	a1 := Atoi(args[0])
	a2 := Atoi(args[2])

	if a1 >= 9223372036800000000 || a2 >= 9223372036800000000 {
		return
	}

	operator := args[1]

	if (operator == "/" || operator == "%") && a2 <= 0 {
		os.Stdout.WriteString("No " + map[string]string{"/": "division", "%": "modulo"}[operator] + " by 0\n")
		return
	}

	for k, v := range map[string]int{"*": a1 * a2, "/": a1 / a2, "%": a1 % a2, "+": a1 + a2, "-": a1 - a2} {
		if operator == k {
			Display(v)
			return
		}
	}
}

func Display(i int) {
	if i < 9223372036854775806 {
		os.Stdout.WriteString(Itoa(i) + "\n")
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
	str := ""
	negative := false
	if i < 0 {
		i *= -1
		negative = true
	}
	for i > 0 {
		if i <= 9 && i > 0 {
			return string(rune(i+48)) + str
		}
		y = i%10
		i = (i - y) / 10
		str = string(rune(y+48)) + str
	}
	return map[bool]string{true: "-" + str, false: str}[negative]
}
