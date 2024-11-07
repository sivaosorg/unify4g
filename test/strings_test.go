package example_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/sivaosorg/unify4g"
)

type defaultTestStruct struct {
	summary        string
	input          interface{}
	expectedOutput interface{}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"non-empty string", "not empty", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unify4g.IsEmpty(tt.input); got != tt.expected {
				t.Errorf("IsEmpty(%q) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestIsNotEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", false},
		{"non-empty string", "not empty", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unify4g.IsNotEmpty(tt.input); got != tt.expected {
				t.Errorf("IsNotEmpty(%q) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestFormatArg(t *testing.T) {
	tests := []struct {
		arg      string
		expected string
	}{
		{"", "``"},
		{"x", "`x`"},
		{"x\nz", "\"x\\nz\""},
	}
	for i, test := range tests {
		result := unify4g.Quote(test.arg)
		if result != test.expected {
			t.Errorf("Test %d: Arg(%s) returned %s. Expected %s.", i, test.arg, result, test.expected)
		}
	}
}

func TestTrimPrefixAll(t *testing.T) {
	tests := []struct {
		s        string
		prefix   string
		expected string
	}{
		{"", "", ""},
		{"a", "x", "a"},
		{"ax", "", "ax"},
		{"ax", "x", "ax"},
		{"", "x", ""},
		{"x", "", "x"},
		{"x", "x", ""},
		{"xa", "", "xa"},
		{"xa", "x", "a"},
		{"xxa", "", "xxa"},
		{"xxa", "x", "a"},
		{"xxxa", "", "xxxa"},
		{"xxxa", "x", "a"},
		{"xxxa", "xx", "xa"},
		{"xxxxa", "xx", "a"},
		{"xxxxxa", "xx", "xa"},
	}
	for i, test := range tests {
		result := unify4g.TrimPrefixAll(test.s, test.prefix)
		if result != test.expected {
			t.Errorf("Test %d: TrimPrefixAll(`%s`, `%s`) returned `%s`. Expected `%s`.", i, test.s, test.prefix, result, test.expected)
		}
	}
}

func TestTrimPrefixN(t *testing.T) {
	tests := []struct {
		s        string
		prefix   string
		n        int
		expected string
	}{
		{"", "", 1, ""},
		{"a", "x", 1, "a"},
		{"ax", "", 1, "ax"},
		{"ax", "x", 1, "ax"},
		{"", "x", 1, ""},
		{"x", "", 1, "x"},
		{"x", "x", 1, ""},
		{"xa", "", 1, "xa"},
		{"xa", "x", 1, "a"},
		{"xxa", "", 1, "xxa"},
		{"xxa", "x", 1, "xa"},
		{"xxxa", "", 1, "xxxa"},
		{"xxxa", "x", 1, "xxa"},
		{"xa", "", 2, "xa"},
		{"xa", "x", 2, "a"},
		{"xxa", "", 2, "xxa"},
		{"xxa", "x", 2, "a"},
		{"xxxa", "", 2, "xxxa"},
		{"xxxa", "x", 2, "xa"},
		{"xxxa", "xx", 1, "xa"},
		{"xxxxa", "xx", 1, "xxa"},
		{"xxxxa", "xx", 2, "a"},
		{"xxxxxa", "xx", 0, "xxxxxa"},
		{"xxxxxa", "xx", 1, "xxxa"},
		{"xxxxxa", "xx", 2, "xa"},
		{"xxxxxa", "xx", 3, "xa"},
		{"xxxxxa", "xx", -1, "xa"},
	}
	for i, test := range tests {
		result := unify4g.TrimPrefixN(test.s, test.prefix, test.n)
		if result != test.expected {
			t.Errorf("Test %d: TrimPrefixAll(`%s`, `%s`, %d) returned `%s`. Expected `%s`.", i, test.s, test.prefix, test.n, result, test.expected)
		}
	}
}

func TestTrimSuffixAll(t *testing.T) {
	tests := []struct {
		s        string
		suffix   string
		expected string
	}{
		{"", "", ""},
		{"a", "x", "a"},
		{"ax", "", "ax"},
		{"xa", "x", "xa"},
		{"", "x", ""},
		{"x", "", "x"},
		{"x", "x", ""},
		{"ax", "", "ax"},
		{"ax", "x", "a"},
		{"axx", "", "axx"},
		{"axx", "x", "a"},
		{"axxx", "", "axxx"},
		{"axxx", "x", "a"},
		{"axxx", "xx", "ax"},
		{"axxxx", "xx", "a"},
		{"axxxxx", "xx", "ax"},
	}
	for i, test := range tests {
		result := unify4g.TrimSuffixAll(test.s, test.suffix)
		if result != test.expected {
			t.Errorf("Test %d: TrimSuffixAll(`%s`, `%s`) returned `%s`. Expected `%s`.", i, test.s, test.suffix, result, test.expected)
		}
	}
}

func TestTrimSuffixN(t *testing.T) {
	tests := []struct {
		s        string
		suffix   string
		n        int
		expected string
	}{
		{"", "", 1, ""},
		{"a", "x", 1, "a"},
		{"xa", "", 1, "xa"},
		{"xa", "x", 1, "xa"},
		{"", "x", 1, ""},
		{"x", "", 1, "x"},
		{"x", "x", 1, ""},
		{"ax", "", 1, "ax"},
		{"ax", "x", 1, "a"},
		{"axx", "", 1, "axx"},
		{"axx", "x", 1, "ax"},
		{"axxx", "", 1, "axxx"},
		{"axxx", "x", 1, "axx"},
		{"ax", "", 2, "ax"},
		{"ax", "x", 2, "a"},
		{"axx", "", 2, "axx"},
		{"axx", "x", 2, "a"},
		{"axxx", "", 2, "axxx"},
		{"axxx", "x", 2, "ax"},
		{"axxx", "xx", 1, "ax"},
		{"axxxx", "xx", 1, "axx"},
		{"axxxx", "xx", 2, "a"},
		{"axxxxx", "xx", 0, "axxxxx"},
		{"axxxxx", "xx", 1, "axxx"},
		{"axxxxx", "xx", 2, "ax"},
		{"axxxxx", "xx", 3, "ax"},
		{"axxxxx", "xx", -1, "ax"},
	}
	for i, test := range tests {
		result := unify4g.TrimSuffixN(test.s, test.suffix, test.n)
		if result != test.expected {
			t.Errorf("Test %d: TrimSuffixN(`%s`, `%s`, %d) returned `%s`. Expected `%s`.", i, test.s, test.suffix, test.n, result, test.expected)
		}
	}
}

func TestTrimSequenceAll(t *testing.T) {
	tests := []struct {
		s        string
		sequence string
		expected string
	}{
		{"", "", ""},
		{"a", "x", "a"},
		{"ax", "", "ax"},
		{"ax", "x", "a"},
		{"", "x", ""},
		{"x", "", "x"},
		{"x", "x", ""},
		{"xa", "", "xa"},
		{"xa", "x", "a"},
		{"xxa", "", "xxa"},
		{"xxa", "x", "a"},
		{"xxxa", "", "xxxa"},
		{"xxxa", "x", "a"},
		{"xxxa", "xx", "xa"},
		{"xxxxa", "xx", "a"},
		{"xxxxxa", "xx", "xa"},

		{"", "", ""},
		{"a", "x", "a"},
		{"ax", "", "ax"},
		{"xa", "x", "a"},
		{"", "x", ""},
		{"x", "", "x"},
		{"x", "x", ""},
		{"ax", "", "ax"},
		{"ax", "x", "a"},
		{"axx", "", "axx"},
		{"axx", "x", "a"},
		{"axxx", "", "axxx"},
		{"axxx", "x", "a"},
		{"axxx", "xx", "ax"},
		{"axxxx", "xx", "a"},
		{"axxxxx", "xx", "ax"},

		{"xxxaxxxxx", "xx", "xax"},
		{"xxxxxaxxx", "xx", "xax"},
		{"xxxxaxx", "xx", "a"},
		{"xx", "xx", ""},
	}
	for i, test := range tests {
		result := unify4g.TrimSequenceAll(test.s, test.sequence)
		if result != test.expected {
			t.Errorf("Test %d: TrimSequenceAll(`%s`, `%s`) returned `%s`. Expected `%s`.", i, test.s, test.sequence, result, test.expected)
		}
	}
}

func TestReplaceAllStrings(t *testing.T) {
	tests := []struct {
		s        []string
		old      string
		new      string
		expected []string
	}{
		{nil, "", "", []string{}},
		{[]string{}, "", "", []string{}},
		{[]string{"x"}, "x", "y", []string{"y"}},
		{[]string{"in", "if"}, "i", "o", []string{"on", "of"}},
	}
	for i, test := range tests {
		result := unify4g.ReplaceAllStrings(test.s, test.old, test.new)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test %d: RangeReplaceAll(%v, `%s`, `%s`) returned %v. Expected %v.", i, test.s, test.old, test.new, result, test.expected)
		}
	}
}

func TestHash(t *testing.T) {
	cases := []defaultTestStruct{
		{"Normal Test", "Handy", "E80649A6418B6C24FCCB199DAB7CB5BD6EC37593EA0285D52D717FCC7AEE5FB3"},
		{"string with number", "123456", "8D969EEF6ECAD3C29A3A629280E686CF0C3F5D5A86AFF3CA12020C923ADC6C92"},
		{"mashup", "Handy12345", "C82333DB3A6D91F98BE188C6C7B928DF4960B9EC3F3EB8CB50293368C673BE3D"},
		{"with symbols", "#handy_12Ax", "507512071AAEA24A94ECBB0F32EE74169FD59160EE9232819C504F39656E61F7"},
	}
	for _, tc := range cases {
		t.Run(tc.summary, func(t *testing.T) {
			r := unify4g.Hash(tc.input.(string))

			if r != strings.ToLower(tc.expectedOutput.(string)) {
				t.Errorf("Test has failed!\n\tInput: %s,\n\tExpected: %d, \n\tGot: %s", tc.input, tc.expectedOutput, r)
			}
		})
	}
}

// TestOnlyDigits tests the OnlyDigits function with various input strings.
func TestOnlyDigits(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"123abc456", "123456"},      // Mixed digits and letters
		{"abc!@#", ""},               // No digits
		{"   123  456   ", "123456"}, // Spaces around digits
		{"", ""},                     // Empty string
		{"  ", ""},                   // Whitespace-only string
		{"9876543210", "9876543210"}, // Digits-only string
		{"12.34-56", "123456"},       // Digits with special characters
	}

	for _, tt := range tests {
		result := unify4g.OnlyDigits(tt.input)
		if result != tt.expected {
			t.Errorf("OnlyDigits(%q) = %q; expected %q", tt.input, result, tt.expected)
		}
	}
}

func TestIndent(t *testing.T) {
	tests := []struct {
		left     string
		input    string
		expected string
	}{

		{"\t", "Lorem ipsum dolor sit amet", "\tLorem ipsum dolor sit amet"},
		{"\t", "Lorem ipsum\ndolor sit amet\n", "\tLorem ipsum\n\tdolor sit amet\n\t"},
		{"\t", "", "\t"},
		{"", "Lorem", "Lorem"},
	}
	for _, test := range tests {
		output := unify4g.Indent(test.input, test.left)
		if output != test.expected {
			t.Errorf("OnlyDigits(%q) = %q; expected %q", test.input, output, test.expected)
		}
	}
}

func TestAbbreviate(t *testing.T) {
	tests := []struct {
		input    string
		maxWidth int
		expected string
	}{
		// Test case 1: No abbreviation needed, string is shorter than maxWidth
		{
			input:    "Hello",
			maxWidth: 10,
			expected: "Hello",
		},
		// Test case 2: Exact maxWidth, no abbreviation needed
		{
			input:    "Hello",
			maxWidth: 5,
			expected: "Hello",
		},
		// Test case 3: Abbreviation required, string is longer than maxWidth
		{
			input:    "This is a long string",
			maxWidth: 10,
			expected: "This is...",
		},
		// Test case 4: Abbreviation with a very small maxWidth
		{
			input:    "This is a long string",
			maxWidth: 5,
			expected: "Th...",
		},
		// Test case 5: Empty string
		{
			input:    "",
			maxWidth: 5,
			expected: "",
		},
		// Test case 6: String exactly at the boundary of abbreviation
		{
			input:    "This is long",
			maxWidth: 12,
			expected: "This is long",
		},
		// Test case 7: String is just at the limit (no abbreviation)
		{
			input:    "This is short",
			maxWidth: 13,
			expected: "This is short",
		},
		// Test case 8: String smaller than 4 and maxWidth smaller than 4
		{
			input:    "abc",
			maxWidth: 3,
			expected: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := unify4g.Abbreviate(tt.input, tt.maxWidth)
			if result != tt.expected {
				t.Errorf("Abbreviate(%q, %d) = %q; expected %q", tt.input, tt.maxWidth, result, tt.expected)
			}
		})
	}
}

func TestAppendIfMissing(t *testing.T) {
	tests := []struct {
		str      string
		suffix   string
		expected string
		suffixes []string
	}{
		{"example", "txt", "exampletxt", nil},                      // Append missing suffix
		{"example.txt", "txt", "example.txt", nil},                 // Suffix already exists
		{"image", "jpg", "imagejpg", nil},                          // Append suffix when missing
		{"report", "csv", "reportcsv", nil},                        // Basic append case
		{"document", "csv", "documentcsv", []string{"csv", "doc"}}, // Multiple suffixes, already ends with one
		{"hello", "o", "hello", nil},                               // Edge case: ends with same letter
		{"", "suffix", "", nil},                                    // Empty string
	}

	for _, tt := range tests {
		result := unify4g.AppendIfMissing(tt.str, tt.suffix, tt.suffixes...)
		if result != tt.expected {
			t.Errorf("AppendIfMissing(%q, %q) = %q; want %q", tt.str, tt.suffix, result, tt.expected)
		}
	}
}

func TestAppendIfMissingIgnoreCase(t *testing.T) {
	tests := []struct {
		str      string
		suffix   string
		expected string
		suffixes []string
	}{
		{"example", "Txt", "exampleTxt", nil},                // Append case-insensitive suffix
		{"example.txt", "txt", "example.txt", nil},           // Suffix already exists
		{"example.txt", "TXT", "example.txt", nil},           // Case-insensitive check
		{"picture.PNG", "png", "picture.PNG", nil},           // Case-insensitive check with suffix present
		{"report.PDF", "pdf", "report.PDF", nil},             // Case-insensitive check with suffix present
		{"file", "txt", "filetxt", nil},                      // Append suffix when missing
		{"presentation.PPT", "ppt", "presentation.PPT", nil}, // Case-insensitive suffix present
		{"greeting", "ing", "greeting", nil},                 // Ends with case-insensitive suffix
		{"hello", "lo", "hello", nil},                        // Ends with case-insensitive suffix
		{"sample", "suffix", "samplesuffix", nil},            // Case-insensitive append
		{"", "suffix", "", nil},                              // Empty string
	}

	for _, tt := range tests {
		result := unify4g.AppendIfMissingIgnoreCase(tt.str, tt.suffix, tt.suffixes...)
		if result != tt.expected {
			t.Errorf("AppendIfMissingIgnoreCase(%q, %q) = %q; want %q", tt.str, tt.suffix, result, tt.expected)
		}
	}
}
