package piscine

import "ft"

func ForEach(f func(int), a []int) {
	i := 0
	for {
		if i >= sliceLen(a) {
			break
		}
		f(a[i])
		i++
	}
}

func PrintNbr(n int) {
	if n == 0 {
		ft.PrintRune('0')
		return
	}

	if n < 0 {
		ft.PrintRune('-')
		n = -n
	}
	var digits [10]rune
	digitCount := 0

	for n > 0 {
		digits[digitCount] = rune('0' + n%10)
		n /= 10
		digitCount++
	}
	for i := digitCount - 1; i >= 0; i-- {
		ft.PrintRune(digits[i])
	}
}

func sliceLen(a []int) int {
	length := 0
	for range a {
		length++
	}
	return length

}
