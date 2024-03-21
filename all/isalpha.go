package piscine

func IsAlpha(s string) bool {
	for _, value := range s {
		if value >= '0' && value <= '9' || (value >= 'a' && value <= 'z') || len(s) == 0 || (value >= 'A' && value <= 'Z') {
			continue
		} else {
			return false
		}
	}
	return true
}
