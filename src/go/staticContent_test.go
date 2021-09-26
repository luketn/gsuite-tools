package main

import (
	"testing"
)

func TestGetStaticContent(t *testing.T) {
	content, contentType, err := GetStaticContent("../resources/ui", "/index.html")
	if err != nil {
		t.Errorf("Failed to read index.html:\n%s", err)
	} else {
		if len(content) == 0 {
			t.Errorf("Failed to read index.html. %d length", len(content))
		}else {
			if contentType != "text/html" {
				t.Errorf("content type check failed: should be: text/html but was: %s", contentType)
			}
		}
	}
}
