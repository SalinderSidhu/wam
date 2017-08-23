package addon

/*
Util represents an interface containing web and file utilities for downloading
and managing World of Warcraft addons.
*/
type Util interface {
	ExtractZip(string, string) error
	GetInfo(string) (string, error)
	Download(string) error
}

// Utils represents an array of Util interfaces
type Utils []Util
