package piscine

func FindPrevPrime(nb int) int {
	if !IsPrime(nb) {
		nb--
	}
	return nb
}
