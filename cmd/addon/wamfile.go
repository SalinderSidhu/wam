package addon

/*
WamFile represents the JSON marshaling structure of wam.json. This file contains
ids and dates of installed addons.
*/
type WamFile struct {
	Path   string `json:"path"`
	Addons []struct {
		ID        string `json:"id"`
		DateEpoch int64  `json:"dateEpoch"`
	} `json:"addons"`
}
