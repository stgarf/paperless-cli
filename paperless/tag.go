package paperless

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

/*
Color is an integer representing the tag color of a paperless tag. There's a
map[int]string to resolve the integer to human-readable names and a map[string]int
for vice-versa.

See https://godoc.org/github.com/stgarf/paperless-cli/paperless/#pkg-variables
for more information and color mappings.
*/
type Color int

/*
ColorValueToName is a map[int]string for Paperless tag colors.

See https://godoc.org/github.com/stgarf/paperless-cli/paperless/#Color
*/
var ColorValueToName = map[int]string{
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

/*
ColorNameToValue is a map[string]int for Paperless tag colors.

See https://godoc.org/github.com/stgarf/paperless-cli/paperless/#Color
*/
var ColorNameToValue = map[string]int{
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

/*
How should we represent a Tag color when trying to stringify it? This returns the struct as a string.

See https://godoc.org/github.com/stgarf/paperless-cli/paperless/#pkg-variables
for more information and color mappings.
*/
func (c Color) String() string {
	return ColorValueToName[int(c)]
}

// Tag is a struct representation of Paperless' /api/tags/<id> JSON response.
type Tag struct {
	ID                int               `json:"id"`
	Slug              string            `json:"slug"`
	Name              string            `json:"name"`
	Color             Color             `json:"colour"`
	Match             string            `json:"match"`
	MatchingAlgorithm MatchingAlgorithm `json:"matching_algorithm"`
	IsInsensitive     bool              `json:"is_insensitive"`
}

// How should we represent a Tag object when trying to stringify it? This returns the struct as a string.
func (t Tag) String() string {
	return fmt.Sprintf("ID: %v, Slug: %v, Name: %v, Color: %v, Match: %v, Matching Algorithm: %v, Is Insensitive: %v",
		t.ID, t.Slug, t.Name, t.Color, t.Match, t.MatchingAlgorithm, t.IsInsensitive)
}

// TagList is a slice of https://godoc.org/github.com/stgarf/paperless-cli/paperless/#Tag structs.
type TagList []Tag

// GetTag returns a https://godoc.org/github.com/stgarf/paperless-cli/paperless/#TagList matching the search string.
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
	results, err := p.MakeGetRequest(u)
	if err != nil {
		log.Errorf("An error occurred making request: %v", err.Error())
	}

	// Append results so far to TagList tl
	for _, tag := range results {
		json.Unmarshal([]byte(tag.Raw), &t)
		tl = append(tl, t)
	}
	return tl, nil
}

// GetTags returns a https://godoc.org/github.com/stgarf/paperless-cli/paperless/#TagList.
func (p Paperless) GetTags() (TagList, error) {
	// A place to store the results
	var t Tag
	var tl TagList

	// Make the request
	p.Root += "/tags/"
	u := fmt.Sprint(p)
	results, err := p.MakeGetRequest(u)
	if err != nil {
		log.Errorf("An error occurred making request: %v", err.Error())
	}

	// Append results so far to TagList tl
	for _, tag := range results {
		json.Unmarshal([]byte(tag.Raw), &t)
		tl = append(tl, t)
	}
	return tl, nil
}
