package rssparse_test

import (
	"io/ioutil"
	"testing"

	"github.com/setheck/image-puller/rssparse"
)

func TestEnclosureUrlsFromRssStringSingleItem(t *testing.T) {
	buffer, _ := ioutil.ReadFile("testdata/singleItemValidRss.rss")
	enclosure := rssparse.EnclosureUrlsFromRssString(string(buffer))
	if len(enclosure) != 1 {
		t.Error("Enclosure count Mismatch")
	}
	for _, enc := range enclosure {
		verifyEnclosureData(enc, t)
	}
}

func TestEnclosureUrlsFromRssStringMultiItem(t *testing.T) {
	buffer, _ := ioutil.ReadFile("testdata/multiItemValidRss.rss")
	enclosure := rssparse.EnclosureUrlsFromRssString(string(buffer))
	if len(enclosure) != 2 {
		t.Error("Enclosure count Mismatch")
	}
	for _, enc := range enclosure {
		verifyEnclosureData(enc, t)
	}
}

func TestEnclosureUrlsFromRssBytesSingleItem(t *testing.T) {
	buffer, _ := ioutil.ReadFile("testdata/singleItemValidRss.rss")
	enclosure := rssparse.EnclosureUrlsFromRssBytes(buffer)
	if len(enclosure) != 1 {
		t.Error("Enclosure count Mismatch")
	}
	for _, enc := range enclosure {
		verifyEnclosureData(enc, t)
	}
}

func TestEnclosureUrlsFromRssBytesMultiItem(t *testing.T) {
	buffer, _ := ioutil.ReadFile("testdata/multiItemValidRss.rss")
	enclosure := rssparse.EnclosureUrlsFromRssBytes(buffer)
	if len(enclosure) != 2 {
		t.Error("Enclosure count Mismatch")
	}
	for _, enc := range enclosure {
		verifyEnclosureData(enc, t)
	}
}

//Helper Verify Methods
func verifyEnclosureData(e rssparse.Enclosure, t *testing.T) {
	if e.URL != "RssItemEnclosureUrl" {
		t.Error("Url Mismatch")
	}
	if e.Type != "RssItemEnclosureType" {
		t.Error("Type Mismatch")
	}
	if e.Length != "RssItemEnclosureLength" {
		t.Error("Length Mismatch")
	}
}
