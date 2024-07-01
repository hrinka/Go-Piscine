package piscine

func Split(s, sep string) []string {
	var result []string
	start := 0
	for {
		index := findIndex(s, sep, start)
		if index == -1 {
			result = append(result, s[start:])
			break
		}
		result = append(result, s[start:index])
		start = index + len(sep)
	}
	return result
}

func findIndex(s, sep string, start int) int {
	n := len(sep)
	for i := start; i <= len(s)-n; i++ {
		if s[i:i+n] == sep {
			return i
		}
	}
	return -1
}
