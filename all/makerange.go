package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	arry := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		arry[i] = min + i
	}
	return arry
}
