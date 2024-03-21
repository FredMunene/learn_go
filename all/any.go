package piscine

func Any(f func(string) bool, a []string) bool {
	N := func(s string) bool {
		for _, char := range s {
			if char < '0' || char > '9' {
				return false
			}
		}
		return true
	}
	for _, str := range a {
		if f(str) || N(str) {
			return true
		}
	}
	return false
}
