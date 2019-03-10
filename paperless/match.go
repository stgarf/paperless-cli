package paperless

// MatchingAlgorithm represents the match algorithm used
type MatchingAlgorithm int

var _AlgoValueToName = map[int]string{
	1: "Any",
	2: "All",
	3: "Literal",
	4: "Regular Expression",
	5: "Fuzzy Match",
}

var _AlgoNameToValue = map[string]int{
	"Any":                1,
	"All":                2,
	"Literal":            2,
	"Regular Expression": 4,
	"Fuzzy Match":        5,
}

func (m MatchingAlgorithm) String() string {
	return _AlgoValueToName[int(m)]
}
