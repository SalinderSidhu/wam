package util

import (
	"encoding/json"
	"os"
)

// JSONIn reads JSON data from a filepath specified by f and loads it into a
// JSON marshalled struct specified by v.
func JSONIn(f string, v interface{}) error {
	in, err := os.Open(f)
	if err != nil {
		return err
	}
	defer in.Close()
	return json.NewDecoder(in).Decode(v)
}

// JSONOut writes a JSON marshalled struct specified by v to a filepath
// specified by f.
func JSONOut(f string, v interface{}) error {
	out, err := os.Create(f)
	if err != nil {
		return err
	}
	defer out.Close()
	return json.NewEncoder(out).Encode(v)
}
