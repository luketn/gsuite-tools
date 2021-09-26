package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func GetContentType(fileName string) (string, error) {
	index := strings.LastIndex(fileName, ".")
	if index == -1 {
		if fileName == "/" {
			return "text/html", nil
		} else {
			return "", errors.New("Path has no extension! Invalid request.")
		}
	}
	switch fileName[index:] {
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
				var f *os.File = nil
				f, err = os.Open(path)
				if err == nil {
					var info os.FileInfo = nil
					info, err = f.Stat()
					if err == nil {
						size := info.Size()
						var fileContent = make([]byte, size)
						var readBytes = 0
						readBytes, err = f.Read(fileContent)
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
