package paperless

import (
	"encoding/json"
	"fmt"
	"log"
)

// TagResults respresents the result of an API call after unmarshaling
type TagResults struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Tags     []Tag  `json:"results"`
}

// Tag represents a Paperless tag
type Tag struct {
	ID                int    `json:"id"`
	Slug              string `json:"slug"`
	Name              string `json:"name"`
	Color             int    `json:"colour"`
	Match             string `json:"match"`
	MatchingAlgorithm int    `json:"matching_algorithm"`
	IsInsensitive     bool   `json:"is_insensitive"`
}

func (t Tag) String() string {
	return fmt.Sprintf("ID: %v, Slug: %v, Name: %v, Color: %v, Match: %v, Matching Algorithm: %v, Is Insensitive: %v",
		t.ID, t.Slug, t.Name, t.Color, t.Match, t.MatchingAlgorithm, t.IsInsensitive)
}

// GetTags returns a slice of Tag items
func (p Paperless) GetTags() ([]Tag, error) {
	p.Root += "/tags"
	tagData, err := p.MakeRequest("GET")
	if err != nil {
		log.Fatalln(err)
	}
	tags := TagResults{}
	json.Unmarshal(tagData, &tags)
	return tags.Tags, nil
}

// GetTag returns a slice of Tags based on the search string
func (p Paperless) GetTag(s string, caseSensitive bool) ([]Tag, error) {
	if caseSensitive {
		p.Root += "/tags/?name__contains=" + s
	} else {
		p.Root += "/tags/?name__icontains=" + s
	}
	tagData, err := p.MakeRequest("GET")
	if err != nil {
		log.Fatalln(err)
	}
	tags := TagResults{}
	json.Unmarshal(tagData, &tags)
	return tags.Tags, nil
}

// // Color is a tag color
// type Color struct {
// 	Value int
// }

// // nolint: golint
// var (
// 	ColorRegentStBlue      = Color{1}
// 	ColorMatisse           = Color{2}
// 	ColorFeijoa            = Color{3}
// 	ColorForestGree        = Color{4}
// 	ColorSweetPink         = Color{5}
// 	ColorAlizarinCrimson   = Color{6}
// 	ColorMacaroniAndCheese = Color{7}
// 	ColorFlushOrange       = Color{8}
// 	ColorLavenderGray      = Color{9}
// 	ColorRoyalPurple       = Color{10}
// 	ColorPaarl             = Color{11}
// 	ColorBlack             = Color{12}
// 	ColorSilver            = Color{13}
// )

// // MatchingAlgorithm are Tag matching algorithms
// type MatchingAlgorithm struct {
// 	Value int `json:"matching_algorithm"`
// }

// // nolint: golint
// var (
// 	MatchingAlgorithmAny               = MatchingAlgorithm{1}
// 	MatchingAlgorithmAll               = MatchingAlgorithm{2}
// 	MatchingAlgorithmLiteral           = MatchingAlgorithm{3}
// 	MatchingAlgorithmRegularExpression = MatchingAlgorithm{4}
// 	MatchingAlgorithmFuzzyMatch        = MatchingAlgorithm{5}
// )
