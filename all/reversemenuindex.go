package piscine

func ReverseMenuIndex(menu []string) []string {
	reversed := make([]string, len(menu))
	for i := 0; i < len(menu); i++ {
		reversed[len(menu)-1-i] = menu[i]
	}
	return reversed
}
