package piscine

func ToLower(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			char = char + 32
		}
		result = result + string(char)
	}
	return result
}
