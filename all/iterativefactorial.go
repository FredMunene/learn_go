package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb < 0 || nb > 21 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		for i := 1; i < nb+1; i++ {
			result *= i
		}
		return result
	}
}
