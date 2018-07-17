package addon

// Profile represents the JSON marshaling structure used to identify installed
// addons.
type Profile struct {
	Path   string `json:"path"`
	Addons []struct {
		Date int64  `json:"date"`
		ID   string `json:"id"`
	} `json:"addons"`
}
