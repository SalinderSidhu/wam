package addon

// Metadata represents an addon containing specific information such as its
// name, version, last updated date and source url.
type Metadata struct {
	Date    int64
	Name    string
	Source  string
	Version string
}
