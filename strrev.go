package piscine

import "strings"

func StrRev(s string) {
	fn := ""
	for i := len(strings.Split(s, "")); i >= 0; i-- {
		fn += strings.Split(s, "")[i]
	}
}
