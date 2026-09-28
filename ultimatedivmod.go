package piscine

func UltimateDivMod(a *int, b *int) {
	olda := *a
	*a = *a / *b
	*b = olda % *b
}
