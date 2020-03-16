package paperless

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Document is a struct representation of Paperless' /api/documents/<id> JSON response.
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

// How should we represent a Document object when trying to stringify it? This returns the struct as a string.
func (d Document) String() string {
	return fmt.Sprintf("ID: %v, Correspondent: %v, Title: %v, FileType: %v, "+
		"Tags: %v, Checksum: %v, Created: %v, Modified: %v, Added: %v, FileName: "+
		"%v, DownloadUrl: %v, ThumbnailUrl: %v", d.ID, d.Correspondent, d.Title,
		d.FileType, d.Tags, d.Checksum, d.Created, d.Modified, d.Added,
		d.FileName, d.DownloadURL, d.ThumbnailURL)
}

// DocumentList is a slice of
// https://godoc.org/github.com/stgarf/paperless-cli/paperless/#Document structs.
type DocumentList []Document

// GetDocument returns a https://godoc.org/github.com/stgarf/paperless-cli/paperless/#DocumentList matching the search string.
func (p Paperless) GetDocument(s string, caseSensitive bool) (DocumentList, error) {
	// A place to store the results
	var document Document
	var docList DocumentList

	// Build the ID to Name maps
	go p.mapCorrespondents()
	go p.mapTags()

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
		json.Unmarshal([]byte(doc.Raw), &document)
		// For each doc, resolve it's correspondents and doc
		idList := []string{}
		correspondentID := corrIDToName[urlToID(document.Correspondent)]
		for _, tag := range document.Tags {
			tagID := tagIDToName[urlToID(tag)]
			idList = append(idList, tagID)
		}
		document.Correspondent = correspondentID
		document.Tags = idList
		docList = append(docList, document)
	}
	return docList, nil
}

// GetDocuments returns a https://godoc.org/github.com/stgarf/paperless-cli/paperless/#DocumentList.
func (p Paperless) GetDocuments() (DocumentList, error) {
	// A place to store the results
	var document Document
	var docList DocumentList

	// Build the ID to Name maps
	go p.mapCorrespondents()
	go p.mapTags()

	// Make the request
	p.Root += "/documents/"
	u := fmt.Sprint(p)
	results, err := p.MakeGetRequest(u)
	if err != nil {
		log.Errorf("An error occurred making request: %v", err.Error())
	}

	// Append results so far to DocumentList docList
	for _, doc := range results {
		json.Unmarshal([]byte(doc.Raw), &document)
		// For each doc, resolve it's correspondent and tag names
		// instead of Paperless API urls
		idList := []string{}
		correspondentID := corrIDToName[urlToID(document.Correspondent)]
		for _, tag := range document.Tags {
			tagID := tagIDToName[urlToID(tag)]
			idList = append(idList, tagID)
		}
		document.Correspondent = correspondentID
		document.Tags = idList
		docList = append(docList, document)
	}
	return docList, nil
}

// corrIDToName is a map[int]string of Paperless correspondents to their id number.
var corrIDToName map[int]string

// tagIDToName is a map[int]string of Paperless tags to their id number.
var tagIDToName map[int]string

// mapCorrespondents calls (paperless).GetCorrespondents() and populates the corrIDToName map.
func (p Paperless) mapCorrespondents() {
	corrIDToName = make(map[int]string)
	p.Root = "/api"
	list, err := p.GetCorrespondents()
	if err != nil {
		log.Panicln(err)
	}
	for _, c := range list {
		corrIDToName[c.ID] = c.Name
	}
}

// mapTags calls (paperless).GetTags() and populates the tagIDToName map.
func (p Paperless) mapTags() {
	tagIDToName = make(map[int]string)
	p.Root = "/api"
	list, err := p.GetTags()
	if err != nil {
		log.Panicln(err)
	}
	for _, c := range list {
		tagIDToName[c.ID] = c.Name
	}
}

// urlToID takes a Paperless Tag or Correspondet ID url and
// returns the ID off the end of it.
func urlToID(s string) int {
	r := regexp.MustCompile("/([0-9]+)/")
	i, _ := strconv.ParseInt(strings.Trim(r.FindString(s), "/"), 10, 32)
	return int(i)
}
