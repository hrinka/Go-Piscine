package piscine

import (
	"os"
	"ft"
)

func PrintProgramName() {
	printProgramName := os.Args[0]
	for _, r := range printProgramName {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}