package unify4go

import (
	"regexp"
	"unicode/utf8"
)

var (
	// Len is an alias for the utf8.RuneCountInString function, which returns the number of runes
	// (Unicode code points) in the given string. This function treats erroneous and short
	// encodings as single runes of width 1 byte. It is useful for determining the
	// length of a string in terms of its Unicode characters, rather than bytes,
	// allowing for accurate character counting in UTF-8 encoded strings.
	Len = utf8.RuneCountInString

	// RegexpDupSpaces is a precompiled regular expression that matches one or more consecutive
	// whitespace characters (including spaces, tabs, and newlines). This can be used for tasks
	// such as normalizing whitespace in strings by replacing multiple whitespace characters
	// with a single space, or for validating string formats where excessive whitespace should
	// be trimmed or removed.
	RegexpDupSpaces = regexp.MustCompile(`\s+`)
)
