package curse

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"../addon"
	"github.com/PuerkitoBio/goquery"
)

// Utils implements the corresponding interface from addon.Utils
type Utils struct {
	addonURL     string
	defaultPaths map[string]string
}

// NewUtils creates an instance of Utils
func NewUtils() addon.Utils {
	return &Utils{
		addonURL: "https://mods.curse.com/addons/wow/%s",
		defaultPaths: map[string]string{
			"windows": "C:/Program Files (x86)/World of Warcraft/Interface/AddOns",
			"darwin":  "/Applications/Battle.net/World of Warcraft/Interface/AddOns",
		},
	}
}

/*
Init creates a new addon profile file (wam.json) with the World of Warcraft
installation path p. The default path for the current OS is selected if p is not
provided. Return an error if one occured.
*/
func (u *Utils) Init(p string) error {
	// Obtain the executable directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	// Name and full path of the wam.json file
	fpath := fmt.Sprintf("%s/wam.json", dir)
	// Check if file does not exist
	if _, err := os.Stat(fpath); !os.IsNotExist(err) {
		// Return specific error to prevent overwriting existing file
		return fmt.Errorf("existing addon profile found in wam.json")
	}
	// Create a new file to store the contents of the wam file
	out, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer out.Close()
	// Create a wam file and assign wow installation path
	wfile := addon.WamFile{}
	if wfile.Path = p; wfile.Path == "" {
		wfile.Path = u.defaultPaths[runtime.GOOS]
	}
	// Output Wam file data to the wam.json file
	e := json.NewEncoder(out)
	if err := e.Encode(wfile); err != nil {
		return err
	}
	return nil
}

/*
GetData returns an addon data object parsed from Curse using a Curse addon id.
*/
func (u *Utils) GetData(id string) (*addon.Data, error) {
	// Parse id an obtain addon data from Curse
	data, err := u.parseCurse(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
Install downloads, extracts and installs an addon from Curse using a Curse
addon id. Return an error if one occured.
*/
func (u *Utils) Install(id string) error {
	// Obtain the addon profile to store data of newly installed addons
	wFile, err := u.parseWamFile()
	if err != nil {
		return err
	}
	// Obtain the executable directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	// Name and full path of the addon zip file based on the id
	fpath := fmt.Sprintf("%s/%s.zip", dir, id)
	// Parse id and obtain addon data from Curse
	data, err := u.parseCurse(id)
	if err != nil {
		return err
	}
	// Download the addon zip file using the URL link
	if err = u.downloadZip(data.URL, fpath); err != nil {
		return err
	}
	// Extract the zip file to a tmp folder
	if err = u.extractZip(fpath, wFile.Path); err != nil {
		return err
	}
	// Delete the downloaded zip file from the tmp folder
	if err = os.Remove(fpath); err != nil {
		return err
	}
	return nil
}

func (u *Utils) parseCurse(id string) (*addon.Data, error) {
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
	return &addon.Data{Name: n, DateEpoch: e, Version: v, URL: l}, nil
}

func (u *Utils) parseWamFile() (*addon.WamFile, error) {
	// Obtain the executable directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, err
	}
	// Name and full path of the wam.json file
	fpath := fmt.Sprintf("%s/wam.json", dir)
	// Check if wam.json exists
	if _, err := os.Stat(fpath); os.IsNotExist(err) {
		return nil, fmt.Errorf(
			"addon profile required, please create using \"wam init\"")
	}
	// Open the wam.json
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// Create wam file object with json data from wam.json
	wfile := addon.WamFile{}
	if err = json.NewDecoder(f).Decode(&wfile); err != nil {
		return nil, err
	}
	return &wfile, nil
}

func (u *Utils) downloadZip(src, dest string) error {
	// Obtain file's contents from src using a GET request
	res, err := http.Get(src)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// Create a zip file to output the file's contents
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	// Copy the data from the result body into the output file
	if _, err = io.Copy(out, res.Body); err != nil {
		return err
	}
	return nil
}

func (u *Utils) extractZip(src, dest string) error {
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
			if _, err = io.Copy(f, rc); err != nil {
				return err
			}
		}
	}
	return nil
}
