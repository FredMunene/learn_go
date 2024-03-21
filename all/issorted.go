package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	urefu := 0
	for i := range a {
		urefu = i + 1
	}
	ascending := true
	descending := true

	for i := 1; i < urefu; i++ {
		if !(f(a[i-1], a[i]) >= 0) {
			descending = false
		}
	}
	for i := 1; i < urefu; i++ {
		if !(f(a[i-1], a[i]) <= 0) {
			ascending = false
		}
	}
	return ascending || descending
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
