package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

	resource := RetrieveResource(testServer.URL+"/"+testFileName, testFileType)
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

//TODO: How do i test filesystem in go?
// func TestSaveResource(t *testing.T) {
// 	testDirectory := "testdata"
// 	testFileName := "testfile.jpg"
// 	testFileData := "some bytes"

// 	resource := &puller.RemoteResource{
// 		Filename: testFileName,
// 		Filetype: "image",
// 		Data:     []byte(testFileData),
// 	}

// 	absoluteFilePath, err := filepath.Abs(testDirectory)
// 	if err != nil {
// 		t.Errorf("Error determining filepath: %s", err)
// 	}

// 	err = resource.SaveResource(absoluteFilePath)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fileData, err := ioutil.ReadFile(absoluteFilePath)
// 	if err != nil {
// 		t.Errorf("Error readinf file: %s", testFileName)
// 	}
// 	if string(fileData) != testFileData {
// 		t.Error("File data was not read correctly!")
// 	}
// }
