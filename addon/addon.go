package addon

// Addon represents an interface defining functionality for managing World of
// Warcraft addons.
type Addon interface {
	InitMetadata(string) (*Metadata, error)
	InitProfile(string) error
	Install(string) error
}
