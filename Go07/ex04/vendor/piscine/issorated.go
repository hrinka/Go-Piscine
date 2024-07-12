package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if lenSlice(a) == 0 {
		return true
	}
	for i := 0; i < lenSlice(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			return false
		}
	}
	return true
}

func lenSlice(a []int) int {
	length := 0
	for range a {
		length++
	}
	return length
}
