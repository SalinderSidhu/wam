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
