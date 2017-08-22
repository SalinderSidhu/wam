package curse

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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

	return &addon.Data{ID: id, Name: n, Epoch: e, Version: v, URL: l}, nil
}

/*
GetInfo returns a string containing the following information about an addon
(specified by id) from curse: name, date and version
*/
func (u *Util) GetInfo(id string) (string, error) {
	// Parse id an obtain addon data from curse
	data, err := u.parse(id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s (%s) Updated: %s", data.Name, data.Version,
		time.Unix(data.Epoch, 0).Format(time.RFC822Z)), nil
}

/*
Download function finds and downloads the latest version of an addon (specified
by id) from Curse. Return an error if one occured, otherwise return nil
*/
func (u *Util) Download(id string) error {
	// Parse id and obtain addon data from curse
	data, err := u.parse(id)
	if err != nil {
		return err
	}
	// Obtain the addon file content using a GET request
	res, err := http.Get(data.URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// Create a file to save the downloaded addon
	out, err := os.Create(fmt.Sprintf("%s.zip", id))
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
