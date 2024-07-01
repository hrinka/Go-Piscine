package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.SplitWhiteSpaces("Hello how are you?"))
	fmt.Println(piscine.SplitWhiteSpaces("Hellohowareyou    ?"))
	fmt.Println(piscine.SplitWhiteSpaces("Hello\thow\nare you?"))
	fmt.Println(piscine.SplitWhiteSpaces(" Hello\t how\n are you? "))
}
