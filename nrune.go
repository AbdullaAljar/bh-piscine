package piscine

func NRune(s string, n int) rune {
	rune1 := []rune(s)
	length := n
	if length > 0 && length <= len(rune1) {
		return rune1[length-1]
	}
	return 0
}
