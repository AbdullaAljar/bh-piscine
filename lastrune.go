package piscine

func LastRune(s string) rune {
	rune1 := []rune(s)
	length := len(rune1)
	if length > 0 {
		return rune1[length-1]
	}
	return 0
}
