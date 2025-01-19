package utils

import (
	"encoding/json"
	"io"
)

func FormatJsonString(x any) string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

func UnmarshalReader(reader io.Reader, v interface{}) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(v)
}
