package paperless

import (
	"fmt"
	"net/url"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

// Color is a tag color
type Color int

// // nolint: golint
var _ColorValueToName = map[int]string{
	1:  "Regent St Blue",
	2:  "Matisse",
	3:  "Feijoa",
	4:  "Forest Green",
	5:  "Sweet Pink",
	6:  "Alizarin Crimson",
	7:  "Macaroni and Cheese",
	8:  "Flush Orange",
	9:  "Lavender Gray",
	10: "Royal Purple",
	11: "Paarl",
	12: "Black",
	13: "Silver",
}

var _ColorNameToValue = map[string]int{
	"RegentStBlue":      1,
	"Matisse":           2,
	"Feijoa":            3,
	"ForestGreen":       4,
	"SweetPink":         5,
	"AlizarinCrimson":   6,
	"MacaroniAndCheese": 7,
	"FlushOrange":       8,
	"LavenderGray":      9,
	"RoyalPurple":       10,
	"Paarl":             11,
	"Black":             12,
	"Silver":            13,
}

func (c Color) String() string {
	return _ColorValueToName[int(c)]
}

// Tag represents a Paperless tag
type Tag struct {
	ID                int               `json:"id"`
	Slug              string            `json:"slug"`
	Name              string            `json:"name"`
	Color             Color             `json:"colour"`
	Match             string            `json:"match"`
	MatchingAlgorithm MatchingAlgorithm `json:"matching_algorithm"`
	IsInsensitive     bool              `json:"is_insensitive"`
}

func (t Tag) String() string {
	return fmt.Sprintf("ID: %v, Slug: %v, Name: %v, Color: %v, Match: %v, Matching Algorithm: %v, Is Insensitive: %v",
		t.ID, t.Slug, t.Name, t.Color, t.Match, t.MatchingAlgorithm, t.IsInsensitive)
}

// TagList is a list/slice of Tag structs
type TagList []Tag

// GetTags returns a slice of Tag items
func (p Paperless) GetTags() (TagList, error) {
	// A place to store the results
	var t Tag
	var tl TagList

	// Make the request
	p.Root += "/tags/"
	u := fmt.Sprint(p)
	creds := []string{p.Username, p.Password}
	resp, err := MakeGetRequest(creds, u)
	if err != nil {
		log.Errorf("An error occurred making request: %v", err.Error())
	}

	// Start processing the request JSON
	tags := gjson.Get(string(resp), "results").Array()
	totalHave := len(tags)
	totalExpected := gjson.Get(string(resp), "count").Int()
	nextURL := gjson.Get(string(resp), "next").String()

	// Append results so far to TagList tl
	for _, tag := range tags {
		gjson.Unmarshal([]byte(tag.Raw), &t)
		tl = append(tl, t)
	}
	// Check if we have all the results or not
	for totalHave < int(totalExpected) {
		log.Debugf("Have: %v, Wanted: %v, Next URL: %v", totalHave, totalExpected, nextURL)
		vals, err := url.Parse(nextURL)
		if err != nil {
			log.Fatalf("Error occurred parsing to URL: %v", err.Error())
		}
		queryParams := vals.Query()
		p.Root = fmt.Sprintf("%v?page=%v", vals.Path, queryParams.Get("page"))
		u = fmt.Sprint(p)
		log.Debugf("Fetching next page of results at: %v", u)
		resp, err = MakeGetRequest(creds, u)
		if err != nil {
			log.Errorf("An error occurred making request: %v", err.Error())
		}
		tags = gjson.Get(string(resp), "results").Array()
		totalHave += len(tags)
		nextURL = gjson.Get(string(resp), "next").String()
		for _, tag := range tags {
			gjson.Unmarshal([]byte(tag.Raw), &t)
			tl = append(tl, t)
		}
	}
	return tl, nil
}

// GetTag returns a slice of Tags based on the search string
func (p Paperless) GetTag(s string, caseSensitive bool) (TagList, error) {
	// A place to store the results
	var t Tag
	var tl TagList

	// Make the request
	if caseSensitive {
		p.Root += "/tags/?name__contains=" + s
	} else {
		p.Root += "/tags/?name__icontains=" + s
	}
	u := fmt.Sprint(p)
	creds := []string{p.Username, p.Password}
	resp, err := MakeGetRequest(creds, u)
	if err != nil {
		log.Errorf("An error occurred making request: %v", err.Error())
	}

	// Start processing the request JSON
	tags := gjson.Get(string(resp), "results").Array()
	totalHave := len(tags)
	totalExpected := gjson.Get(string(resp), "count").Int()
	nextURL := gjson.Get(string(resp), "next").String()

	// Append results so far to TagList tl
	for _, tag := range tags {
		gjson.Unmarshal([]byte(tag.Raw), &t)
		tl = append(tl, t)
	}
	// Check if we have all the results or not
	for totalHave < int(totalExpected) {
		log.Debugf("Have: %v, Wanted: %v, Next URL: %v", totalHave, totalExpected, nextURL)
		vals, err := url.Parse(nextURL)
		if err != nil {
			log.Fatalf("Error occurred parsing to URL: %v", err.Error())
		}
		queryParams := vals.Query()
		p.Root = fmt.Sprintf("%v?page=%v", vals.Path, queryParams.Get("page"))
		u = fmt.Sprint(p)
		log.Debugf("Fetching next page of results at: %v", u)
		resp, err = MakeGetRequest(creds, u)
		if err != nil {
			log.Errorf("An error occurred making request: %v", err.Error())
		}
		tags = gjson.Get(string(resp), "results").Array()
		totalHave += len(tags)
		nextURL = gjson.Get(string(resp), "next").String()
		for _, tag := range tags {
			gjson.Unmarshal([]byte(tag.Raw), &t)
			tl = append(tl, t)
		}
	}
	return tl, nil
}
