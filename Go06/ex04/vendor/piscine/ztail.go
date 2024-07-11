package piscine

import (
	"fmt"
	"os"
)

func Ztail(filename string, count int) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("open %s: %v", filename, err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("stat %s: %v", filename, err)
	}

	size := stat.Size()
	if count > int(size) {
		count = int(size)
	}

	buffer := make([]byte, count)
	_, err = file.Seek(-int64(count), os.SEEK_END)
	if err != nil {
		return fmt.Errorf("seek %s: %v", filename, err)
	}
	_, err = file.Read(buffer)
	if err != nil {
		return fmt.Errorf("read %s: %v", filename, err)
	}

	fmt.Print(string(buffer))
	return nil
}
