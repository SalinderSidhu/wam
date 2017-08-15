package curse

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"../addon"
	"github.com/PuerkitoBio/goquery"
)

/*
Util represents web and file utilities for downloading and managing World of
Warcraft addons from the Curse website
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
GetData returns an addon data object containing information about a Curse addon
using the specified id and a nil error. Return nil for the addon data and a
specific error otherwise.
*/
func (u *Util) GetData(id string) (*addon.Data, error) {
	// Resolve the addon page from Curse
	doc, err := goquery.NewDocument(fmt.Sprintf(u.addonURL, id))
	if err != nil {
		return nil, err
	}
	// Check for 404 page if the Curse addon was not found
	h := doc.Find("#content section header h2").First().Text()
	if h == "Not found" {
		return nil, errors.New("Addon not found")
	}
	// Parse specific information from the page
	n := doc.Find("#project-overview > header > h2").First().Text()
	v := strings.Split(doc.Find("li.newest-file").First().Text(), ": ")[1]
	d, _ := doc.Find("li.updated abbr").Attr("data-epoch")
	t, err := strconv.ParseInt(d, 10, 64)
	if err != nil {
		return nil, err
	}
	// Parse download link and obtain the file url of the addon
	dDoc, err := goquery.NewDocument(fmt.Sprintf(u.addonURL, id) + "/download")
	if err != nil {
		return nil, err
	}
	l, _ := dDoc.Find("#file-download a").Attr("data-href")

	return &addon.Data{ID: id, Name: n, Date: t, Version: v, URL: l}, nil
}

/*
Download downloads the latest version of a Curse addon from the addon data.
Return the file name of the downloaded addon and a nil error. Return nil for
the file name and a specific error otherwise.
*/
func (u *Util) Download(d *addon.Data) (string, error) {
	// Obtain the addon file content using a GET request
	res, err := http.Get(d.URL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	// Create a file to save the downloaded content
	out, err := os.Create(fmt.Sprintf("%s.zip", d.ID))
	if err != nil {
		return "", err
	}
	defer out.Close()
	// Copy the data from the request into the output file
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.zip", d.ID), nil
}
