package curse

import (
	"../addon"
)

/*
Util represents web and file utilities for downloading and managing World of
Warcraft addons from the Curse website
*/
type Util struct {
	addonURL         string
	addonDownloadURL string
}

// NewUtil creates an instance of Utils
func NewUtil() addon.Util {
	return &Util{
		addonURL:         "https://mods.curse.com/addons/wow/%d",
		addonDownloadURL: "https://mods.curse.com/addons/wow/%d/download",
	}
}

/*
Download the current version of the Curse addon. Return an error if any occur,
return nil otherwise.
*/
func (u *Util) Download(id string) error {
	return nil
}
