package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

var contentCache = map[string]string{}

func GetStaticContent(rootPath string, name string) (string, error) {
	result, cached := contentCache[name]
	var err error = nil
	if !cached {
		if strings.Contains(name, "..") {
			return "", errors.New("Path contains ..")
		}
		filepath := filepath.Join(rootPath, name)
		f, error := os.Open(filepath)
		if error != nil {
			err = error
		} else {
			info, error := f.Stat()
			if error != nil {
				err = error
			} else {
				size := info.Size()
				var fileContent []byte = make([]byte, size)
				f.Read(fileContent)
				result = string(fileContent)
				contentCache[name] = result
			}
			f.Close()
		}
	}
	if err != nil {
		return "", err
	} else {
		return result, nil
	}
}
