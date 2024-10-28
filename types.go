package unify4g

type OptionsConfig struct {
	// Width is an max column width for single line arrays
	// Default is 80
	Width int `json:"width"`
	// Prefix is a prefix for all lines
	// Default is an empty string
	Prefix string `json:"prefix"`
	// Indent is the nested indentation
	// Default is two spaces
	Indent string `json:"indent"`
	// SortKeys will sort the keys alphabetically
	// Default is false
	SortKeys bool `json:"sort_keys"`
}

// Style is the color style
type Style struct {
	Key, String, Number [2]string
	True, False, Null   [2]string
	Escape              [2]string
	Brackets            [2]string
	Append              func(dst []byte, c byte) []byte
}

type result int
type byKind int
type jsonType int

type pair struct {
	keyStart, keyEnd     int
	valueStart, valueEnd int
}

type byKeyVal struct {
	sorted bool
	json   []byte
	buf    []byte
	pairs  []pair
}

// DefaultOptionsConfig is the default options for pretty formats.
var DefaultOptionsConfig = &OptionsConfig{Width: 80, Prefix: "", Indent: "  ", SortKeys: false}

// TerminalStyle is for terminals
var TerminalStyle *Style
