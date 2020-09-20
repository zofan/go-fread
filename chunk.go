package fread

import (
	"bufio"
)

func ChunkSplit(filePath string, split []byte, output chan []byte) error {
	reader, err := NewReader(filePath)
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

func ChunkSplitAny(filePath string, split []byte, output chan []byte) error {
	reader, err := NewReader(filePath)
	if err != nil {
		return err
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	scanner.Split(ScanSplitAny(split))
	for scanner.Scan() {
		output <- scanner.Bytes()
	}

	return scanner.Err()
}
