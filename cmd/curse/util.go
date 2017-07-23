package curse

import (
	"../addon"
)

/*
Util represents web and file utilities for downloading and managing World of
Warcraft addons from the Curse website
*/
type Util struct {
	AddonURL string
}

// NewUtil creates an instance of Utils
func NewUtil() addon.Util {
	return &Util{
		AddonURL: "http://mods.curse.com/addons/wow/%s/download",
	}
}

/*
GetCurrentVersion returns the current version string of the Curse addon
by a Curse identifier
*/
func (u *Util) GetCurrentVersion(id string) (string, error) {
	return "", nil
}
