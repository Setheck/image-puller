package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	defaultFeed = "https://www.nasa.gov/rss/dyn/lg_image_of_the_day.rss"
)

type rss struct {
	RssChannel rssChannel `xml:"channel"`
}

type rssChannel struct {
	RssItem []rssItem `xml:"item"`
}

type rssItem struct {
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Enclosure   linkEnclosure `xml:"enclosure"`
	Description string        `xml:"description"`
}

type linkEnclosure struct {
	URL string `xml:"url,attr"`
}

func pullFeed(feed string) []string {
	resp, err := http.Get(defaultFeed)
	if err != nil {
		fmt.Printf("Error retrieving feed: %s", defaultFeed)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body")
		os.Exit(1)
	}
	// fmt.Print(string(body))

	var rss rss
	xml.Unmarshal(body, &rss)

	var result []string

	for _, v := range rss.RssChannel.RssItem {
		result = append(result, v.Enclosure.URL)
	}
	return result
}

func retrieveImage(imgURL string) (imageName string, imageData []byte) {
	resp, err := http.Get(imgURL)
	if err != nil {
		fmt.Printf("Error retrieving: %s", imgURL)
	}
	defer resp.Body.Close()
	splitString := strings.Split(imgURL, "/")
	imageName = splitString[len(splitString)-1]
	imageData, _ = ioutil.ReadAll(resp.Body)
	return imageName, imageData
}

func saveImage(fileName string, imageData []byte) {
	file, createError := os.Create(fileName)
	if createError != nil {
		log.Fatalf("Failed creating file: %s", fileName)
	}
	_, writeError := io.Copy(file, ioutil.NopCloser(bytes.NewBuffer(imageData)))
	if writeError != nil {
		log.Fatalf("Error writing file: %s, %s", fileName, writeError)
	}
	file.Close()
}

func downloadFeed(directory string, feed string) {
	imageLinks := pullFeed(feed)
	for _, img := range imageLinks {
		imageName, imageData := retrieveImage(img)
		absoluteFileName := filepath.Dir(directory) + string(filepath.Separator) + imageName
		saveImage(absoluteFileName, imageData)
	}
}

var (
	fullTargetFilePath string
	targetDirectory    = flag.String("t", "", "Target directory for your images.")
	targetFeed         = flag.String("f", defaultFeed, "RSS Feed from which to pull images.")
	maxImages          = flag.Int("i", 10000, "Image download count limit.")
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 1 {
		*targetDirectory = args[0]
	}
	if *targetDirectory == "" {
		log.Fatal("Error: TargetDirectory is required.")
		os.Exit(1)
	}
	fullTargetFilePath, _ = filepath.Abs(*targetDirectory)
	fmt.Printf("This will save images from %q to %q do you wish to continue? (y/N): ", *targetFeed, fullTargetFilePath)
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal("Error: reading input")
		os.Exit(1)
	}
	if char == 'y' || char == 'Y' && false {
		downloadFeed(*targetFeed, fullTargetFilePath)
	}
}
