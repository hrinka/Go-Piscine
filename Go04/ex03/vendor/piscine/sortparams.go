package piscine

import (
	"os"
	"ft"
)

func SortParams() {
	args := os.Args[1:]

	n := len(args)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if args[j] > args[j+1] {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	for _, arg := range args {
		for _, r := range arg {
			ft.PrintRune(r)
		}
		ft.PrintRune('\n')
	}
}