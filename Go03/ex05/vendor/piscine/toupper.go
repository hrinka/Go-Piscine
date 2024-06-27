package piscine

func ToUpper(s string) string {
	result := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			result += string(r - 'a' + 'A')
		} else {
			result += string(r)
		}
	}
	return result
}