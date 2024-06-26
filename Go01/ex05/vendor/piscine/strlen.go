package piscine

func Strlen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
