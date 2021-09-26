package main

import (
	"testing"
)

func TestGetStaticContent(t *testing.T) {
	content, err := GetStaticContent("../resources/ui", "/index.html")
	if err != nil {
		t.Errorf("Failed to read index.html:\n%s", err)
	} else {
		if len(content) == 0 {
			t.Errorf("Failed to read index.html. %d length", len(content))
		}
	}
}
