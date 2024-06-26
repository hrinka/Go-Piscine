package main

import (
	"ft"
	"os"
)

const (
	yes     = true
	no      = false
	EvenMsg = "I have an even number of arguments"
	OddMsg  = "I have an odd number of arguments"
)

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func isEven(nbr int) bool {
	if even(nbr) {
		return yes
	} else {
		return no
	}
}

func even(nbr int) bool {
	return nbr%2 == 0
}

func main() {
	lengthOfArg := len(os.Args) - 1
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
