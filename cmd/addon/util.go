package addon

/*
Util represents an interface containing web and file utilities for downloading
and managing World of Warcraft addons.
*/
type Util interface {
	Install(string) error
	GetData(string) (*Data, error)
}

// Utils represents an array of Util interfaces
type Utils []Util
