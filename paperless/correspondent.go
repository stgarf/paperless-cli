package paperless

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// CorrResults respresents the result of an API call after unmarshaling
type CorrResults struct {
	Count          int             `json:"count"`
	Next           string          `json:"next"`
	Previous       string          `json:"previous"`
	Correspondents []Correspondent `json:"results"`
}

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

// GetCorrespondents returns a slice of Correspondent items
func (p Paperless) GetCorrespondents() ([]Correspondent, error) {
	p.Root += "/correspondents"
	cData, err := p.MakeRequest("GET")
	if err != nil {
		log.Fatalln(err)
	}
	corrs := CorrResults{}
	json.Unmarshal(cData, &corrs)
	// FIXME (sgarf): // We're not fetching all the results, fix this
	if len(corrs.Correspondents) < corrs.Count {
		log.Warnln("We're not done fetching correspondents!!!")
	}
	return corrs.Correspondents, nil
}

// GetCorrespondent returns a slice of Tags based on the search string
func (p Paperless) GetCorrespondent(s string, caseSensitive bool) ([]Correspondent, error) {
	if caseSensitive {
		p.Root += "/correspondents/?name__contains=" + s
	} else {
		p.Root += "/correspondents/?name__icontains=" + s
	}
	cData, err := p.MakeRequest("GET")
	if err != nil {
		log.Fatalln(err)
	}
	corrs := CorrResults{}
	json.Unmarshal(cData, &corrs)
	// FIXME (sgarf): // We're not fetching all the results, fix this
	if len(corrs.Correspondents) < corrs.Count {
		log.Warnln("We're not done fetching correspondents!!!")
	}
	return corrs.Correspondents, nil
}
