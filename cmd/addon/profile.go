package addon

import (
	"encoding/json"
	"os"
)

/*
Profile represents the JSON marshaling strucutre of wam.json
This struct contains IDs, dates and files of the installed addons.
*/
type Profile struct {
	Path   string `json:"path"`
	Addons []struct {
		ID           string   `json:"id"`
		DateEpoch    int64    `json:"dateEpoch"`
		DataLocation []string `json:"dataLocation"`
	} `json:"addons"`
}

/*
Read JSON data from file and load into Profile object. Return an error if one
occured.
*/
func (p *Profile) Read(f string) error {
	// Open file containing profile data
	in, err := os.Open(f)
	if err != nil {
		return err
	}
	defer in.Close()
	// Parse JSON data from file into Profile
	if err = json.NewDecoder(in).Decode(p); err != nil {
		return err
	}
	return nil
}

/*
Write data from Profile into a file as JSON. Return an error if one occured.
*/
func (p *Profile) Write(f string) error {
	// Create a new file to store contents of profile
	out, err := os.Create(f)
	if err != nil {
		return err
	}
	defer out.Close()
	// Output Profile data to file as JSON
	if err := json.NewEncoder(out).Encode(p); err != nil {
		return err
	}
	return nil
}
