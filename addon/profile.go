package addon

import (
	"encoding/json"
	"os"
)

// A Profile represents the JSON marshaling structure of wam.json.
// This struct contains IDs, dates and files of the installed addons.
type Profile struct {
	Path   string `json:"path"`
	Addons []struct {
		DateEpoch    int64    `json:"dateEpoch"`
		DataLocation []string `json:"dataLocation"`
		ID           string   `json:"id"`
	} `json:"addons"`
}

// Read JSON data from file and load into Profile object.
// Return an error if one occured.
func (p *Profile) Read(f string) error {
	// Open file containing profile data
	in, err := os.Open(f)
	if err != nil {
		return err
	}
	defer in.Close()
	// Parse JSON data from file into Profile
	return json.NewDecoder(in).Decode(p)
}

// Write data from Profile into a file as JSON.
// Return an error if one occurred.
func (p *Profile) Write(f string) error {
	// Create a new file to store contents of profile
	out, err := os.Create(f)
	if err != nil {
		return err
	}
	defer out.Close()
	// Output Profile data to file as JSON
	return json.NewEncoder(out).Encode(p)
}
