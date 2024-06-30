package piscine

import (
	"os"
	"ft"
)

func PrintParams() {
	args := os.Args[1:]// Skip the first argument which is the program name
	for _, arg := range args {
		for _, r := range arg {
			ft.PrintRune(r)
		}
		ft.PrintRune('\n')
	}
}