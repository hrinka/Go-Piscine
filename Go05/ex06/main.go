package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.AppendRange(0, 16))
	fmt.Println(piscine.AppendRange(10, 5))
	fmt.Println(piscine.AppendRange(10, 10))
	fmt.Println(piscine.AppendRange(-1, 3))
}