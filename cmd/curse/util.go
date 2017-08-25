package curse

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"../addon"
	"github.com/PuerkitoBio/goquery"
)

/*
Util represents web and file utilities for downloading and managing World of
Warcraft addons from the Curse website.
*/
type Util struct {
	addonURL string
}

// NewUtil creates an instance of Utils
func NewUtil() addon.Util {
	return &Util{
		addonURL: "https://mods.curse.com/addons/wow/%s",
	}
}

/*
GetData returns an addon data object parsed from Curse using a curse addon id.
*/
func (u *Util) GetData(id string) (*addon.Data, error) {
	// Parse id an obtain addon data from curse
	data, err := u.parse(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
Install downloads, extracts and installs an addon from Curse using a curse
addon id. Return an error if one occured.
*/
func (u *Util) Install(id string) error {
	// Parse id and obtain addon data from curse
	data, err := u.parse(id)
	if err != nil {
		return err
	}
	// Download the addon using the URL link
	err = u.downloadURL(data.URL, fmt.Sprintf("%s.zip", id))
	if err != nil {
		return err
	}
	// Extract the addon zip file
	err = u.extractZip(fmt.Sprintf("%s.zip", id), "./extracted")
	if err != nil {
		return err
	}
	return nil
}

func (u *Util) parse(id string) (*addon.Data, error) {
	// Resolve the addon page from Curse
	doc, err := goquery.NewDocument(fmt.Sprintf(u.addonURL, id))
	if err != nil {
		return nil, err
	}
	// Check for 404 page if the Curse addon was not found
	h := doc.Find("#content section header h2").First().Text()
	if h == "Not found" {
		return nil, fmt.Errorf("%s not found on Curse", id)
	}
	// Parse specific information from the page
	n := doc.Find("#project-overview > header > h2").First().Text()
	v := strings.Split(doc.Find("li.newest-file").First().Text(), ": ")[1]
	d, _ := doc.Find("li.updated abbr").Attr("data-epoch")
	e, err := strconv.ParseInt(d, 10, 64)
	if err != nil {
		return nil, err
	}
	// Parse download link and obtain the file url of the addon
	dDoc, err := goquery.NewDocument(fmt.Sprintf(u.addonURL, id) + "/download")
	if err != nil {
		return nil, err
	}
	l, _ := dDoc.Find("#file-download a").Attr("data-href")
	return &addon.Data{Name: n, Epoch: e, Version: v, URL: l}, nil
}

func (u *Util) downloadURL(url, fname string) error {
	// Obtain file's contents using a GET request
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// Create a file to output the file's contents
	out, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer out.Close()
	// Copy the data from the request body into the output file
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}
	return nil
}

func (u *Util) extractZip(src, dest string) error {
	// Open zip file for reading
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()
	// Open and read each file and/or folder
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		// Create files and folders at dest from zip contents
		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, f.Mode())
		} else {
			var fdir string
			lastIndex := strings.LastIndex(fpath, string(os.PathSeparator))
			if lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}
			err = os.MkdirAll(fdir, f.Mode())
			if err != nil {
				return err
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
