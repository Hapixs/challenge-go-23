package piscine

func Split(str, sep string) []string {
	for len(str) > 0 {
		i := Index(str, sep)
		if i < 0 {
			return []string{str}
		}
		return append([]string{str[:i]}, Split(str[i+len([]rune(sep)):], sep)...)
	}
	return []string{str}
	/*
		[]string{"xUnQLgHALgc1d", ">abc<>]UqmQFzHetoL4R", ">abc<>]R9oaFnlolSxfd", ">abc<>]y0JmAME3YK6VP", ">abc<>]PUUck2cdh24eG", ">abc<>]CkqoJyLfR1i0d", ">abc<>]TgRjXxVWqNPTZ", ">abc<>]bYd8nKWnrmVVg"} instead of
		[]string{"xUnQLgHALgc1d", "UqmQFzHetoL4R", "R9oaFnlolSxfd", "y0JmAME3YK6VP", "PUUck2cdh24eG", "CkqoJyLfR1i0d", "TgRjXxVWqNPTZ", "bYd8nKWnrmVVg"}
	*/
}
