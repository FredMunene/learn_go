package tests

//returna elem of a slice that does not have a correspodning pair

func Unmatch(a []int) int {
	// Re-order the array
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	// Check for unpaired element
	for i := 0; i < len(a)-1; i += 2 {
		if a[i] != a[i+1] {
			return a[i]
		}
	}
	// If no unpaired element found, return the last element
	return a[len(a)-1]
}
