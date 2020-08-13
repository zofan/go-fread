package fread

import "bytes"

var (
	XmlSplit  = `</`
	JsonSplit = `,`
	CsvSplit  = []byte{',', ';', '\t'}
	AnySplit  = []byte("[]{}<>(),!?;:\"'`/| \n\r\t")
)

func scanSplit(split string) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	splitBytes := []byte(split)

	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.Index(data, splitBytes); i >= 0 {
			return i + len(splitBytes), data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
}

func scanSplitAny(splitBytes []byte) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			for _, sb := range splitBytes {
				if data[i] == sb {
					return i + 1, data[0:i], nil
				}
			}
		}
		if atEOF && len(data) > 0 {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
}
