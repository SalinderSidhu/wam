package addon

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/salindersidhu/wam/util"
)

// Curse represents the corresponding interface from addon.Addon
type Curse struct {
	Source       string
	DefaultPaths map[string]string
}

// NewCurse creates an instance of Curse.
func NewCurse() Addon {
	return &Curse{
		Source: "https://www.curseforge.com/wow/addons/%s",
		DefaultPaths: map[string]string{
			"windows": "C:/Program Files/World of Warcraft/_retail_/Interface/AddOns",
			"darwin":  "/Applications/Battle.net/World of Warcraft/_retail_/Interface/AddOns",
		},
	}
}

// InitMetadata returns the metadata object of an addon parsed from Curse where
// the curse addon id is specified by id. Return an error if one occurred.
func (c *Curse) InitMetadata(id string) (*Metadata, error) {
	data, err := c.parseCurse(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// InitProfile creates a new profile (wam.json) with the World of Warcraft
// installation path specified by f. Return an error if one occurred.
func (c *Curse) InitProfile(f string) error {
	// Obtain the executable directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	// Name and full path of the wam.json file
	fPath := fmt.Sprintf("%s/wam.json", dir)
	// Check if file does not exist
	if _, err := os.Stat(fPath); !os.IsNotExist(err) {
		// Return specific error to prevent overwriting existing file
		return fmt.Errorf("existing profile found in wam.json")
	}
	// Create a wam file and assign wow installation path
	pFile := &Profile{}
	if pFile.Path = f; pFile.Path == "" {
		pFile.Path = c.DefaultPaths[runtime.GOOS]
	}
	// Output Wam file data to the wam.json file
	return util.JSONOut(fPath, pFile)
}

// Install downloads, extracts and installs an addon from Curse using a Curse
// addon id. Return an error if one occurred.
func (c *Curse) Install(id string) error {
	// Obtain the executable directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	// Obtain the profile to store data of newly installed addons
	pFile := &Profile{}
	if err := util.JSONIn(fmt.Sprintf("%s/wam.json", dir), pFile); err != nil {
		return err
	}
	// Name and full path of the addon zip file based on the id
	fPath := fmt.Sprintf("%s/%s.zip", dir, id)
	// Parse id and obtain addon data from Curse
	data, err := c.parseCurse(id)
	if err != nil {
		return err
	}
	// Download the addon zip file using the URL link
	if err = util.ZipDownload(data.Source, fPath); err != nil {
		return err
	}
	// Extract the zip file to a tmp folder
	if err = util.ZipExtract(fPath, pFile.Path); err != nil {
		return err
	}
	// Delete the downloaded zip file from the tmp folder
	return os.Remove(fPath)
}

func (c *Curse) parseCurse(id string) (*Metadata, error) {
	// Resolve the addon page from Curse
	doc, err := goquery.NewDocument(fmt.Sprintf(c.Source, id))
	if err != nil {
		return nil, err
	}
	// Check for 404 page if the Curse addon was not found
	h := doc.Find("h2").First().Text()
	if h == "Not found" {
		return nil, fmt.Errorf("%s not found on Curse", id)
	}
	// Parse specific information from the page
	n := doc.Find("h2.font-bold.text-lg.break-all").First().Text()
	v := strings.Split(doc.Find("div.flex > span:nth-child(3)").First().Text(), ": ")[1]
	d, _ := doc.Find("span.mr-2.text-gray-500 abbr").Attr("data-epoch")

	e, err := strconv.ParseInt(d, 10, 64)
	if err != nil {
		return nil, err
	}
	// Parse download link and obtain the file url of the addon
	dDoc, err := goquery.NewDocument(fmt.Sprintf(c.Source, id) + "/download")
	if err != nil {
		return nil, err
	}
	dlPart, _ := dDoc.Find("p.text-sm a").Attr("href")
	l := fmt.Sprintf(c.Source, strings.Split(dlPart, "/wow/addons/")[1])
	return &Metadata{Name: n, Date: e, Version: v, Source: l}, nil
}
