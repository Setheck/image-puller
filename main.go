package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/setheck/image-puller/rssparse"

	"github.com/setheck/image-puller/puller"
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
		feed := puller.RetrieveResource(*targetFeed, "feed")
		enclosures := rssparse.EnclosureUrlsFromRssBytes(feed.Data)

		var saveCount int
		for _, enc := range enclosures {
			if saveCount >= *maxImages {
				break
			}
			resource := puller.RetrieveResource(enc.URL, enc.Type)
			resource.SaveResource(fullTargetFilePath)
			saveCount++
		}
	}
}
