package piscine

func UltimatePointOne(n ***int) {
	***n = 1 // On affect la valeur au pointeur qui lui même pointe vers un pointeur qui lui pointe vers une variable à modifier
}
