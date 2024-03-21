package piscine

func ConcatParams(args []string) string {
	arryLength := len(args) * 2        // the lenth of arry slice we need
	arry := make([]string, arryLength) // our slice
	index := 0
	for i := 0; i < len(args); i++ { // iterate args
		arry[index] = args[i] // assigns
		index++               // increments once

		if i != len(args)-1 { // assigns newline but not to last index
			arry[index] = "\n"
			index++ // increments once
		}
	}
	var result string
	for _, value := range arry { // iterates through arry(string slice )
		result += value // appends slice to result
	}
	return result
}
