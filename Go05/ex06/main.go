package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "HelloHAhowHAreHAyou?"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
}
