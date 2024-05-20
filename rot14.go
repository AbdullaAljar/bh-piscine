package piscine

func rot14(b rune) rune {
	if (b >= 'A' && b <= 'M') || (b >= 'a' && b <= 'm') {
		return b + 14
	}

	if (b >= 'N' && b <= 'Z') || (b >= 'n' && b <= 'z') {
		return b - 12
	}

	return b
}

func Rot14(s string) string {
	result := ""
	for _, letter := range s {
		result += string(rot14(letter))
	}
	return result
}
