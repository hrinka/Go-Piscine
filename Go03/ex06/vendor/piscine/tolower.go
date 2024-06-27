package piscine

func ToLower(s string) string {
	result := ""
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			result += string(r + 'a' - 'A')
		} else {
			result += string(r)
		}
	}
	return result
}