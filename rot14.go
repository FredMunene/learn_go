package tests

var result string

func Rot14(s string) string {
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			result += string((char-'a'+14)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			result += string((char-'A'+14)%26 + 'A')
		} else {
			result += string(char)
		}
	}
	return result
}
