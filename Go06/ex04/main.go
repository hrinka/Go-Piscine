package main

import (
	"fmt"
	"os"
	"piscine"
)

func atoi(s string) int {
	result := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return -1
		}
		result = result*10 + int(r-'0')
	}
	return result
}

func main() {
	if len(os.Args) < 3 || os.Args[1] != "-c" {
		fmt.Println("Usage: go run . -c <number> <file1> [<file2> ...]")
		return
	}

	n := atoi(os.Args[2])
	if n <= 0 {
		fmt.Println("The number of bytes must be a positive integer")
		return
	}

	files := os.Args[3:]
	error0ccured := false

	for _, file := range files {
		if len(files) > 1 {
			fmt.Printf("==> %s <==\n", file)
		}

		err := piscine.Ztail(file, n)
		if err != nil {
			fmt.Println(err)
			error0ccured = true
		}
		if len(files) > 1 {
			fmt.Println()
		}
	}
	if error0ccured {
		fmt.Println("One or more files counld not be processed.")
	}
}
