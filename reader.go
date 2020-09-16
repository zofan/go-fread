package fread

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func GetReader(fileName string) (io.ReadCloser, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	var reader io.ReadCloser

	ext := filepath.Ext(filepath.Base(fileName))

	head, err := ReadHead(file, 10)
	if err != nil {
		return nil, err
	}

	if strings.Contains(ext, `.zip`) || bytes.HasPrefix(head, []byte{'\x50', '\x4B'}) {
		stat, err := file.Stat()
		if err != nil {
			return nil, err
		}

		zipFile, err := zip.NewReader(file, stat.Size())
		if err != nil {
			return nil, err
		}

		if len(zipFile.File) == 0 {
			return nil, errors.New(`zip does not contains files`)
		}

		sort.Slice(zipFile.File, func(i, j int) bool {
			return zipFile.File[i].UncompressedSize64 > zipFile.File[j].UncompressedSize64
		})

		reader, err = zipFile.File[0].Open()
		if err != nil {
			return nil, err
		}
	} else if strings.Contains(ext, `.gz`) || bytes.HasPrefix(head, []byte{'\x1F', '\x8B'}) {
		reader, err = gzip.NewReader(file)
		if err != nil {
			return nil, err
		}
	} else {
		reader = file
	}

	return reader, err
}

func ReadHead(file *os.File, size int) ([]byte, error) {
	var head = make([]byte, size)

	_, err := file.ReadAt(head, 0)
	if err != nil {
		return nil, err
	}
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	return head, nil
}
