package piscine

func LoafOfBread(str string) string {
	result := ""
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	word := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			word += string(str[i])
		} else {
			if len(word) >= 5 {
				result += word[:5] + " "
				word = ""
			} else {
				result += word + " "
				word = ""
			}
		}

		if len(word) == 5 {
			result += word + " "
			word = ""
		}
	}

	if len(word) >= 5 {
		result += word[:5]
	} else {
		result += word
	}
	result += "\n"
	return result
}
