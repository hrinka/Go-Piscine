package piscine

import (
	"os"
	"ft"
)

func RevParams() {
	args := os.Args[1:]
	for i := len(args) - 1; i >= 0; i-- {
		for _, r := range args[i] {
			ft.PrintRune(r)
		}
		ft.PrintRune('\n')
	}
}