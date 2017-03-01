package main

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	var testFileName = "testdata/urls.txt"
	var etalon = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	var urls, err = readFile(&testFileName)
	if err != nil {
		t.Errorf("Error while read test data from \"%s\": %s", testFileName, err.Error())
		t.FailNow()
	}
	if len(urls) != len(etalon) {
		t.Errorf("Line count from \"%s\" is %d, want %d", testFileName, len(urls), len(etalon))
	}
}
