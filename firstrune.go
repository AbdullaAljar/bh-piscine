package piscine

func FirstRune(s string) rune {
	rune1 := []rune(s)
	for _, letter := range rune1 {
		return letter
	}
	return 0
}
