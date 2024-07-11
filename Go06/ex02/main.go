package main

import (
	"fmt"
	"os"
	"piscine"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	piscine.DisplayFile(os.Args[1])
}
