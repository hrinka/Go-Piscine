package piscine

import (
	"fmt"
	"ft"
)

func PrintComb2() {
	for i := 0; i <= 99; i++ {
		for j := i + 1; j <= 99; j++ {
			if !(i == 0 && j == 1) {
				fmt.Print(", ")
			}
			ft.PrintRune(rune('0' + i/10))
			ft.PrintRune(rune('0' + i%10))
			fmt.Print(" ")
			ft.PrintRune(rune('0' + j/10))
			ft.PrintRune(rune('0' + j%10))

		}
	}
	fmt.Println()
}
