package factorial

func Factorial(n int) int {
	total := 1
	for n; n > 0; n-- {
		total = total * n
	}

	return total

}
