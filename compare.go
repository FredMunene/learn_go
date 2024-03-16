package piscine

func Compare(a, b string) int {
	str1 := []rune(a)
	str2 := []rune(b)

	minLen := len(str1)
	if len(str2) < minLen {
		minLen = len(str2)
	}

	for i := 0; i < minLen; i++ {
		if str1[i] < str2[i] {
			return -1
		} else if str1[i] > str2[i] {
			return 1
		}
	}
	if len(str1) < len(str2) {
		return -1
	} else if len(str1) > len(str2) {
		return 1
	}
	return 0
}

/*package piscine

//Compare returns an integer comparing two strings lexicographically. The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func Compare(a, b string) int {
	minLen := len(a)
	if len(b) < len(a) {
		minLen = len(b)
	}
	for i := 0; i < minLen; i++ {
		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}
	// overhauls above considering length of string
	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}
	return 0
}
*/
