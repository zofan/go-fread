package fread

import (
	"bufio"
)

func ReadLines(file string, output chan string) error {
	reader, err := GetReader(file)
	if err != nil {
		return err
	}
	defer reader.Close()

	s := bufio.NewScanner(reader)
	for s.Scan() {
		output <- s.Text()
	}

	return s.Err()
}
