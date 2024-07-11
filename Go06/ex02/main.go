package main

import (
	"fmt"
	"os"
	"piscine"
)

func countArgs(args []string) int {
	count := 0
	for range args {
		count++
	}
	return count

}

func main() {
	argCount := countArgs(os.Args)

	if argCount == 1 {
		fmt.Println("File name missing")
		return
	} else if argCount > 2 {
		fmt.Println("Too many arguments")
		return
	}
	piscine.DisplayFile(os.Args[1])
}
