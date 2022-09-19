package piscine

func SplitWhiteSpaces(str string) []string {
	for i, s := range str {
		if s == ' ' || s == '\n' || s == '\t' {
			if len(str[:i]) > 0 {
				return append([]string{str[:i]}, SplitWhiteSpaces(str[i+1:])...)
			} else {
				return append([]string{}, SplitWhiteSpaces(str[i+1:])...)
			}
		}
	}
	return []string{str} /*

		[]string{"MIX", "V,c6)^$os", "5-7o|im29BrG?", "CsNe6U<^1+w1", " 2]I5ch>#g,V1{", "/46UIg&RxuIe", " \\oOM)~8$Vcw]Z", "K0l6y\\JFTpP&V", "SIZ.-&0@}%i(\""} instead of
		[]string{"MIX", "V,c6)^$os", "5-7o|im29BrG?", "CsNe6U<^1+w1", "2]I5ch>#g,V1{", "/46UIg&RxuIe", "\\oOM)~8$Vcw]Z", "K0l6y\\JFTpP&V", "SIZ.-&0@}%i(\""}
	*/
}
