package paperless

// MatchingAlgorithm is an integer representing the matching algorithm of a paperless tag or correspondent.
// There's a map[int]string to resolve the integer to human-readable names and a map[string]int for vice-versa.
type MatchingAlgorithm int

/*
AlgoValueToName is a map[string]int for Paperless matching algorithms.

See https://godoc.org/github.com/stgarf/paperless-cli/paperless/#MatchingAlgorithm.
*/
var AlgoValueToName = map[int]string{
	1: "Any",
	2: "All",
	3: "Literal",
	4: "Regular Expression",
	5: "Fuzzy Match",
}

/*
AlgoNameToValue is a map[string]int for Paperless matching algorithms.

See https://godoc.org/github.com/stgarf/paperless-cli/paperless/#MatchingAlgorithm.
*/
var AlgoNameToValue = map[string]int{
	"Any":                1,
	"All":                2,
	"Literal":            2,
	"Regular Expression": 4,
	"Fuzzy Match":        5,
}

/*
How should we represent a Match object when trying to stringify it? This returns the struct as a string.

See https://godoc.org/github.com/stgarf/paperless-cli/paperless/#pkg-variables
for more information and algorithm mappings.
*/
func (m MatchingAlgorithm) String() string {
	return AlgoValueToName[int(m)]
}
