package fread

import (
	"archive/zip"
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var (
	ZipSelector = ZipMaxSizeSelector

	ErrZipNotFileContains = errors.New(`fread: zip does not contains files`)
)

func NewReader(filePath string) (io.ReadCloser, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	head, err := Head(file, 10)
	if err != nil {
		return nil, err
	}

	var reader io.ReadCloser

	ext := filepath.Ext(filepath.Base(filePath))
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
			return nil, ErrZipNotFileContains
		}

		reader, err = ZipSelector(zipFile)
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

func ChunkSplit(filePath string, split []byte, output chan []byte) error {
	r, err := NewReader(filePath)
	if err != nil {
		return err
	}
	defer r.Close()

	s := bufio.NewScanner(r)
	s.Split(ScanSplit(split))
	for s.Scan() {
		output <- s.Bytes()
	}

	return s.Err()
}

func ChunkSplitAny(filePath string, split []byte, output chan []byte) error {
	r, err := NewReader(filePath)
	if err != nil {
		return err
	}
	defer r.Close()

	s := bufio.NewScanner(r)
	s.Split(ScanSplitAny(split))
	for s.Scan() {
		output <- s.Bytes()
	}

	return s.Err()
}

func Lines(filePath string, output chan []byte) error {
	r, err := NewReader(filePath)
	if err != nil {
		return err
	}
	defer r.Close()

	s := bufio.NewScanner(r)
	for s.Scan() {
		output <- s.Bytes()
	}

	return s.Err()
}

func Head(file *os.File, size int) ([]byte, error) {
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

func ZipMaxSizeSelector(zipFile *zip.Reader) (io.ReadCloser, error) {
	sort.Slice(zipFile.File, func(i, j int) bool {
		return zipFile.File[i].UncompressedSize64 > zipFile.File[j].UncompressedSize64
	})

	return zipFile.File[0].Open()
}
