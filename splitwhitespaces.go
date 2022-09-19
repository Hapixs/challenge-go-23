package piscine

func SplitWhiteSpaces(str string) []string {
	for i, s := range str {
		if s == ' ' || s == '\n' || s == '\t' && len(str) > 0 {
			if len(str[:i]) > 0 {
				return append([]string{str[:i]}, SplitWhiteSpaces(str[i+1:])...)
			} else {
				return append([]string{}, SplitWhiteSpaces(str[i+1:])...)
			}
		}
	}
	if len(str) > 0 {
		return []string{str}
	}
	return []string{}
	/*
		[]string{",tN&3E>i.r#.c", "U?s)Uu>6xD|&m", "@gtP~}X0ZNPRQ", "}#{3kA4#lPNw+", "w|>Ou^RiLZw]?", "/gJc.U'0(RfP=", "*eywKok@QE8Nm", "2?n.Zz?Q+|'N", ""} instead of
		[]string{",tN&3E>i.r#.c", "U?s)Uu>6xD|&m", "@gtP~}X0ZNPRQ", "}#{3kA4#lPNw+", "w|>Ou^RiLZw]?", "/gJc.U'0(RfP=", "*eywKok@QE8Nm", "2?n.Zz?Q+|'N"}
	*/
}
