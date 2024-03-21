package piscine

// function that, for an int slice, applies a function on each element of that slice
func ForEach(f func(int), a []int) {
	for _, value := range a {
		f(value)
	}
}
