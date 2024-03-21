package piscine

func Unmatch(a []int) []int {
	var oddNumber []int
	for _, digit := range a {
		if IsOdd(oddNumber, digit) {
			oddNumber = append(oddNumber, digit)
		}
	}
	return oddNumber
}

func IsOdd(a []int, num int) bool {
	for _, val := range a {
		if val == num {
			return false
		}
	}
	return true
}
