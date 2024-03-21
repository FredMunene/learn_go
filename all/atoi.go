package piscine

// func accepts number as string as input and outputs int
func Atoi(s string) int {
	sign := 1
	var result int
	for index, value := range s {
		if index == 0 && value == '-' {
			sign = -1
		}
		if index >= 1 && !(value >= '0' && value <= '9') {
			return 0
		}
	}
	for _, value := range s {
		if value >= '0' && value <= '9' {
			result = result*10 + int(value-'0')
		} else {
			result = 0
		}
	}
	return result * sign
}

//returns 0 if number is invalid
// non-valid strings will be tested
// handling of + or - signs
//
