package recursive

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return n * factorial(n-1)
}
