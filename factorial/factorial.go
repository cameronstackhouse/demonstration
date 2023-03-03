package factorial

func Factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total = total * i
	}

	return total

}
