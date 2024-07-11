package piscine

import (
	"fmt"
	"os"
)

func countArgs(args []string) int {
	count := 0
	for range args {
		count++
	}
	return count
}

func Cat(files []string) {
	if countArgs(files) == 0 {
		readFromStdin()
		return
	}
	for _, file := range files {
		err := printFile(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
		}
	}
}

func readFromStdin() {
	buffer := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(buffer)
		if err != nil {
			break
		}
		os.Stdout.Write(buffer[:n])
	}
}

func printFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			break
		}
		os.Stdout.Write(buffer[:n])
	}
	return nil
}
