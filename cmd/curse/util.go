package curse

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

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
		addonURL:         "https://mods.curse.com/addons/wow/%s",
		addonDownloadURL: "https://mods.curse.com/addons/wow/%s/download",
	}
}

/*
Download the current version of the Curse addon. Return an error if any occur,
return nil otherwise.
*/
func (u *Util) Download(id string) error {
	// Resolve the addon page from Curse
	page, err := u.resolve(fmt.Sprintf(u.addonDownloadURL, id))
	if err != nil {
		return err
	}
	// Parse the addon page for the download URL
	fileURL, err := u.parse(page, `data-href="([^"]*)`, `data-href="`)
	if err != nil {
		return err
	}
	// Obtain the file contents using a GET request
	res, err := http.Get(fileURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// Create a file to save the downloaded content
	out, err := os.Create(fmt.Sprintf("DL_%s.zip", id))
	if err != nil {
		return err
	}
	defer out.Close()
	// Copy the data from the request into the output file
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}
	return nil
}

func (u *Util) parse(s string, regex string, split string) (string, error) {
	re, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	rx := re.FindStringSubmatch(s)
	return strings.Split(rx[0], split)[1], nil
}

func (u *Util) resolve(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	bytes, _ := ioutil.ReadAll(res.Body)
	return string(bytes), nil
}
