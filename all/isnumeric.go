package piscine

func IsNumeric(s string) bool {
	for _, value := range s {
		if value >= '0' && value <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
