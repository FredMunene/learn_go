package piscine

func Swap(a *int, b *int) {
	c := *a // Store the value pointed to by 'a' in 'c'
	*a = *b // Assign the value pointed to by 'b' to 'a'
	*b = c  // Assign the value stored in 'c' to 'b'
}
