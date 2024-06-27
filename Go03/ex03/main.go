package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Index("Hello!", "o"))
	fmt.Println(piscine.Index("Salut!", "lut"))
	fmt.Println(piscine.Index("Bye!", "h0c"))
}
