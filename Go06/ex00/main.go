package main

import (
	"ft"
	"os"
)

const (
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
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func countArgs(args []string) int {
	count := 0
	for range args {
		count++
	}
	return count
}

func main() {
	lengthOfArg := countArgs(os.Args) - 1
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
