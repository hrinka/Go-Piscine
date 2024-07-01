package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.MakeRange(10, 5))
	fmt.Println(piscine.MakeRange(10, 10))
	fmt.Println(piscine.MakeRange(-1, 3))
	fmt.Println(piscine.MakeRange(0, 16))
}
