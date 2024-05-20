package piscine

func UltimateDivMod(a *int, b *int) {
	tempA := *a
	*a = tempA / (*b)
	*b = tempA % (*b)
}
