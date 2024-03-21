package tests

func IterativeFactorial(nb int) int {
	result := nb
	if nb <= 0 {
		return 0
	}
	for i := 1; i < nb; i++ {
		result *= (nb - i)

	}
	return result
}
