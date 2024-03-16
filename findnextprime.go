package piscine

func FindNextPrime(nb int) int {
	if !IsPrime(nb) {
		nb++
	}
	return nb
}
