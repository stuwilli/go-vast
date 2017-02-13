package vast_test

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"

	vast2 "github.com/stuwilli/go-vast2"
)

func loadFixture(path string) (*vast2.VAST, error) {

	xmlFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()
	b, _ := ioutil.ReadAll(xmlFile)

	var v vast2.VAST
	err = xml.Unmarshal(b, &v)
	return &v, err
}

func TestVastVersion(t *testing.T) {

	v, err := loadFixture("testdata/document_with_one_wrapper_ad.xml")

	if err != nil {
		t.Fatal(err)
	}

	if v.Version != 2.0 {
		t.Error("Expected Vast.Version to be 2.0, got", v.Version)
	}
}
