package piscine

func UltimateDivMod(a *int, b *int) {
	if *b != 0 {
		div := *a / *b
		mod := *a % *b
		*a = div
		*b = mod
	}
}
