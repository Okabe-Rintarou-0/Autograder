package utils

import "encoding/json"

func FormatJsonString(x any) string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}
