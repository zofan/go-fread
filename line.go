package fread

import (
	"bufio"
)

func ReadLine(file string, output chan string) error {
	reader, err := GetReader(file)
	if err != nil {
		return err
	}

	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		output <- scanner.Text()
	}

	return nil
}
