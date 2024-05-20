package piscine

func AlphaCount(s string) int {
	sentence := []rune(s)
	count := 0

	for _, letter := range sentence {
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
			count++
		}
	}
	return count
}
