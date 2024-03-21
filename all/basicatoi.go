package piscine

func BasicAtoi(s string) int {
	var result int
	for _, digit := range s {
		result = result*10 + int(digit-'0')
	}
	return result
}
