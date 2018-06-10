package main

import (
	"strings"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	fullTargetFilePath string
	targetDirectory    = flag.String("t", "", "Target directory for your images.")
	targetFeed         = flag.String("f", DefaultFeed, "RSS Feed from which to pull images.")
	maxImages          = flag.Int("i", 10000, "Image download count limit.")
)

// DefaultFeed is the nasa image of the day rss.
const DefaultFeed = "https://www.nasa.gov/rss/dyn/lg_image_of_the_day.rss"

func main() {
	if parseArgs() {
		fullTargetFilePath, _ = filepath.Abs(*targetDirectory)
		if confirmFullFilePathAndFeed(fullTargetFilePath, *targetFeed) {
			feed := RetrieveResource(*targetFeed, "feed")
			enclosures := EnclosureUrlsFromRssBytes(feed.Data)

			for _, enc := range enclosures[0:*maxImages] {
				resource := RetrieveResource(enc.URL, enc.Type)
				if err := resource.SaveResource(fullTargetFilePath); err != nil {
					log.Fatalf("Error saving resource: %s", err)
				}
			}
		}
	}
}

func parseArgs() bool {
	flag.Parse()
	args := flag.Args()
	if len(args) == 1 {
		*targetDirectory = args[0]
	}
	if *targetDirectory == "" {
		log.Fatal("Error: TargetDirectory is required.")
		return false
	}
	return true
}

func confirmFullFilePathAndFeed(path, feed string) bool {
	fmt.Printf("This will save images from %q to %q do you wish to continue? (y/N): ", feed, path)
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal("Error: reading input")
	}

	return strings.ToUpper(string(char)) == "Y"
}
