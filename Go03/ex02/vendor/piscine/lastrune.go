package piscine

func LastRune(s string) rune {
	if len(s) == 0 {
		return 0
	}
	var lastRune rune
	for _, r := range s {
		lastRune = r
	}
	return lastRune
}
