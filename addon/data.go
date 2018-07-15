package addon

// Data represents an Addon with specific information such as a name, version,
// last updated date and download url.
type Data struct {
	DateEpoch int64
	Name      string
	URL       string
	Version   string
}
