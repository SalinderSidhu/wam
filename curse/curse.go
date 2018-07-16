package curse

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/salindersidhu/wam/addon"
)

// Addon represents the corresponding interface from addon.Addon
type Addon struct {
	source string
}

// NewAddon creates an instance of Addon.
func NewAddon() addon.Addon {
	return &Addon{
		source: "https://www.curseforge.com/wow/addons/%s",
	}
}

// InitMetadata returns the metadata object of an addon parsed from Curse where
// the curse addon id is specified by id. Return an error if one occurred.
func (a *Addon) InitMetadata(id string) (*addon.Metadata, error) {
	data, err := a.parseCurse(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// InitProfile creates a new addon profile (wam.json) with the World of
// Warcraft installation path specified by fpath. Return an error if one
// occurred.
func (a *Addon) InitProfile(fpath string) error {
	return nil
}

// Install downloads, extracts and installs an addon from Curse using a Curse
// addon id. Return an error if one occurred.
func (a *Addon) Install(id string) error {
	return nil
}

func (a *Addon) parseCurse(id string) (*addon.Metadata, error) {
	// Resolve the addon page from Curse
	doc, err := goquery.NewDocument(fmt.Sprintf(a.source, id))
	if err != nil {
		return nil, err
	}
	// Check for 404 page if the Curse addon was not found
	h := doc.Find("h2").First().Text()
	if h == "Not found" {
		return nil, fmt.Errorf("%s not found on Curse", id)
	}
	// Parse specific information from the page
	n := doc.Find("#content section header h2").First().Text()
	v := strings.Split(doc.Find(".stats--game-version").First().Text(), ": ")[1]
	d, _ := doc.Find(".stats--last-updated abbr").Attr("data-epoch")
	fmt.Println(d)

	e, err := strconv.ParseInt(d, 10, 64)
	if err != nil {
		return nil, err
	}
	// Parse download link and obtain the file url of the addon
	dDoc, err := goquery.NewDocument(fmt.Sprintf(a.source, id) + "/download")
	if err != nil {
		return nil, err
	}
	dlPart, _ := dDoc.Find("a.download__link").Attr("href")
	l := fmt.Sprintf(a.source, strings.Split(dlPart, "/wow/addons/")[1])
	return &addon.Metadata{Name: n, Date: e, Version: v, Source: l}, nil
}
