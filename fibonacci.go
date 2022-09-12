package piscine

func Fibonacci(index int) int {
	if index <= 0 {
		return -1
	}
	suite := []int{}
	for i := 0; i < index+1; i++ {
		if i == 0 {
			suite = append(suite, 0)
			continue
		}
		if i == 1 {
			suite = append(suite, 1)
			continue
		}
		suite = append(suite, suite[i-1]+suite[i-2])
	}
	return suite[index]
}
