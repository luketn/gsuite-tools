package main

import "os"

var ContentCache map[string]string

func GetStaticContent(name string) (string, error) {
	result, cached := ContentCache[name]
	var err error = nil
	if cached {
		f, error := os.Open("html/index.html")
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
				result = string(fileContent);
				ContentCache[name] = result
		}
		f.Close()
			}
	}
	if err != nil {
		return "", err
	}else {
		return result, nil
	}
}
