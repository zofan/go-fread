package fread

import "bytes"

var (
	XmlSplit     = `</`
	JsonSplit    = `,`
	CsvSplit     = []byte{',', ';', '\t'}
	SpaceSplit   = []byte(" \n\r\t")
	SpecialSplit = []byte("`~!@#$%^&*()_+-=[{]};:'\"\\|,<.>/?№() \n\r\t")
)

func ScanSplit(split string) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	splitBytes := []byte(split)

	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.Index(data, splitBytes); i >= 0 {
			return i + len(splitBytes), data[:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
}

func ScanSplitAny(splitBytes []byte) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if bytes.IndexByte(splitBytes, data[i]) >= 0 {
				return i + 1, data[:i], nil
			}
		}
		if atEOF && len(data) > 0 {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
}

func ScanSplitNotAny(splitBytes []byte) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if bytes.IndexByte(splitBytes, data[i]) == -1 {
				return i + 1, data[:i], nil
			}
		}
		if atEOF && len(data) > 0 {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
}
