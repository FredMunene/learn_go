package piscine

// func removes zero-values from slice
func Compact(ptr *[]string) int { // slice
	count := 0 // count for non-zero elems

	for i := 0; i < len(*ptr); {
		if (*ptr)[i] != "" {
			count++
			i++
		} else {
			*ptr = append((*ptr)[:i], (*ptr)[i+1:]...)
		}
	}
	return count
}
