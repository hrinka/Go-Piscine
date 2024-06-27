package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n; i++ {
		if isAlphaNumeric(runes[i]) {
			if i == 0 || !isAlphaNumeric(runes[i-1]) {
				if runes[i] >= 'a' && runes[i] <= 'z' {
					runes[i] = runes[i] - 32
				}
			} else {
				if runes[i] >= 'A' && runes[i] <= 'Z' {
					runes[i] = runes[i] + 32
				}
			}
		}
	}
	return string(runes)
}

func isAlphaNumeric(r rune) bool {
	return (r >= '0' && r <= '9') || (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}