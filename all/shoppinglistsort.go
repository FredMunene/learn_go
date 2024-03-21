package piscine

func ShoppingListSort(slice []string) []string {
	for f := 0; f < len(slice)-1; f++ {
		for d := f + 1; d < len(slice); d++ {
			if len(slice[f]) > len(slice[d]) {
				slice[f], slice[d] = slice[d], slice[f]
			}
		}
	}
	return slice
}
