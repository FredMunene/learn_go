package piscine

func LastRune(s string) rune {
	str := []rune(s)
	lastIndex := len(str) - 1
	return str[lastIndex]
}
