package fread

import (
	"encoding/json"
	"os"
)

func ReadJson(path string, i interface{}) error {
	fh, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}
	defer fh.Close()

	return json.NewDecoder(fh).Decode(i)
}
