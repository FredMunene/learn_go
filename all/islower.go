package piscine

func IsLower(s string) bool {
	for _, value := range s {
		if value >= 'a' && value <= 'z' {
			continue
		} else {
			return false
		}
	}
	return true
}
