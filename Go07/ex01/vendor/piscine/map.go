package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, sliceLen(a))
	for i, v := range a {
		result[i] = f(v)
	}
	return result
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= n/i; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func sliceLen(a []int) int {
	length := 0
	for range a {
		length++
	}
	return length

}
