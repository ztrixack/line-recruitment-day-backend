package utils

import (
	"bytes"
	"encoding/json"
)

func Transcode(in, out interface{}) error {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(in); err != nil {
		return err
	}
	if err := json.NewDecoder(buf).Decode(out); err != nil {
		return err
	}

	return nil
}

func JsonFilter(in, out interface{}) error {
	buf, err := json.Marshal(in)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(buf, &out); err != nil {
		return err
	}

	return nil
}
