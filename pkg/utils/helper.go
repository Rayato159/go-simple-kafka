package utils

import (
	"encoding/json"
)

func CompressToJsonBytes(obj any) []byte {
	raw, _ := json.Marshal(obj)
	return raw
}
