package piscine

func ToUpper(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			char = char - 32
		}
		result = result + string(char)
	}
	return result
}
