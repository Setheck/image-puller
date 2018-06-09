package puller_test

import (
	"path"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/setheck/image-puller/puller"
)

type testData struct {
	data     string
	expected string
}

// func TestEnclosureUrlsFromRss(t *testing.T) {
// 	for _, testElement := range []testData{
// 		{`<rss><channel><item><enclosure url=""/></item></channel></rss>`, ""},
// 		{`<rss><channel><item><enclosure url="testurl"/></item></channel></rss>`, "testurl"},
// 	} {
// 		resultAry := main.EnclosureUrlsFromRss([]byte(testElement.data))
// 		if resultAry[0] != testElement.expected {
// 			t.Errorf("Error")
// 		}

// 	}
// const rssFeed = `<rss><channel><item><enclosure url="testurl"/></item></channel></rss>`

// want := []string { "testurl" }
// if want[0] != enclosureUrlsFromRss([]byte(rssFeed))[0]{
// 	t.Errorf("error")
// }
// }

// func TestMakeRequestReadBytes(t *testing.T) {
// 	const responseString = "<rss></rss>"
// 	testServer := httptest.NewServer(
// 		http.HandlerFunc(
// 			func(w http.ResponseWriter, r *http.Request) {
// 				//w.Header().Set("Content-Type", "application/xml")
// 				fmt.Fprint(w, responseString)
// 			}))
// 	defer testServer.Close()

// 	byteResponse := makeRequestReadBytes(testServer.URL)

// 	fmt.Println(string(byteResponse))
// 	got := string(byteResponse)
// 	if got != responseString {
// 		t.Errorf("Resulting response %q did not match expected %q", got, responseString)
// 	}
// }

func TestRetrieveResources(t *testing.T) {
	testFileName := "test.jpg"
	testFileType := "image"
	testFileData := "some file data blah blah blah"
	testServer := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				//w.Header().Set("Content-Type", "application/xml")
				fmt.Fprint(w, testFileData)
			}))
	defer testServer.Close()

	resource := puller.RetrieveResource(testServer.URL+"/"+testFileName, testFileType)
	if resource.Filename != testFileName {
		t.Errorf("Expected Filename: %q, got %q", testFileName, resource.Filename)
	}
	if resource.Filetype != testFileType {
		t.Errorf("Expected Filetype: %q, got %q", testFileType, resource.Filetype)
	}
	if string(resource.Data) != testFileData {
		t.Errorf("Expected Filetype: %q, got %q", testFileData, resource.Data)
	}
}

func TestSaveResource(t *testing.T) {
	testDirectory := "testdata"
	testFileName := "testfile.jpg"
	testFileData := "some bytes"

	resource := &puller.RemoteResource{
		Filename: testFileName,
		Filetype: "image",
		Data:     []byte(testFileData),
	}

	absoluteFilePath, err := filepath.Abs(testDirectory)
	if err != nil {
		t.Errorf("Error determining filepath: %s", err)
	}

	err = resource.SaveResource(absoluteFilePath)
	if err != nil {
		t.Error(err)
	}
	fileData, err := ioutil.ReadFile(absoluteFilePath)
	if err != nil {
		t.Errorf("Error readinf file: %s", testFileName)
	}
	if string(fileData) != testFileData {
		t.Error("File data was not read correctly!")
	}
}
