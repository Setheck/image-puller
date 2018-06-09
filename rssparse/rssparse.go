package rssparse

import (
	"encoding/xml"
)

//Rss root rss node
type Rss struct {
	RssChannel RssChannel `xml:"channel"`
}

//RssChannel channel node
type RssChannel struct {
	RssItem []RssItem `xml:"item"`
}

//RssItem item node
type RssItem struct {
	Title         string    `xml:"title"`
	Link          string    `xml:"link"`
	LinkEnclosure Enclosure `xml:"enclosure"`
	Description   string    `xml:"description"`
}

//Enclosure rss link enclosure
type Enclosure struct {
	URL    string `xml:"url,attr"`
	Type   string `xml:"type,attr"`
	Length string `xml:"length,attr"`
}

//EnclosureUrlsFromRssString take
func EnclosureUrlsFromRssString(body string) []Enclosure {
	return EnclosureUrlsFromRssBytes([]byte(body))
}

//EnclosureUrlsFromRssBytes scrapes urls from enclosures.
func EnclosureUrlsFromRssBytes(body []byte) []Enclosure {
	var rss Rss
	xml.Unmarshal(body, &rss)

	var result []Enclosure
	for _, v := range rss.RssChannel.RssItem {
		result = append(result, v.LinkEnclosure)
	}
	return result
}
