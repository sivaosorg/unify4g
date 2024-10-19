package unify4go

import (
	"crypto/sha256"
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// IsEmpty checks if the provided string is empty or consists solely of whitespace characters.
//
// The function trims leading and trailing whitespace from the input string `s` using
// strings.TrimSpace. It then evaluates the length of the trimmed string. If the length is
// zero, it indicates that the original string was either empty or contained only whitespace,
// and the function returns true. Otherwise, it returns false.
//
// Parameters:
// - `s`: A string that needs to be checked for emptiness.
//
// Returns:
// - A boolean value:
//   - true if the string is empty or contains only whitespace characters;
//   - false if the string contains any non-whitespace characters.
//
// Example:
//
// result := IsEmpty("   ") // result will be true
// result = IsEmpty("Hello") // result will be false
func IsEmpty(s string) bool {
	trimmed := strings.TrimSpace(s)
	return len(trimmed) == 0
}

// IsNotEmpty checks if the provided string is not empty or does not consist solely of whitespace characters.
//
// This function leverages the IsEmpty function to determine whether the input string `s`
// is empty or contains only whitespace. It returns the negation of the result from IsEmpty.
// If IsEmpty returns true (indicating the string is empty or whitespace), IsNotEmpty will return false,
// and vice versa.
//
// Parameters:
// - `s`: A string that needs to be checked for non-emptiness.
//
// Returns:
// - A boolean value:
//   - true if the string contains at least one non-whitespace character;
//   - false if the string is empty or contains only whitespace characters.
//
// Example:
//
// result := IsNotEmpty("Hello") // result will be true
// result = IsNotEmpty("   ") // result will be false
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// TrimWhitespace removes extra whitespace from the input string,
// replacing any sequence of whitespace characters with a single space.
//
// This function first checks if the input string `s` is empty or consists solely of whitespace
// using the IsEmpty function. If so, it returns an empty string. If the string contains
// non-whitespace characters, it utilizes a precompiled regular expression (regexpDupSpaces)
// to identify and replace all sequences of whitespace characters (including spaces, tabs, and
// newlines) with a single space. This helps to normalize whitespace in the string.
//
// Parameters:
// - `s`: The input string from which duplicate whitespace needs to be removed.
//
// Returns:
//   - A string with all sequences of whitespace characters replaced by a single space.
//     If the input string is empty or only contains whitespace, an empty string is returned.
//
// Example:
//
// result := TrimWhitespace("This   is  an example.\n\nThis is another line.")
// // result will be "This is an example. This is another line."
func TrimWhitespace(s string) string {
	if IsEmpty(s) {
		return ""
	}
	// Use a regular expression to replace all sequences of whitespace characters with a single space.
	s = RegexpDupSpaces.ReplaceAllString(s, " ")
	return s
}

// CleanSpaces removes leading and trailing whitespace characters from a given string and replaces sequences of whitespace characters with a single space.
// It first checks if the input string is empty or consists solely of whitespace characters. If so, it returns an empty string.
// Otherwise, it calls RemoveDuplicateWhitespace to replace all sequences of whitespace characters with a single space, effectively removing duplicates.
// Finally, it trims the leading and trailing whitespace characters from the resulting string using strings.TrimSpace and returns the cleaned string.
func CleanSpaces(s string) string {
	if IsEmpty(s) {
		return ""
	}
	return strings.TrimSpace(TrimWhitespace(s))
}

// Quote formats a string argument for safe output, escaping any special characters
// and enclosing the result in double quotes.
//
// This function uses the fmt.Sprintf function with the %#q format verb to create a quoted
// string representation of the input argument `arg`. The output will escape any special
// characters (such as newlines or tabs) in the string, ensuring that it is suitable for
// safe display or logging. The resulting string will be surrounded by double quotes,
// making it clear where the string begins and ends.
//
// Parameters:
// - `arg`: The input string to be formatted.
//
// Returns:
//   - A string that represents the input `arg` as a quoted string with special characters
//     escaped. This can be useful for creating safe outputs in logs or console displays.
//
// Example:
//
// formatted := Quote("Hello, world!\nNew line here.")
// // formatted will be "\"Hello, world!\\nNew line here.\""
func Quote(arg string) string {
	return fmt.Sprintf("%#q", arg)
}

// TrimPrefixAll returns a new string with all occurrences of prefix at the start of s removed.
// If prefix is the empty string, this function returns s.
func TrimPrefixAll(s string, prefix string) string {
	if IsEmpty(prefix) {
		return s
	}
	for strings.HasPrefix(s, prefix) {
		s = s[len(prefix):]
	}
	return s
}

// TrimPrefixN returns a new string with up to n occurrences of prefix at the start of s removed.
// If prefix is the empty string, this function returns s.
// If n is negative, returns TrimPrefixAll(s, prefix).
func TrimPrefixN(s string, prefix string, n int) string {
	if n < 0 {
		return TrimPrefixAll(s, prefix)
	}
	if IsEmpty(prefix) {
		return s
	}
	for n > 0 && strings.HasPrefix(s, prefix) {
		s = s[len(prefix):]
		n--
	}
	return s
}

// TrimSuffixAll returns a new string with all occurrences of suffix at the end of s removed.
// If suffix is the empty string, this function returns s.
func TrimSuffixAll(s string, suffix string) string {
	if IsEmpty(suffix) {
		return s
	}
	for strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

// TrimSuffixN returns a new string with up to n occurrences of suffix at the end of s removed.
// If suffix is the empty string, this function returns s.
// If n is negative, returns TrimSuffixAll(s, suffix).
func TrimSuffixN(s string, suffix string, n int) string {
	if n < 0 {
		return TrimSuffixAll(s, suffix)
	}
	if IsEmpty(suffix) {
		return s
	}
	for n > 0 && strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
		n--
	}
	return s
}

// TrimSequenceAll returns a new string with all occurrences of sequence at the start and end of s removed.
// If sequence is the empty string, this function returns s.
func TrimSequenceAll(s string, sequence string) string {
	return TrimSuffixAll(TrimPrefixAll(s, sequence), sequence)
}

// ReplaceAllStrings takes a slice of strings and replaces all occurrences of a specified
// substring (old) with a new substring (new) in each string of the slice.
//
// This function creates a new slice of strings, where each string is the result of
// replacing all instances of the old substring with the new substring in the corresponding
// string from the input slice. The original slice remains unchanged.
//
// Parameters:
// - `ss`: A slice of strings in which the replacements will be made.
// - `old`: The substring to be replaced.
// - `new`: The substring to replace the old substring with.
//
// Returns:
//   - A new slice of strings with all occurrences of `old` replaced by `new` in each string
//     from the input slice.
//
// Example:
//
// input := []string{"hello world", "world peace", "goodbye world"}
// output := ReplaceAllStrings(input, "world", "universe")
// // output will be []string{"hello universe", "universe peace", "goodbye universe"}
func ReplaceAllStrings(ss []string, old string, new string) []string {
	values := make([]string, len(ss))
	for i, s := range ss {
		values[i] = strings.ReplaceAll(s, old, new)
	}
	return values
}

// Slash is like strings.Join(elems, "/"), except that all leading and trailing occurrences of '/'
// between elems are trimmed before they are joined together. Non-trailing leading slashes in the
// first element as well as non-leading trailing slashes in the last element are kept.
func Slash(elems ...string) string {
	return JoinUnary(elems, "/")
}

// JoinUnary concatenates a slice of strings into a single string, separating each element
// with a specified separator. The function handles various cases of input size and optimizes
// memory allocation based on expected lengths.
//
// Parameters:
// - `elems`: A slice of strings to be concatenated.
// - `separator`: A string used to separate the elements in the final concatenated string.
//
// Returns:
//   - A single string resulting from the concatenation of the input strings, with the specified
//     separator inserted between each element. If the slice is empty, it returns an empty string.
//     If there is only one element in the slice, it returns that element without any separators.
//
// The function performs the following steps:
//  1. Checks if the input slice is empty; if so, it returns an empty string.
//  2. If the slice contains a single element, it returns that element directly.
//  3. A `strings.Builder` is used to efficiently build the output string, with an initial capacity
//     that is calculated based on the number of elements and their average length.
//  4. Each element is appended to the builder, with the specified separator added between them.
//  5. The function also trims any leading or trailing occurrences of the separator from each element
//     to avoid duplicate separators in the output.
//
// Example:
//
// elems := []string{"apple", "banana", "cherry"}
// separator := ", "
// result := JoinUnary(elems, separator)
// // result will be "apple, banana, cherry"
func JoinUnary(elems []string, separator string) string {
	if len(elems) == 0 {
		return ""
	}
	if len(elems) == 1 {
		return elems[0]
	}
	var sb strings.Builder
	const maxGuess = 100
	const guessAverageElementLen = 5
	if len(elems) <= maxGuess {
		sb.Grow((len(elems)-1)*len(separator) + len(elems)*guessAverageElementLen)
	} else {
		sb.Grow((len(elems)-1)*len(separator) + maxGuess*guessAverageElementLen)
	}
	t := TrimSuffixAll(elems[0], separator) + separator
	for _, element := range elems[1 : len(elems)-1] {
		sb.WriteString(t)
		t = TrimSequenceAll(element, separator) + separator
	}
	sb.WriteString(t)
	t = TrimPrefixAll(elems[len(elems)-1], separator)
	sb.WriteString(t)
	return sb.String()
}

// Reverse returns a new string that is the reverse of the input string s.
// This function handles multi-byte Unicode characters correctly by operating on runes,
// ensuring that each character is reversed without corrupting the character encoding.
//
// Parameters:
// - `s`: A string to be reversed.
//
// Returns:
//   - A new string that contains the characters of the input string in reverse order.
//     If the input string has fewer than two characters (i.e., is empty or a single character),
//     it returns the input string as-is.
//
// The function works as follows:
//  1. It checks the length of the input string using utf8.RuneCountInString. If the string has
//     fewer than two characters, it returns the original string.
//  2. The input string is converted to a slice of runes to correctly handle multi-byte characters.
//  3. A buffer of runes is created to hold the reversed characters.
//  4. A loop iterates over the original rune slice from the end to the beginning, copying each
//     character into the buffer in reverse order.
//  5. Finally, the function converts the buffer back to a string and returns it.
//
// Example:
//
// original := "hello"
// reversed := Reverse(original)
// // reversed will be "olleh"
func Reverse(s string) string {
	if utf8.RuneCountInString(s) < 2 {
		return s
	}
	r := []rune(s)
	buffer := make([]rune, len(r))
	for i, j := len(r)-1, 0; i >= 0; i-- {
		buffer[j] = r[i]
		j++
	}
	return string(buffer)
}

// Hash computes the SHA256 hash of the input string s and returns it as a hexadecimal string.
//
// This function performs the following steps:
//  1. It checks if the input string is empty or consists only of whitespace characters using the IsEmpty function.
//     If the string is empty, it returns the original string.
//  2. It creates a new SHA256 hash using the sha256.New() function.
//  3. The input string is converted to a byte slice and written to the hash. If an error occurs during this process,
//     the function returns an empty string.
//  4. Once the string has been written to the hash, it calculates the final hash value using the Sum method.
//  5. The hash value is then formatted as a hexadecimal string using fmt.Sprintf and returned.
//
// Parameters:
// - `s`: The input string to be hashed.
//
// Returns:
//   - A string representing the SHA256 hash of the input string in hexadecimal format.
//     If the input string is empty or if an error occurs during hashing, an empty string is returned.
//
// Example:
//
// input := "hello"
// hashValue := Hash(input)
// // hashValue will contain the SHA256 hash of "hello" in hexadecimal format.
//
// Notes:
//   - This function is suitable for generating hash values for strings that can be used for comparisons,
//     checksums, or other cryptographic purposes. However, if the input string is empty, it returns the empty
//     string as a direct response.
func Hash(s string) string {
	// Check if the input string is empty or consists solely of whitespace characters
	if IsEmpty(s) {
		return s
	}
	// Create a new SHA256 hash
	h := sha256.New()
	// Write the input string to the hash
	if _, err := h.Write([]byte(s)); err != nil {
		// If an error occurs during the write process, return an empty string
		return ""
	}
	// Sum the hash to get the final hash value
	sum := h.Sum(nil)
	// Format the hash value as a hexadecimal string and return it
	return fmt.Sprintf("%x", sum)
}

// OnlyLetters returns a new string containing only the letters from the original string, excluding all non-letter characters such as numbers, spaces, and special characters.
// This function iterates through each character in the input string, checks if it is a letter using the unicode.IsLetter function, and appends it to a slice of runes if it is.
// The function returns a string created from the slice of letters.
func OnlyLetters(sequence string) string {
	// Check if the input string is empty or consists solely of whitespace characters
	if IsEmpty(sequence) {
		return ""
	}
	// Check if the input string has no runes (e.g., it contains only whitespace characters)
	if utf8.RuneCountInString(sequence) == 0 {
		return ""
	}
	// Initialize a slice to store the letters found in the input string
	var letters []rune
	// Iterate through each character in the input string
	for _, r := range sequence {
		// Check if the current character is a letter
		if unicode.IsLetter(r) {
			// If it is a letter, append it to the slice of letters
			letters = append(letters, r)
		}
	}
	// Convert the slice of letters back into a string and return it
	return string(letters)
}

// OnlyDigits returns a new string containing only the digits from the original string, excluding all non-digit characters such as letters, spaces, and special characters.
// This function first checks if the input string is empty or consists solely of whitespace characters. If so, it returns an empty string.
// If the input string is not empty, it uses a regular expression to replace all non-digit characters with an empty string, effectively removing them.
// The function returns the resulting string, which contains only the digits from the original string.
func OnlyDigits(sequence string) string {
	if IsEmpty(sequence) {
		return ""
	}
	if utf8.RuneCountInString(sequence) > 0 {
		re, _ := regexp.Compile(`[\D]`)
		sequence = re.ReplaceAllString(sequence, "")
	}
	return sequence
}

// Indent takes a string `s` and a string `left`, and indents every line in `s` by prefixing it with `left`.
// Empty lines are also indented.
//
// Parameters:
// - `s`: The input string whose lines will be indented. It may contain multiple lines separated by newline characters (`\n`).
// - `left`: The string that will be used as the indentation prefix. This string is prepended to every line of `s`, including empty lines.
//
// Behavior:
// - The function works by replacing each newline character (`\n`) in `s` with a newline followed by the indentation string `left`.
// - It also adds `left` to the beginning of the string, ensuring the first line is indented.
// - Empty lines, if present, are preserved and indented like non-empty lines.
//
// Returns:
// - A new string where every line of the input `s` has been indented by the string `left`.
//
// Example:
//
// Input:
// s = "Hello\nWorld\n\nThis is a test"
// left = ">>> "
//
// Output:
// ">>> Hello\n>>> World\n>>> \n>>> This is a test"
//
// In this example, each line of the input, including the empty line, is prefixed with ">>> ".
func Indent(s string, left string) string {
	return left + strings.Replace(s, "\n", "\n"+left, -1)
}

// RemoveAccents removes accents and diacritics from the input string s,
// converting special characters into their basic ASCII equivalents.
//
// This function processes each rune in the input string and uses the
// normalizeRune function to convert accented characters to their unaccented
// counterparts. The results are collected in a strings.Builder for efficient
// string concatenation.
//
// Parameters:
// - `s`: The input string from which accents and diacritics are to be removed.
//
// Returns:
//   - A new string that contains the same characters as the input string,
//     but with all accents and diacritics removed. Characters that do not
//     have a corresponding unaccented equivalent are returned as they are.
//
// Example:
//
// input := "Café naïve"
// output := RemoveAccents(input)
// // output will be "Cafe naive"
//
// Notes:
//   - This function is useful for normalizing strings for comparison,
//     searching, or displaying in a consistent format. It relies on the
//     normalizeRune function to perform the actual character conversion.
func RemoveAccents(s string) string {
	var buff strings.Builder
	buff.Grow(len(s))
	for _, r := range s {
		buff.WriteString(normalize_rune(r))
	}
	return buff.String()
}

// Slugify converts a string to a slug which is useful in URLs, filenames.
// It removes accents, converts to lower case, remove the characters which
// are not letters or numbers and replaces spaces with "-".
//
// Example:
//
//	unify4go.Slugify("'We löve Motörhead'") //Output: we-love-motorhead
//
// Normalzation is done with unify4go.ReplaceAccents function using a rune replacement map
// You can use the following code for better normalization before unify4go.Slugify()
//
//	str := "'We löve Motörhead'"
//	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
//	str = transform.String(t, str) //We love Motorhead
//
// Slugify doesn't support transliteration. You should use a transliteration
// library before Slugify like github.com/rainycape/unidecode
//
// Example:
//
//	import "github.com/rainycape/unidecode"
//
//	str := unidecode.Unidecode("你好, world!")
//	unify4go.Slugify(str) //Output: ni-hao-world
func Slugify(s string) string {
	return SlugifySpecial(s, "-")
}

// SlugifySpecial converts a string to a slug with the delimiter.
// It removes accents, converts string to lower case, remove the characters
// which are not letters or numbers and replaces spaces with the delimiter.
//
// Example:
//
//	unify4go.SlugifySpecial("'We löve Motörhead'", "-") //Output: we-love-motorhead
//
// SlugifySpecial doesn't support transliteration. You should use a transliteration
// library before SlugifySpecial like github.com/rainycape/unidecode
//
// Example:
//
//	import "github.com/rainycape/unidecode"
//
//	str := unidecode.Unidecode("你好, world!")
//	unify4go.SlugifySpecial(str, "-") //Output: ni-hao-world
func SlugifySpecial(str string, delimiter string) string {
	str = RemoveAccents(str)
	delBytes := []byte(delimiter)
	n := make([]byte, 0, len(str))
	isPrevSpace := false
	for _, r := range str {
		if r >= 'A' && r <= 'Z' {
			r -= 'A' - 'a'
		}
		//replace non-alpha chars with delimiter
		switch {
		case (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9'):
			n = append(n, byte(int8(r)))
			isPrevSpace = false
		case !isPrevSpace:
			if len(n) > 0 {
				n = append(n, delBytes...)
			}
			fallthrough
		default:
			isPrevSpace = true
		}
	}
	ln := len(n)
	ld := len(delimiter)
	if ln >= ld && string(n[ln-ld:]) == delimiter {
		n = n[:ln-ld]
	}
	return string(n)
}

// ToSnakeCase converts the input string s to snake_case format,
// where all characters are lowercase and spaces are replaced with underscores.
//
// This function first trims any leading or trailing whitespace from the input
// string and then converts all characters to lowercase. It subsequently
// replaces all spaces in the string with underscores to achieve the desired
// snake_case format.
//
// Parameters:
// - `s`: The input string to be converted to snake_case.
//
// Returns:
//   - A new string formatted in snake_case. If the input string is empty or
//     contains only whitespace, the function will return an empty string.
//
// Example:
//
// input := "Hello World"
// output := ToSnakeCase(input)
// // output will be "hello_world"
//
// Notes:
//   - This function is useful for generating variable names, file names,
//     or other identifiers that conform to snake_case naming conventions.
func ToSnakeCase(s string) string {
	s = strings.TrimSpace(strings.ToLower(s))
	return strings.Replace(s, " ", "_", -1)
}

// ToCamelCase converts the input string s to CamelCase format,
// where the first letter of each word is capitalized and all spaces
// are removed.
//
// This function first trims any leading or trailing whitespace from the input
// string. It then iterates over each character in the string, capitalizing the
// first character of each word (defined as a sequence of characters following
// a space) while removing all spaces from the final result. The first character
// of the string remains unchanged unless it follows a space.
//
// Parameters:
// - `s`: The input string to be converted to CamelCase.
//
// Returns:
//   - A new string formatted in CamelCase. If the input string has fewer than
//     two characters, it returns the original string unchanged. If the input
//     string contains only spaces, it returns an empty string.
//
// Example:
//
// input := "hello world"
// output := ToCamelCase(input)
// // output will be "HelloWorld"
//
// Notes:
//   - This function is useful for generating variable names or identifiers that
//     conform to CamelCase naming conventions.
func ToCamelCase(s string) string {
	s = strings.TrimSpace(s)
	if Len(s) < 2 {
		return s
	}
	var buff strings.Builder
	var prev string
	for _, r := range s {
		c := string(r)
		if c != " " {
			if prev == " " {
				c = strings.ToUpper(c)
			}
			buff.WriteString(c)
		}
		prev = c
	}
	return buff.String()
}

// SplitCamelCase splits a CamelCase string into its component words.
//
// This function takes a string in CamelCase format and separates it into
// individual words based on transitions between upper and lower case letters,
// as well as transitions between letters and digits. It handles the following cases:
//
//   - A transition from a lowercase letter to an uppercase letter indicates the
//     start of a new word.
//   - A transition from an uppercase letter to a lowercase letter indicates
//     the continuation of a word, unless preceded by a digit.
//   - A digit following a letter also indicates a split between words.
//
// The function also trims any leading or trailing whitespace from the input string.
// If the input string has fewer than two characters, it returns a slice containing
// the original string.
//
// Parameters:
// - `s`: The input CamelCase string to be split into words.
//
// Returns:
//   - A slice of strings containing the individual words extracted from the
//     input string.
//
// Example:
//
// input := "CamelCaseString123"
// output := SplitCamelCase(input)
// // output will be []string{"Camel", "Case", "String", "123"}
//
// Notes:
//   - This function is useful for parsing identifiers or names that follow
//     CamelCase conventions, making them easier to read and understand.
func SplitCamelCase(s string) []string {
	s = strings.TrimSpace(s)
	if Len(s) < 2 {
		return []string{s}
	}
	var prev rune
	var start int
	words := []string{}
	runes := []rune(s)
	for i, r := range runes {
		if i != 0 {
			switch {
			case unicode.IsDigit(r) && unicode.IsLetter(prev):
				fallthrough
			case unicode.IsUpper(r) && unicode.IsLower(prev):
				words = append(words, string(runes[start:i]))
				start = i
			case unicode.IsLower(r) && unicode.IsUpper(prev) && start != i-1:
				words = append(words, string(runes[start:i-1]))
				start = i - 1
			}
		}
		prev = r
	}
	if start < len(runes) {
		words = append(words, string(runes[start:]))
	}
	return words
}

// RemovePrefixes removes specified prefixes from the start of a given string.
//
// This function checks the input string `s` and removes any prefixes provided
// in the `prefix` variadic parameter. If the string is empty or if no prefixes
// are provided, the original string is returned unchanged. The function will
// attempt to remove each specified prefix in the order they are provided.
//
// Parameters:
//   - `s`: The input string from which prefixes will be removed.
//   - `prefix`: A variadic parameter that takes one or more prefixes to be removed
//     from the beginning of the string.
//
// Returns:
//   - A string with the specified prefixes removed. If no prefixes are matched,
//     or if the string is empty, the original string is returned.
//
// Example:
//
// input := "prefix_example"
// output := RemovePrefixes(input, "prefix_", "test_")
// // output will be "example"
//
// Notes:
//   - This function is useful for cleaning up strings by removing unwanted or
//     redundant prefixes in various contexts.
func RemovePrefixes(s string, prefixes ...string) string {
	if IsEmpty(s) {
		return s
	}
	if len(prefixes) == 0 {
		return s
	}
	for _, v := range prefixes {
		s = strings.TrimPrefix(s, v)
	}
	return s
}
