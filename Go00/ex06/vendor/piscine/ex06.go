package piscine

import (
	"fmt"
	"ft"
)

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	printComb([]rune{}, '0', n, false)
	fmt.Println()
}

func printComb(current []rune, start rune, n int, first bool) {
	if len(current) == n {
		if first {
			fmt.Print(", ")
		}
		for _, r := range current {
			ft.PrintRune(r)
		}
		first = true
		return
	}
	for r := start; r <= '9'; r++ {
		next := append(current, r)
		printComb(next, r+1, n, first)
		first = true
	}
}
