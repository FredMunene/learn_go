package piscine

func BasicAtoi2(s string) int {
	var result int
	for _, digit := range s {
		if digit < '0' || digit > '9' {
			return 0 // If any character is not a digit, return 0
		}
		result = result*10 + int(digit-'0')
	}
	return result
}
