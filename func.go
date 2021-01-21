package fread

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadJson(path string, i interface{}) error {
	fh, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer fh.Close()

	raw, err := ioutil.ReadAll(fh)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, i)
}
