package paperless

import (
	"fmt"
	"net/url"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

// Correspondent represents a Paperless correspondent
type Correspondent struct {
	ID                int    `json:"id"`
	Slug              string `json:"slug"`
	Name              string `json:"name"`
	Match             string `json:"match"`
	MatchingAlgorithm int    `json:"matching_algorithm"`
	IsInsensitive     bool   `json:"is_insensitive"`
}

func (c Correspondent) String() string {
	return fmt.Sprintf("ID: %v, Slug: %v, Name: %v, Match: %v, Matching Algorithm: %v, Is Insensitive: %v",
		c.ID, c.Slug, c.Name, c.Match, c.MatchingAlgorithm, c.IsInsensitive)
}

// CorrespondentList is a list/slice of Correspondent
type CorrespondentList []Correspondent

// GetCorrespondents returns a slice of Correspondent items
func (p Paperless) GetCorrespondents() (CorrespondentList, error) {
	// A place to store the results
	var c Correspondent
	var cl CorrespondentList

	// Make the request
	p.Root += "/correspondents"

	u := fmt.Sprint(p)
	creds := []string{p.Username, p.Password}
	resp, err := MakeGetRequest(creds, u)
	if err != nil {
		log.Errorf("An error occurred making request: %v", err.Error())
	}

	// Start processing the request JSON
	corrs := gjson.Get(string(resp), "results").Array()
	totalHave := len(corrs)
	totalExpected := gjson.Get(string(resp), "count").Int()
	nextURL := gjson.Get(string(resp), "next").String()

	// Append results so far to CorrespondentList cl
	for _, corrs := range corrs {
		gjson.Unmarshal([]byte(corrs.Raw), &c)
		cl = append(cl, c)
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
		corrs = gjson.Get(string(resp), "results").Array()
		totalHave += len(corrs)
		nextURL = gjson.Get(string(resp), "next").String()
		for _, corr := range corrs {
			gjson.Unmarshal([]byte(corr.Raw), &c)
			cl = append(cl, c)
		}
	}

	return cl, nil
}

// GetCorrespondent returns a slice of Tags based on the search string
func (p Paperless) GetCorrespondent(s string, caseSensitive bool) (CorrespondentList, error) {
	// A place to store the results
	var c Correspondent
	var cl CorrespondentList

	// Make the request
	if caseSensitive {
		p.Root += "/correspondents/?name__contains=" + s
	} else {
		p.Root += "/correspondents/?name__icontains=" + s
	}
	u := fmt.Sprint(p)
	creds := []string{p.Username, p.Password}
	resp, err := MakeGetRequest(creds, u)
	if err != nil {
		log.Errorf("An error occurred making request: %v", err.Error())
	}

	// Start processing the request JSON
	corrs := gjson.Get(string(resp), "results").Array()
	totalHave := len(corrs)
	totalExpected := gjson.Get(string(resp), "count").Int()
	nextURL := gjson.Get(string(resp), "next").String()

	// Append results so far to CorrespondentList cl
	for _, corr := range corrs {
		gjson.Unmarshal([]byte(corr.Raw), &c)
		cl = append(cl, c)
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
		corrs = gjson.Get(string(resp), "results").Array()
		totalHave += len(corrs)
		nextURL = gjson.Get(string(resp), "next").String()
		for _, corr := range corrs {
			gjson.Unmarshal([]byte(corr.Raw), &c)
			cl = append(cl, c)
		}
	}
	return cl, nil
}
