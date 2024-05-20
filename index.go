package piscine

func Index(s string, toFind string) int {
	sentence := []rune(s)
	toFindRunes := []rune(toFind)

	for i := 0; i <= len(sentence)-1; i++ {
		find := true

		for j := 0; j <= len(toFindRunes)-1; j++ {
			if sentence[i+j] != toFindRunes[j] {
				find = false
				break
			}
		}

		if find {
			return i
		}
	}

	return -1
}
