package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsPrintable("hello\n"))
	fmt.Println(piscine.IsPrintable("hello!"))
	fmt.Println(piscine.IsPrintable("abcdffF~\t"))
}
