package fread

import (
	"bufio"
)

func ReadChunkSplit(file string, split string, output chan []byte) error {
	reader, err := GetReader(file)
	if err != nil {
		return err
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	scanner.Split(ScanSplit(split))
	for scanner.Scan() {
		output <- scanner.Bytes()
	}

	return scanner.Err()
}

func ReadChunkSplitAny(file string, splitBytes []byte, output chan []byte) error {
	reader, err := GetReader(file)
	if err != nil {
		return err
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	scanner.Split(ScanSplitAny(splitBytes))
	for scanner.Scan() {
		output <- scanner.Bytes()
	}

	return scanner.Err()
}
