package piscine

func SortIntegerRable(table []int) {
	n := len(table)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}
