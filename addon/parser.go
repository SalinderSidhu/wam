package addon

// A Parser represents an interface containing functionality for downloading,
// installing, updating and removing World of Warcraft addons.
type Parser interface {
	GetData(string) (*Data, error)
	Init(string) error
	Install(string) error
}
