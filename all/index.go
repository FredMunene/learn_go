package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			found := true
			for j := 1; j < len(toFind); j++ {
				if s[i+j] != toFind[j] {
					found = false
					break
				}
			}
			if found {
				return i
			}

		}
	}
	return -1
}
