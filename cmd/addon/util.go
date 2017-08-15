package addon

/*
Util represents an interface containing web and file utilities for downloading
and managing World of Warcraft addons
*/
type Util interface {
	Download(*Data) (string, error)
	GetData(string) (*Data, error)
}

// Data represents a structure containing specific information about an addon
type Data struct {
	ID      string
	Name    string
	Date    int64
	Version string
	URL     string
}

// Utils represents an array of Util interfaces
type Utils []Util
