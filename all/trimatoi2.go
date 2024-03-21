package piscine

func Trtoi(s string) int {
	sign := 0
	num := 0

	for _, value := range s {
		if value == '-' {
			sign = -1
		}
		if value >= '0' && value <= '9' {
			num = num*10 + int(value-'0')
		} else {
			num = 0
		}
	}
	return num * sign
}
