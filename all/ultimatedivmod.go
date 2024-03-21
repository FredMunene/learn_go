package piscine

func UltimateDivMod(a *int, b *int) {
	d := *a / *b
	e := *a % *b
	*a = d
	*b = e
}
