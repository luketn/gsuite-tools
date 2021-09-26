package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func GetContentType(fileName string) (string, error) {
	switch fileName[strings.LastIndex(fileName, "."):] {
	case ".txt":
		return "text/plain", nil
	case ".html":
		return "text/html", nil
	case ".css":
		return "text/css", nil
	case ".js":
		return "application/javascript", nil
	case ".png":
		return "image/png", nil
	case ".ico":
		return "image/vnd.microsoft.icon", nil
	default:
		return "", errors.New("Unknown type for " + fileName)
	}
}

var contentCache = map[string]string{}

// Returns the static content in the file with name 'name' in directory 'rootPath'. Returns the content, the content type and the error (or nil)
func GetStaticContent(rootPath string, name string) (string, string, error) {
	cached := false
	result := ""
	contentType, err := GetContentType(name)
	if err == nil {
		result, cached = contentCache[name]
		if !cached {
			if strings.Contains(name, "..") {
				err = errors.New("Path contains ..")
			} else {
				path := filepath.Join(rootPath, name)
				f, err := os.Open(path)
				if err == nil {
					info, err := f.Stat()
					if err == nil {
						size := info.Size()
						var fileContent = make([]byte, size)
						readBytes, err := f.Read(fileContent)
						if err == nil && int64(readBytes) == size {
							result = string(fileContent)
							contentCache[name] = result
						}
					}
					err = f.Close()
				}
			}
		}
	}
	return result, contentType, err
}
