package paperless

import (
	"fmt"
	"net/url"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

// Document represents a Paperless document
type Document struct {
	ID            int      `json:"id"`
	Correspondent string   `json:"correspondent"`
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	FileType      string   `json:"file_type"`
	Tags          []string `json:"tags"`
	Checksum      string   `json:"checksum"`
	Created       string   `json:"created"`
	Modified      string   `json:"modified"`
	Added         string   `json:"added"`
	FileName      string   `json:"file_name"`
	DownloadURL   string   `json:"download_url"`
	ThumbnailURL  string   `json:"thumbnail_url"`
}

func (d Document) String() string {
	return fmt.Sprintf("ID: %v, Correspondent: %v, Title: %v, FileType: %v, "+
		"Tags: %v, Checksum: %v, Created: %v, Modified: %v, Added: %v, FileName: "+
		"%v, DownloadUrl: %v, ThumbnailUrl: %v", d.ID, d.Correspondent, d.Title,
		d.FileType, d.Tags, d.Checksum, d.Created, d.Modified, d.Added,
		d.FileName, d.DownloadURL, d.ThumbnailURL)
}

// DocumentList is a list/slice of Document structs
type DocumentList []Document

// GetDocuments returns a slice of Document items
func (p Paperless) GetDocuments() (DocumentList, error) {
	// A place to store the results
	var document Document
	var docList DocumentList

	// Make the request
	p.Root += "/documents/"
	u := fmt.Sprint(p)
	results, err := p.MakeGetRequest(u)
	if err != nil {
		log.Errorf("An error occurred making request: %v", err.Error())
	}

	// Append results so far to DocumentList docList
	for _, doc := range results {
		gjson.Unmarshal([]byte(doc.Raw), &document)
		// For each doc, resolve it's correspondent and tag names
		// instead of Paperless API urls
		idList := []string{}
		correspondentID := p.GetNameByID(getPath(document.Correspondent))
		for _, tag := range document.Tags {
			tagID := p.GetNameByID(getPath(tag))
			idList = append(idList, tagID)
		}
		document.Correspondent = correspondentID
		document.Tags = idList
		docList = append(docList, document)
	}
	return docList, nil
}

// GetDocument returns a slice of Documents based on the search string
func (p Paperless) GetDocument(s string, caseSensitive bool) (DocumentList, error) {
	// A place to store the results
	var document Document
	var docList DocumentList

	// Make the request
	if caseSensitive {
		p.Root += "/documents/?title__contains=" + s
	} else {
		p.Root += "/documents/?title__icontains=" + s
	}
	u := fmt.Sprint(p)
	results, err := p.MakeGetRequest(u)
	if err != nil {
		log.Errorf("An error occurred making request: %v", err.Error())
	}

	// Append results so far to DocumentList docList
	for _, doc := range results {
		gjson.Unmarshal([]byte(doc.Raw), &document)
		// For each doc, resolve it's correspondents and doc
		idList := []string{}
		correspondentID := p.GetNameByID(getPath(document.Correspondent))
		for _, tag := range document.Tags {
			tagID := p.GetNameByID(getPath(tag))
			idList = append(idList, tagID)
		}
		document.Correspondent = correspondentID
		document.Tags = idList
		docList = append(docList, document)
	}
	return docList, nil
}

func getPath(s string) string {
	url, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	return url.Path
}
