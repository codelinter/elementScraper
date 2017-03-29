package elementScraper

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update tesdata expected files")

// fakeCC implements replaceImpl type
func fakeCC(raw string) string {
	fkOutput := newfakeOutput()
	return fkOutput.output()
}

func TestElementScraper(t *testing.T) {
	input := TestInput
	scrptr := NewElementScraper(input, `<script>`).ElementFunc(fakeCC)

	got := scrptr.Run()

	filename := filepath.Join("testdata", "scraper.txt")
	if *update {
		ioutil.WriteFile(filename, []byte(got), 0644)
	}
	expected, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
	}
	if got != string(expected) {
		t.Errorf("Got \n%v \n  but Expected \n%v", got, string(expected))
	}
}
