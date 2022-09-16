package piscine

func AlphaCount(s string) int {
	return map[bool]int{true: map[bool]int{true: 1 + AlphaCount(string(s[1:])), false: 0}[IsAlpha(s)], false: 0}[StrLen(s) > 0]
}
