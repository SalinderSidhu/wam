package addon

// Metadata represents an addon containing specific information such as its
// name, version, last updated date and source url.
type Metadata struct {
	Date    int64
	Name    string
	Source  string
	Version string
}

// Profile represents the JSON marshaling structure used to identify installed
// addons.
type Profile struct {
	Path   string `json:"path"`
	Addons []struct {
		Date int64  `json:"date"`
		ID   string `json:"id`
	} `json:"addons"`
}

// Addon represents an interface defining functionality for managing World of
// Warcraft addons.
type Addon interface {
	InitMetadata(string) (*Metadata, error)
	InitProfile(string) error
	Install(string) error
}
