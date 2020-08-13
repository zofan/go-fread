package fread

import (
	"bufio"
	"strings"
)

func ReadTldZone(file string, output chan string) error {
	reader, err := GetReader(file)
	if err != nil {
		return err
	}

	defer reader.Close()

	var lastValue string

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.SplitN(line, "\t", 2)

		if len(cols) < 2 {
			continue
		}

		value := cols[0]
		if lastValue == value {
			continue
		}

		lastValue = value
		output <- value
	}

	return nil
}
