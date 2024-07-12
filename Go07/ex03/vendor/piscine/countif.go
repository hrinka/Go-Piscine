package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for i := 0; i < lenStringSlice(tab); i++ {
		if f(tab[i]) {
			count++
		}
	}
	return count
}

func IsNumeric(s string) bool {
	for i := 0; i < lenString(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func lenStringSlice(slice []string) int {
	length := 0
	for range slice {
		length++
	}
	return length
}

func lenString(s string) int {
	length := 0
	for range s {
		length++
	}
	return length
}
