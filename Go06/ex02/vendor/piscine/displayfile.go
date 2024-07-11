package piscine

import (
	"fmt"
	"ft"
	"os"
)

func DisplayFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println(err)
			}
			break
		}
		for i := 0; i < n; i++ {
			ft.PrintRune(rune(buffer[i]))
		}
	}
}
