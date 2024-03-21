package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	// given a string, count amountof times each item appears in it and give a summary
	wordCount := make(map[string]int)
	start := 0
	// Iterate over each character in the input string
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' || i == len(str)-1 {
			// Extract the word from the input string
			var word string
			if i == len(str)-1 {
				word = str[start:]
			} else {
				word = str[start:i]
			}
			wordCount[word]++
			start = i + 1
		}
	}
	return wordCount
}
