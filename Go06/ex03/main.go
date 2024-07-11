package main

import (
	"os"
	"piscine"
)

func main() {
	args := os.Args[1:]
	argCount := piscine.countArgs(args)

	if argCount == 0 {
		piscine.Cat([]string{})
	} else {
		piscine.Cat(args)
	}
}
