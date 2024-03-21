package tests

func RecursiveFactorial(nb int) int {
	result := nb
	if nb < 0 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	result *= RecursiveFactorial(nb - 1)
	return result
}
