package paperless

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

// Correspondent is a struct representation of Paperless' /api/correspondents/<id> JSON response.
type Correspondent struct {
	ID                int               `json:"id"`
	Slug              string            `json:"slug"`
	Name              string            `json:"name"`
	Match             string            `json:"match"`
	MatchingAlgorithm MatchingAlgorithm `json:"matching_algorithm"`
	IsInsensitive     bool              `json:"is_insensitive"`
}

// How should we represent a Correspondent object when trying to stringify it? This returns the struct as a string.
func (c Correspondent) String() string {
	return fmt.Sprintf("ID: %v, Slug: %v, Name: %v, Match: %v, Matching Algorithm: %v, Is Insensitive: %v",
		c.ID, c.Slug, c.Name, c.Match, c.MatchingAlgorithm, c.IsInsensitive)
}

// CorrespondentList is a slice of https://godoc.org/github.com/stgarf/paperless-cli/paperless/#Correspondent structs.
type CorrespondentList []Correspondent

// GetCorrespondent returns a https://godoc.org/github.com/stgarf/paperless-cli/paperless/#CorrespondentList matching the search string.
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
	results, err := p.MakeGetRequest(u)
	if err != nil {
		log.Errorf("An error occurred making request: %v", err.Error())
	}

	// Append results so far to CorrespondentList cl
	for _, corr := range results {
		gjson.Unmarshal([]byte(corr.Raw), &c)
		cl = append(cl, c)
	}
	return cl, nil
}

// GetCorrespondents returns a https://godoc.org/github.com/stgarf/paperless-cli/paperless/#CorrespondentList.
func (p Paperless) GetCorrespondents() (CorrespondentList, error) {
	// A place to store the results
	var c Correspondent
	var cl CorrespondentList

	// Make the request
	p.Root += "/correspondents"
	u := fmt.Sprint(p)
	results, err := p.MakeGetRequest(u)
	if err != nil {
		log.Errorf("An error occurred making request: %v", err.Error())
	}

	// Append results so far to CorrespondentList cl
	for _, corr := range results {
		gjson.Unmarshal([]byte(corr.Raw), &c)
		cl = append(cl, c)
	}
	return cl, nil
}
