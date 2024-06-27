package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsNumeric("0123456789"))
	fmt.Println(piscine.IsNumeric("01,02,03"))
	fmt.Println(piscine.IsNumeric("012345678x"))

}
