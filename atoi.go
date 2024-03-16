package piscine

//func accepts number as string as input and outputs int
/*    Write a function that simulates the behaviour of the Atoi function in Go. Atoi transforms a number represented as a string in a number represented as an int.
Atoi returns 0 if the string is not considered as a valid number. For this exercise non-valid string chains will be tested. Some will contain non-digits characters.
For this exercise the handling of the signs + or - does have to be taken into account.
This function will only have to return the int. For this exercise the error result of Atoi is not required.
*/
func Atoi(s string) int {
	sign := 1
	var result int
	for index, value := range s {
		if index == 0 && value == '-' { //sets sign for integer
			sign = -1
		}
		if index >= 1 && !(value >= '0' && value <= '9') { // ensures string is valid
			return 0
		}
	}
	for _, value := range s {
		if value >= '0' && value <= '9' {
			result = result*10 + int(value-'0') // converts string to int and assigns it to result
		} else {
			result = 0
		}
	}
	return result * sign //updates the sign
}

//returns 0 if number is invalid
// non-valid strings will be tested
// handling of + or - signs
//
