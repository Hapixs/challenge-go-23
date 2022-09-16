package piscine

func BasicJoin(strings []string) string {
	if len(strings) < 1 {
		return ""
	}
	return strings[0] + map[bool]string{true: BasicJoin(strings[1:]), false: ""}[len(strings) > 1]
}
