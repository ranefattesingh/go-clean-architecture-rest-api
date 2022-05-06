package utils

import (
	"encoding/json"
	"io"
)

func FromJSON(r io.Reader, d interface{}) error {
	return json.NewDecoder(r).Decode(d)
}

func ToJSON(w io.Writer, d interface{}) error {
	return json.NewEncoder(w).Encode(d)
}
