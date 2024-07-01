package main

import (
	"piscine"
)

func main() {
	a := piscine.SplitWhiteSpaces("HELLO HOW ARE YOU?")
	piscine.PrintWordsTables(a)
}
