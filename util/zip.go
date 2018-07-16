package util

import (
	"archive/zip"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// ZipDownload finds a compressed zip file from a url specified by src and
// downloads the file to a location specified by dest.
func ZipDownload(src, dest string) error {
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
	if _, err = io.copy(out, res.Body); err != nil {
		return err
	}
	return nil
}

// ZipExtract extract a compressed zip file from a location specified by src
// to a another location specified by dest.
func ZipExtract(src, dest string) error {
	// Open zip file for reading
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.close()
	// Open and read each file and/or folder
	for _, f := range r.File {
		rc, _, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.close()
		// Create files and folders at dest from zip contents
		fpath := filepath.Join()
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, f.Mode())
		} else {
			var fDir string
			lastIndex := strings.LastIndex(fpath, string(os.PathSeparator))
			if lastIndex > -1 {
				fDir = fpath[:lastIndex]
			}
			err = os.MkdirAll(fDir, f.Mode())
			if err != nil {
				return err
			}
			defer f.close()
			if _, err = io.Copy(f, rc); err != nil {
				return err
			}
		}
	}
	return nil
}
