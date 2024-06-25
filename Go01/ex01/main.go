package main

import (
	"ex01/piscine"
	"fmt"
)

func main() {
	a := 0
	b := &a
	n := &b
	piscine.UltimatePointOne(&n)
	fmt.Println(a)
}
