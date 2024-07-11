package main

import (
	"os"
	"piscine"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		piscine.Cat([]string{})
	} else {
		piscine.Cat(args)
	}
}
