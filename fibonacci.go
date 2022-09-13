package piscine

func Fibonacci(index int) int {
	for i, a, b := 1, 0, 1; i <= index && index > 0; i++ {
		a, b = b, a+b
		if i == index {
			return a
		}
	}
	return -1
}
