package addon

/*
Data represents an Addon object with specific information such as name,
author, version, last updated date and download URL
*/
type Data struct {
	URL       string
	Name      string
	DateEpoch int64
	Version   string
}

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
