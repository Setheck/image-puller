package main

import (
	"io/ioutil"
	"testing"
)

func TestEnclosureUrlsFromRssStringSingleItem(t *testing.T) {
	buffer, _ := ioutil.ReadFile("testdata/singleItemValidRss.rss")
	enclosure := EnclosureUrlsFromRssString(string(buffer))
	if len(enclosure) != 1 {
		t.Error("Enclosure count Mismatch")
	}
	for _, enc := range enclosure {
		verifyEnclosureData(enc, t)
	}
}

func TestEnclosureUrlsFromRssStringMultiItem(t *testing.T) {
	buffer, _ := ioutil.ReadFile("testdata/multiItemValidRss.rss")
	enclosure := EnclosureUrlsFromRssString(string(buffer))
	if len(enclosure) != 2 {
		t.Error("Enclosure count Mismatch")
	}
	for _, enc := range enclosure {
		verifyEnclosureData(enc, t)
	}
}

func TestEnclosureUrlsFromRssBytesSingleItem(t *testing.T) {
	buffer, _ := ioutil.ReadFile("testdata/singleItemValidRss.rss")
	enclosure := EnclosureUrlsFromRssBytes(buffer)
	if len(enclosure) != 1 {
		t.Error("Enclosure count Mismatch")
	}
	for _, enc := range enclosure {
		verifyEnclosureData(enc, t)
	}
}

func TestEnclosureUrlsFromRssBytesMultiItem(t *testing.T) {
	buffer, _ := ioutil.ReadFile("testdata/multiItemValidRss.rss")
	enclosure := EnclosureUrlsFromRssBytes(buffer)
	if len(enclosure) != 2 {
		t.Error("Enclosure count Mismatch")
	}
	for _, enc := range enclosure {
		verifyEnclosureData(enc, t)
	}
}

//Helper Verify Methods
func verifyEnclosureData(e Enclosure, t *testing.T) {
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
