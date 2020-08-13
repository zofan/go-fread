package fread

import (
	"bufio"
)

func ReadChunkSplit(file string, split string, output chan string) error {
	reader, err := GetReader(file)
	if err != nil {
		return err
	}

	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	scanner.Split(scanSplit(split))
	for scanner.Scan() {
		output <- scanner.Text()
	}

	return nil
}

func ReadChunkSplitAny(file string, splitBytes []byte, output chan string) error {
	reader, err := GetReader(file)
	if err != nil {
		return err
	}

	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	scanner.Split(scanSplitAny(splitBytes))
	for scanner.Scan() {
		output <- scanner.Text()
	}

	return nil
}
