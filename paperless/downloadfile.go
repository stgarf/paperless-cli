package paperless

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

var selection int

func (p Paperless) writeFile(document Document) {
	downloadURL := p.DownloadString(document.DownloadURL)

	client := http.Client{Timeout: time.Second * 5}
	log.Debugf("downloading from: %v", downloadURL)
	req := ReturnAuthenticatedRequest(p.Username, p.Password)
	req.Method = "GET"
	urlPtr, _ := url.Parse(downloadURL)
	req.URL = urlPtr

	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error downloading file: %v", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(document.FileName)
	if err != nil {
		log.Panicln(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Panicln(err)
	}
}

// DownloadFiles pulls either a paperless document or allows selection from a list
func (p Paperless) DownloadFiles(filename string) {
	documents, err := Paperless.GetDocument(p, filename, true)
	if err != nil {
		log.Errorf("An error occurred: %s", err)
		return
	}
	if len(documents) == 0 {
		log.Errorf("No matches for search string %s", filename)
		return
	}
	if len(documents) > 1 {
		fmt.Printf("%d files found. Please select from:\n", len(documents))
		for index, element := range documents {
			index++
			fmt.Println(strconv.Itoa(index) + ": " + element.FileName)
		}
		selection = -1
		fmt.Scanln(&selection)
		for selection < 1 || selection > len(documents) {
			log.Errorf("%d is out of range of the length of choices", selection)
			fmt.Scanln(&selection)
		}
	}

	document := documents[selection-1]
	p.writeFile(document)
	fmt.Printf("Sucessfully downloaded file with string: %s\n", filename)
}
