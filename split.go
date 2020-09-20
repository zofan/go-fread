package fread

import (
	"bufio"
	"bytes"
)

var (
	XmlSplit        = []byte(`</`)
	SpaceAnySplit   = []byte(" \n\r\t")
	SpecialAnySplit = []byte("`~!@#$%^&+=[]{};'\"\\|,<>/?№() \n\r\t")
)

func ScanSplit(split []byte) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if i := bytes.Index(data, split); i >= 0 {
			return i + len(split), data[:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
}

func ScanSplitAny(split []byte) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if bytes.IndexByte(split, data[i]) >= 0 {
				return i + 1, data[:i], nil
			}
		}
		if atEOF && len(data) > 0 {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
}

func ScanSplitNotAny(split []byte) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if bytes.IndexByte(split, data[i]) == -1 {
				return i + 1, data[:i], nil
			}
		}
		if atEOF && len(data) > 0 {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
}
