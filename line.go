package fread

import (
	"bufio"
)

func Lines(filePath string, output chan []byte) error {
	reader, err := NewReader(filePath)
	if err != nil {
		return err
	}
	defer reader.Close()

	s := bufio.NewScanner(reader)
	for s.Scan() {
		output <- s.Bytes()
	}

	return s.Err()
}
