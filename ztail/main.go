package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 || args[0] != "-c" || !IsNumeric(args[1]) {
		return
	}
	c := Atoi(args[1])
	for i, a := range args[2:] {
		content, err := os.ReadFile(a)
		if err != nil {
			os.Stdout.WriteString(err.Error() + "\n")
			defer os.Exit(1)
			continue
		}
		if i > 0 && i < len(args)-1 {
			os.Stdout.WriteString("\n")
		}
		ct := string(content)
		os.Stdout.WriteString("==> " + a + " <==\n")
		if len(ct) <= c {
			os.Stdout.WriteString(ct)
			continue
		}
		for i := c; i > 0; i-- {
			os.Stdout.WriteString(string(ct[len([]rune(ct))-i]))
		}
	}
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

func IsNumeric(s string) bool {
	for _, c := range s {
		if !(c >= 48 && c <= 57) && c != '-' {
			return false
		}
	}
	return true
}
