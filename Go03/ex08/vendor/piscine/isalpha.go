package piscine

func IsAlpha(s string) bool {
	for _, r := range s {
		if !isAlphaNumeric(r) {
			return false
		}
	}
	return true
}

func isAlphaNumeric(r rune) bool {
	return (r >= '0' && r <= '9') || (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}