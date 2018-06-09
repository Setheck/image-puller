package puller

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//RemoteResource a remote resource image/file/etc
type RemoteResource struct {
	Filename string
	Filetype string
	Data     []byte
}

func makeRequestReadBytes(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error retrieving %s", url)
	}
	defer resp.Body.Close()
	return readBytesFromBody(resp)
}

func readBytesFromBody(resp *http.Response) []byte {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body")
	}
	return body
}

//RetrieveResource get a remote resource.
func RetrieveResource(url string, t string) *RemoteResource {
	log.Printf("Retrieiving %q from %q", t, url)
	splitName := strings.Split(url, "/")
	return &RemoteResource{
		Filename: splitName[len(splitName)-1],
		Filetype: t,
		Data:     makeRequestReadBytes(url),
	}
}

//SaveResource save the resource to a file.
func (r *RemoteResource) SaveResource(targetDirectory string) error {
	absoluteFileName := targetDirectory + string(filepath.Separator) + r.Filename
	log.Printf("Attempting to save file %q", absoluteFileName)
	file, err := os.Create(absoluteFileName)
	if err != nil {
		log.Fatalf("Failed creating file: %s", absoluteFileName)
		return err
	}
	_, err = io.Copy(file, ioutil.NopCloser(bytes.NewBuffer(r.Data)))
	if err != nil {
		log.Fatalf("Error writing file: %s, %s", absoluteFileName, err)
		return err
	}
	file.Close()
	return nil
}

//ToStringForm write string form of object.
func (r *RemoteResource) ToStringForm() string {
	return "RemoteResource{\n" +
		"Filename: " + r.Filename +
		",\n Filetype: " + r.Filetype +
		",\n Data: " + string(r.Data) +
		"\n}"
}
