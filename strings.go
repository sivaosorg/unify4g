package unify4go

import (
	"crypto/sha256"
	"fmt"
	"regexp"
	"strconv"
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

// IsAnyEmpty checks if any of the provided strings are empty.
//
// This function takes a variadic number of string arguments and iterates through
// each string to check if it is empty (i.e., has a length of zero or consists
// solely of whitespace characters). If any string in the provided list is found
// to be empty, the function returns `true`. If all strings are non-empty, it
// returns `false`.
//
// Parameters:
//   - `strings`: A variadic parameter that allows passing multiple strings to be checked.
//
// Returns:
//   - `true` if at least one of the provided strings is empty; `false` if all
//     strings are non-empty.
//
// Example:
//
// result := IsAnyEmpty("hello", "", "world")
// // result will be true because one of the strings is empty.
//
// Notes:
//   - The function utilizes the IsEmpty helper function to determine if a string
//     is considered empty, which may include strings that are only whitespace.
func IsAnyEmpty(strings ...string) bool {
	for _, s := range strings {
		if IsEmpty(s) {
			return true
		}
	}
	return false
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

// IsNoneEmpty checks if all provided strings are non-empty.
//
// This function takes a variadic number of string arguments and iterates through
// each string to verify that none of them are empty (i.e., have a length of zero
// or consist solely of whitespace characters). If any string in the provided list
// is found to be empty, the function immediately returns `false`. If all strings
// are non-empty, it returns `true`.
//
// Parameters:
//   - `strings`: A variadic parameter that allows passing multiple strings to be checked.
//
// Returns:
//   - `true` if all of the provided strings are non-empty; `false` if at least one
//     string is empty.
//
// Example:
//
// result := IsNoneEmpty("hello", "world", "!")
// // result will be true because all strings are non-empty.
//
// result2 := IsNoneEmpty("hello", "", "world")
// // result2 will be false because one of the strings is empty.
//
// Notes:
//   - The function utilizes the IsEmpty helper function to determine if a string
//     is considered empty, which may include strings that are only whitespace.
func IsNoneEmpty(strings ...string) bool {
	for _, s := range strings {
		if IsEmpty(s) {
			return false
		}
	}
	return true
}

// IsBlank checks if a string is blank (empty or contains only whitespace).
//
// This function determines if the input string `s` is considered blank. A string
// is considered blank if it is either an empty string or consists solely of
// whitespace characters (spaces, tabs, newlines, etc.).
//
// The function first checks if the string is empty. If it is, it returns `true`.
// If the string is not empty, it uses a regular expression to check if the
// string contains only whitespace characters. If the string matches this
// condition, it also returns `true`. If neither condition is met, the function
// returns `false`, indicating that the string contains non-whitespace characters.
//
// Parameters:
//   - `s`: The input string to check for blankness.
//
// Returns:
//   - `true` if the string is blank (empty or contains only whitespace);
//     `false` otherwise.
//
// Example:
//
// result1 := IsBlank("")
// // result1 will be true because the string is empty.
//
// result2 := IsBlank("   ")
// // result2 will be true because the string contains only spaces.
//
// result3 := IsBlank("Hello")
// // result3 will be false because the string contains non-whitespace characters.
//
// Notes:
//   - The function uses a regular expression to match strings that consist entirely
//     of whitespace. The regex `^\s+$` matches strings that contain one or more
//     whitespace characters from the start to the end of the string.
func IsBlank(s string) bool {
	if s == "" {
		return true
	}
	if regexp.MustCompile(`^\s+$`).MatchString(s) {
		return true
	}
	return false
}

// IsNotBlank checks if a string is not blank (not empty and contains non-whitespace characters).
//
// This function serves as a logical negation of the `IsBlank` function. It checks
// if the input string `s` contains any non-whitespace characters. A string is
// considered not blank if it is neither empty nor consists solely of whitespace
// characters. This is determined by calling the `IsBlank` function and
// negating its result.
//
// Parameters:
//   - `s`: The input string to check for non-blankness.
//
// Returns:
//   - `true` if the string is not blank (contains at least one non-whitespace character);
//     `false` if the string is blank (empty or contains only whitespace).
//
// Example:
//
// result1 := IsNotBlank("Hello")
// // result1 will be true because the string contains non-whitespace characters.
//
// result2 := IsNotBlank("   ")
// // result2 will be false because the string contains only spaces.
//
// result3 := IsNotBlank("")
// // result3 will be false because the string is empty.
//
// Notes:
//   - This function provides a convenient way to check for meaningful content
//     in a string by confirming that it is not blank.
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// IsAnyBlank checks if any of the provided strings are blank (empty or containing only whitespace).
//
// This function iterates through a variadic list of strings and determines if at least
// one of them is considered blank. A string is considered blank if it is either empty
// or consists solely of whitespace characters. The function returns `true` as soon as
// it finds a blank string; if none of the strings are blank, it returns `false`.
//
// Parameters:
//   - `strings`: A variadic parameter that accepts one or more strings to check for blankness.
//
// Returns:
//   - `true` if at least one of the provided strings is blank;
//     `false` if none of the strings are blank.
//
// Example:
//
// result1 := IsAnyBlank("Hello", "World", " ")
// // result1 will be true because the third string is blank (contains only a space).
//
// result2 := IsAnyBlank("Hello", "World")
// // result2 will be false because both strings are not blank.
//
// Notes:
//   - This function is useful for validating input or ensuring that required fields
//     are not left blank in forms or data processing.
func IsAnyBlank(strings ...string) bool {
	for _, s := range strings {
		if IsBlank(s) {
			return true
		}
	}
	return false
}

// IsNoneBlank checks if none of the provided strings are blank (empty or containing only whitespace).
//
// This function iterates through a variadic list of strings and returns `false` if any of them
// is blank. A string is considered blank if it is either empty or consists solely of whitespace
// characters. If all strings are non-blank, it returns `true`.
//
// Parameters:
//   - `strings`: A variadic parameter that accepts one or more strings to check for blankness.
//
// Returns:
//   - `true` if none of the provided strings are blank;
//     `false` if at least one of the strings is blank.
//
// Example:
//
// result1 := IsNoneBlank("Hello", "World", " ")
// // result1 will be false because the third string is blank (contains only a space).
//
// result2 := IsNoneBlank("Hello", "World")
// // result2 will be true because both strings are non-blank.
//
// Notes:
//   - This function is useful for validating input or ensuring that all required fields
//     contain meaningful data in forms or data processing.
func IsNoneBlank(strings ...string) bool {
	for _, s := range strings {
		if IsBlank(s) {
			return false
		}
	}
	return true
}

// IsNumeric checks if the provided string contains only numeric digits (0-9).
//
// This function iterates through each character of the input string and verifies if each
// character is a digit using the `unicode.IsDigit` function. If it encounters any character
// that is not a digit, it returns `false`. If all characters are digits, it returns `true`.
//
// Parameters:
//   - `str`: The input string to be checked for numeric characters.
//
// Returns:
//   - `true` if the string contains only numeric digits;
//     `false` if the string contains any non-numeric characters.
//
// Example:
//
// result1 := IsNumeric("12345")
// // result1 will be true because the string contains only digits.
//
// result2 := IsNumeric("123A45")
// // result2 will be false because the string contains a non-digit character ('A').
//
// Notes:
//   - This function is useful for validating numeric input, such as in forms or parsing
//     data that should be strictly numeric.
func IsNumeric(str string) bool {
	for _, c := range str {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// IsNumericSpace checks if the provided string contains only numeric digits and whitespace characters.
//
// This function iterates through each character of the input string and verifies if each character
// is either a digit (0-9) or a whitespace character (spaces, tabs, etc.) using the `unicode.IsDigit`
// and `unicode.IsSpace` functions. If it encounters any character that is neither a digit nor a
// whitespace, it returns `false`. If all characters are valid, it returns `true`.
//
// Parameters:
//   - `str`: The input string to be checked for numeric digits and whitespace.
//
// Returns:
//   - `true` if the string contains only digits and whitespace;
//     `false` if the string contains any non-numeric and non-whitespace characters.
//
// Example:
//
// result1 := IsNumericSpace("123 456")
// // result1 will be true because the string contains only digits and a space.
//
// result2 := IsNumericSpace("123A456")
// // result2 will be false because the string contains a non-numeric character ('A').
//
// Notes:
//   - This function is useful for validating input that should be a number but may also include
//     spaces, such as in user forms or data processing where formatting is flexible.
func IsNumericSpace(str string) bool {
	for _, c := range str {
		if !unicode.IsDigit(c) && !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// IsWhitespace checks if the provided string contains only whitespace characters.
//
// This function iterates through each character of the input string and checks if each character
// is a whitespace character (spaces, tabs, newlines, etc.) using the `unicode.IsSpace` function.
// If it encounters any character that is not a whitespace, it returns `false`. If all characters
// are whitespace, it returns `true`.
//
// Parameters:
//   - `str`: The input string to be checked for whitespace.
//
// Returns:
//   - `true` if the string contains only whitespace characters;
//     `false` if the string contains any non-whitespace characters.
//
// Example:
//
// result1 := IsWhitespace("    ")
// // result1 will be true because the string contains only spaces.
//
// result2 := IsWhitespace("Hello")
// // result2 will be false because the string contains non-whitespace characters.
//
// Notes:
//   - This function is useful for determining if a string is blank in terms of visible content,
//     which can be important in user input validation or string processing tasks.
func IsWhitespace(str string) bool {
	for _, c := range str {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
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

// Abbreviate abbreviates a string by shortening it to a specified `maxWidth` and adding ellipses if necessary.
//
// This function shortens a given string `str` to the length specified by `maxWidth`.
// If the string is already shorter than `maxWidth`, it is returned as is. If the
// string needs to be shortened, an ellipsis ("...") is added at the end of the abbreviated
// string. The abbreviation starts from the beginning of the string.
//
// Parameters:
//   - `str`: The input string to abbreviate.
//   - `maxWidth`: The maximum allowed length for the abbreviated string, including the ellipsis.
//
// Returns:
//   - A string shortened to `maxWidth` characters with an ellipsis if abbreviation is required.
//     If the input string is shorter than or equal to `maxWidth`, the original string is returned.
//
// Example:
//
// input := "This is a long string."
// output := Abbreviate(input, 10)
// // output will be "This is..."
func Abbreviate(str string, maxWidth int) string {
	if IsEmpty(str) {
		return str
	}
	return AbbreviateWithOffset(str, 0, maxWidth)
}

// AbbreviateWithOffset abbreviates a string by shortening it to a specified `maxWidth` starting at a given offset.
//
// This function shortens the input string `str` to the length specified by `maxWidth`, starting from the specified
// `offset`. An ellipsis ("...") is added at the beginning or end of the abbreviated string depending on the offset.
// If the string is already shorter than `maxWidth`, it is returned as is. If the offset is past the length of the string,
// the function adjusts the offset to ensure the abbreviation can be made.
//
// Parameters:
//   - `str`: The input string to abbreviate.
//   - `offset`: The index from which the abbreviation should start.
//   - `maxWidth`: The maximum allowed length for the abbreviated string, including the ellipsis.
//
// Returns:
//   - A string shortened to `maxWidth` characters with an ellipsis. If the input string is shorter
//     than or equal to `maxWidth`, or the offset is invalid, the original string is returned.
//
// Example:
//
// input := "This is a long string."
// output := AbbreviateWithOffset(input, 5, 10)
// // output will be "...s is a..."
//
// Notes:
//   - The function ensures that the abbreviation makes sense even if the offset or maxWidth
//     values are adjusted to fit the input string's length.
func AbbreviateWithOffset(str string, offset int, maxWidth int) string {
	size := len(str)
	if IsEmpty(str) || maxWidth < 4 || size <= maxWidth {
		return str
	}
	if offset > size {
		offset = size
	}
	if size-offset < maxWidth-3 {
		offset = size - (maxWidth - 3)
	}
	abbrevMarker := "..."
	if offset <= 4 {
		return str[0:maxWidth-3] + abbrevMarker
	}
	if maxWidth < 7 {
		return str
	}
	if offset+maxWidth-3 < size {
		return abbrevMarker + Abbreviate(str[offset:], maxWidth-3)
	}
	return abbrevMarker + str[size-(maxWidth-3):]
}

// EndsWith checks if the given string `str` ends with the specified suffix `suffix`.
//
// This function performs a case-sensitive comparison to determine if `str` ends with `suffix`.
// If either the string or the suffix is empty, it will return true only if both are empty.
// It uses the internal helper function `internalEndsWith` to handle the logic.
//
// Parameters:
//   - `str`: The input string to check.
//   - `suffix`: The suffix to check for at the end of the string.
//
// Returns:
//   - `true` if `str` ends with the specified `suffix`, `false` otherwise.
//
// Example:
//
// input := "hello world"
// suffix := "world"
// result := EndsWith(input, suffix)
// // result will be true
func EndsWith(str string, suffix string) bool {
	return internalEndsWith(str, suffix, false)
}

// EndsWithIgnoreCase performs a case-insensitive check to determine if a string ends with a specified suffix.
//
// This function converts both `str` and `suffix` to lowercase before checking if `str` ends with `suffix`.
// Like `EndsWith`, it will return true if both the string and the suffix are empty.
// It uses the internal helper function `internalEndsWith` to handle the case-insensitive comparison.
//
// Parameters:
//   - `str`: The input string to check.
//   - `suffix`: The suffix to check for at the end of the string.
//
// Returns:
//   - `true` if `str` ends with the specified `suffix` (case-insensitively), `false` otherwise.
//
// Example:
//
// input := "hello world"
// suffix := "WORLD"
// result := EndsWithIgnoreCase(input, suffix)
// // result will be true
func EndsWithIgnoreCase(str string, suffix string) bool {
	return internalEndsWith(str, suffix, true)
}

// AppendIfMissing appends a specified suffix to a string if it is not already present.
//
// This function checks if the input string `str` ends with a specified `suffix` or any of the additional
// suffixes provided in the variadic `suffixes` parameter. If none of the specified suffixes are present,
// it appends the main suffix to `str`. The comparison is case-sensitive.
//
// Parameters:
//   - `str`: The input string to check and potentially modify.
//   - `suffix`: The main suffix to append if `str` does not already end with it.
//   - `suffixes`: An optional variadic parameter specifying additional suffixes to check against.
//
// Returns:
//   - The original string if it already ends with the specified suffix (or one of the provided suffixes).
//     Otherwise, it returns the string with the `suffix` appended.
//
// Example:
//
// input := "example"
// suffix := "txt"
// result := AppendIfMissing(input, suffix)
// // result will be "exampletxt" if input does not already end with "txt".
//
// Notes:
//   - This function is case-sensitive.
func AppendIfMissing(str string, suffix string, suffixes ...string) string {
	return internalAppendIfMissing(str, suffix, false, suffixes...)
}

// AppendIfMissingIgnoreCase appends a specified suffix to a string if it is not already present,
// using case-insensitive comparison.
//
// This function checks if the input string `str` ends with a specified `suffix` or any of the additional
// suffixes provided in the variadic `suffixes` parameter. If none of the specified suffixes are present,
// it appends the main suffix to `str`. The comparison is case-insensitive.
//
// Parameters:
//   - `str`: The input string to check and potentially modify.
//   - `suffix`: The main suffix to append if `str` does not already end with it.
//   - `suffixes`: An optional variadic parameter specifying additional suffixes to check against.
//
// Returns:
//   - The original string if it already ends with the specified suffix (or one of the provided suffixes).
//     Otherwise, it returns the string with the `suffix` appended.
//
// Example:
//
// input := "example"
// suffix := "Txt"
// result := AppendIfMissingIgnoreCase(input, suffix)
// // result will be "exampleTxt" if input does not already end with "Txt" or any suffix provided.
//
// Notes:
//   - This function is case-insensitive, so "Txt" and "txt" will be considered equivalent.
func AppendIfMissingIgnoreCase(str string, suffix string, suffixes ...string) string {
	return internalAppendIfMissing(str, suffix, true, suffixes...)
}

// Capitalize capitalizes the first letter of the given string.
//
// This function converts the first character of the input string `str` to its uppercase form,
// while leaving the rest of the string unchanged. If the input string is empty, it returns the string as is.
//
// Parameters:
//   - `str`: The input string that needs to have its first letter capitalized.
//
// Returns:
//   - A new string with the first character in uppercase and the rest of the string unchanged.
//     If the input string is empty, the original string is returned.
//
// Example:
//
// input := "hello"
// output := Capitalize(input)
// // output will be "Hello"
//
// Notes:
//   - The function only modifies the first character. If the first character is already uppercase,
//     or if it's not a letter, the string is returned unchanged.
func Capitalize(str string) string {
	if IsEmpty(str) {
		return str
	}
	for i, c := range str {
		return string(unicode.ToUpper(c)) + str[i+1:]
	}
	return ""
}

// Chomp removes the trailing newline or carriage return characters from the given string.
//
// This function trims newline (`\n`), carriage return (`\r`), or the combination of both (`\r\n`) from the
// end of the input string `str`. If the string consists solely of one newline or carriage return character,
// it returns an empty string. If the string does not end with any of these characters, the original string
// is returned unchanged.
//
// Parameters:
//   - `str`: The input string from which trailing newline or carriage return characters will be removed.
//
// Returns:
//   - A new string with the trailing newline or carriage return characters removed, or the original string
//     if none are found.
//
// Example:
//
// input := "hello\n"
// output := Chomp(input)
// // output will be "hello"
//
// input := "world\r\n"
// output := Chomp(input)
// // output will be "world"
//
// input := "test"
// output := Chomp(input)
// // output will be "test" (no change)
//
// Notes:
//   - This function removes at most one newline or carriage return sequence from the end of the string.
//   - It handles the common line-ending conventions: `\n` (Unix), `\r\n` (Windows), and `\r` (legacy Mac).
func Chomp(str string) string {
	if IsEmpty(str) {
		return str
	}
	if len(str) == 1 && (str[0:1] == "\n" || str[0:1] == "\r") {
		return ""
	}
	if EndsWith(str, "\r\n") {
		return str[:len(str)-2]
	} else if EndsWith(str, "\n") || EndsWith(str, "\r") {
		return str[:len(str)-1]
	} else {
		return str
	}
}

// Chop removes the last character or the last newline sequence from the given string.
//
// This function works similarly to `Chomp` in that it removes newline sequences (`\n`, `\r`, `\r\n`),
// but if the input string does not end with a newline, it removes the last character instead.
// If the string is empty, it returns the string unchanged.
//
// Parameters:
//   - `str`: The input string from which the last character or newline sequence will be removed.
//
// Returns:
//   - A new string with either the trailing newline sequence or the last character removed, or
//     the original string if it is empty.
//
// Example:
//
// input := "hello\n"
// output := Chop(input)
// // output will be "hello" (removes the newline)
//
// input := "world"
// output := Chop(input)
// // output will be "worl" (removes the last character)
//
// Notes:
//   - `Chop` first attempts to remove a newline sequence using `Chomp`. If a newline is found and removed,
//     it returns the result. Otherwise, it removes the last character of the string.
//   - If the string contains only one character, `Chop` will return an empty string.
func Chop(str string) string {
	if IsEmpty(str) {
		return str
	}
	sc := Chomp(str)
	if len(str) > len(sc) {
		return sc
	}
	return str[0 : len(str)-1]
}

// Contains checks if the given string contains a specified substring.
//
// This function checks if the input string `str` contains the substring `search`.
// It performs a case-sensitive comparison to determine if the substring is present.
//
// Parameters:
//   - `str`: The input string in which to search for the substring.
//   - `search`: The substring to search for within `str`.
//
// Returns:
//   - `true` if the `search` substring is found within `str`, `false` otherwise.
//
// Example:
//
// input := "hello world"
// search := "world"
// result := Contains(input, search)
// // result will be true as "world" is present in "hello world".
func Contains(str string, search string) bool {
	return strings.Contains(str, search)
}

// ContainsAny checks if the given string contains any of the specified substrings.
//
// This function checks if the input string `str` contains any of the substrings provided
// in the variadic `search` parameter. It returns `true` as soon as one of the substrings is found.
// The comparison is case-sensitive.
//
// Parameters:
//   - `str`: The input string in which to search for substrings.
//   - `search`: A variadic parameter that accepts one or more substrings to search for.
//
// Returns:
//   - `true` if any of the specified substrings are found within `str`, `false` otherwise.
//
// Example:
//
// input := "hello world"
// search := []string{"foo", "world"}
// result := ContainsAny(input, search...)
// // result will be true because "world" is found in "hello world".
//
// Notes:
//   - The function iterates over the `search` substrings and checks each one individually.
//   - The search is case-sensitive.
func ContainsAny(str string, search ...string) bool {
	for _, s := range search {
		if Contains(str, (string)(s)) {
			return true
		}
	}
	return false
}

// ContainsIgnoreCase checks if the string contains the specified substring,
// ignoring case differences.
//
// This function converts both the input string `str` and the `search` substring
// to lowercase before performing the check, ensuring that the comparison is
// case-insensitive.
//
// Parameters:
//   - `str`: The input string in which to search for the substring.
//   - `search`: The substring to search for within `str`.
//
// Returns:
//   - `true` if the `search` substring is found within `str` (case-insensitive),
//     `false` otherwise.
//
// Example:
//
// input := "Hello World"
// search := "world"
// result := ContainsIgnoreCase(input, search)
// // result will be true as "world" is found in "Hello World" ignoring case.
func ContainsIgnoreCase(str string, search string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(search))
}

// ContainsNone checks if the string contains none of the specified substrings.
//
// This function checks the input string `str` to ensure that it does not contain
// any of the substrings provided in the variadic `search` parameter. It returns
// `true` if none of the substrings are present and `false` if at least one is found.
//
// Parameters:
//   - `str`: The input string to check against the substrings.
//   - `search`: A variadic parameter that accepts one or more substrings to check for.
//
// Returns:
//   - `true` if `str` does not contain any of the specified substrings,
//     `false` otherwise.
//
// Example:
//
// input := "hello world"
// search := []string{"foo", "bar"}
// result := ContainsNone(input, search...)
// // result will be true because neither "foo" nor "bar" are present in "hello world".
func ContainsNone(str string, search ...string) bool {
	return !ContainsAny(str, search...)
}

// ContainsAnyCharacter checks if the string contains any of the characters
// specified in the `search` string.
//
// This function iterates through each character in the `search` string and checks
// if any of those characters are present in the input string `str`.
//
// Parameters:
//   - `str`: The input string in which to search for characters.
//   - `search`: A string containing characters to search for within `str`.
//
// Returns:
//   - `true` if any character from `search` is found in `str`, `false` otherwise.
//
// Example:
//
// input := "hello"
// search := "aeiou"
// result := ContainsAnyCharacter(input, search)
// // result will be true because "hello" contains the character 'e'.
func ContainsAnyCharacter(str string, search string) bool {
	for _, c := range search {
		if Contains(str, (string)(c)) {
			return true
		}
	}
	return false
}

// ContainsNoneCharacter checks if the string contains none of the characters
// specified in the `search` string.
//
// This function calls `ContainsAnyCharacter` to determine if any character
// from `search` is present in the input string `str`. It returns `true` if
// none of the characters are found and `false` if at least one is found.
//
// Parameters:
//   - `str`: The input string to check against the characters.
//   - `search`: A string containing characters to check for within `str`.
//
// Returns:
//   - `true` if `str` does not contain any characters from `search`,
//     `false` otherwise.
//
// Example:
//
// input := "hello"
// search := "xyz"
// result := ContainsNoneCharacter(input, search)
// // result will be true because "hello" contains none of the characters 'x', 'y', or 'z'.
func ContainsNoneCharacter(str string, search string) bool {
	return !ContainsAnyCharacter(str, search)
}

// ContainsOnly checks if a string contains only the specified characters.
//
// This function evaluates the input string `str` and verifies that every character
// in the string is present in the provided variadic parameter `search`. If the input
// string is empty, the function returns `false`. If all characters in the string
// are found in the `search` slice, it returns `true`; otherwise, it returns `false`.
//
// Parameters:
//   - `str`: The input string to be checked.
//   - `search`: A variadic parameter that accepts one or more characters to check
//     against the characters in `str`.
//
// Returns:
//   - `true` if all characters in `str` are present in `search`; `false` otherwise.
//
// Example:
//
// inputStr := "abc"
// allowedChars := []string{"a", "b", "c", "d"}
// result := ContainsOnly(inputStr, allowedChars...)
// // result will be true because "abc" contains only the allowed characters.
//
// Notes:
//   - The function performs a character-by-character check and is case-sensitive.
func ContainsOnly(str string, search ...string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !stringInSlice((string)(c), search) {
			return false
		}
	}
	return true
}

// IsAllLowerCase checks if the string contains only lowercase characters.
func IsAllLowerCase(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !unicode.IsLower(c) {
			return false
		}
	}
	return true
}

// IsAllUpperCase checks if the string contains only uppercase characters.
func IsAllUpperCase(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !unicode.IsUpper(c) {
			return false
		}
	}
	return true
}

// IsAlpha checks if the string contains only Unicode letters.
func IsAlpha(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

// IsAlphanumeric checks if a string contains only alphanumeric characters.
//
// This function evaluates the input string `str` and determines if it consists solely
// of letters (both uppercase and lowercase) and digits. If the string is empty,
// the function returns `false`. The check is performed character by character,
// returning `false` as soon as a non-alphanumeric character is found; if all
// characters are alphanumeric, it returns `true`.
//
// Parameters:
//   - `str`: The input string to be checked for alphanumeric content.
//
// Returns:
//   - `true` if the string contains only letters and digits; `false` otherwise.
//
// Example:
//
// inputStr := "Hello123"
// result := IsAlphanumeric(inputStr)
// // result will be true because "Hello123" consists only of letters and digits.
//
// Notes:
//   - This function is case-sensitive; both uppercase and lowercase letters are accepted.
//   - Special characters and whitespace will cause the function to return false.
func IsAlphanumeric(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// IsAlphaSpace checks if a string contains only alphabetic characters and spaces.
//
// This function evaluates the input string `str` to determine if it consists solely
// of letters (both uppercase and lowercase) and whitespace characters. If the string
// is empty, the function returns `false`. The check is performed character by character,
// returning `false` as soon as a non-alphabetic and non-space character is found;
// if all characters are either alphabetic or spaces, it returns `true`.
//
// Parameters:
//   - `str`: The input string to be checked for alphabetic characters and spaces.
//
// Returns:
//   - `true` if the string contains only letters and spaces; `false` otherwise.
//
// Example:
//
// inputStr := "Hello World"
// result := IsAlphaSpace(inputStr)
// // result will be true because "Hello World" consists only of letters and spaces.
//
// Notes:
//   - This function is case-sensitive; both uppercase and lowercase letters are accepted.
//   - Any numeric characters, punctuation, or special characters will cause the function
//     to return false.
func IsAlphaSpace(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// IsAlphanumericSpace checks if a string contains only alphanumeric characters and spaces.
//
// This function evaluates the input string `str` to determine if it consists solely
// of letters (both uppercase and lowercase), digits, and whitespace characters. If the
// string is empty, the function returns `false`. The check is performed character by
// character, returning `false` as soon as a non-alphanumeric and non-space character
// is found; if all characters are either alphanumeric or spaces, it returns `true`.
//
// Parameters:
//   - `str`: The input string to be checked for alphanumeric characters and spaces.
//
// Returns:
//   - `true` if the string contains only letters, digits, and spaces; `false` otherwise.
//
// Example:
//
// inputStr := "Hello123 World"
// result := IsAlphanumericSpace(inputStr)
// // result will be true because "Hello123 World" consists only of letters, digits, and spaces.
//
// Notes:
//   - This function is case-sensitive; both uppercase and lowercase letters are accepted.
//   - Any special characters or punctuation will cause the function to return false.
func IsAlphanumericSpace(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) && !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// JoinBool concatenates a slice of boolean values into a single string,
// with each value separated by the specified separator.
//
// This function converts each boolean in the input slice `a` to its string representation
// ("true" or "false") using strconv.FormatBool, and then joins them together using the
// specified separator `sep`.
//
// Parameters:
//   - `a`: A slice of boolean values to be joined.
//   - `sep`: A string used to separate the boolean values in the output.
//
// Returns:
//   - A string that contains the boolean values joined together, separated by `sep`.
//
// Example:
//
// bools := []bool{true, false, true}
// result := JoinBool(bools, ", ")
// // result will be "true, false, true"
func JoinBool(a []bool, sep string) string {
	strs := make([]string, len(a))
	for idx, i := range a {
		strs[idx] = strconv.FormatBool(i)
	}
	return JoinUnary(strs, sep)
}

// JoinFloat64 concatenates a slice of float64 values into a single string,
// using the default format ('G') and bit size (32) for float conversion.
//
// This function converts each float64 in the input slice `a` to its string representation
// using strconv.FormatFloat with default formatting options, and then joins them together
// using the specified separator `sep`.
//
// Parameters:
//   - `a`: A slice of float64 values to be joined.
//   - `sep`: A string used to separate the float64 values in the output.
//
// Returns:
//   - A string that contains the float64 values joined together, separated by `sep`.
//
// Example:
//
// floats := []float64{1.23, 4.56, 7.89}
// result := JoinFloat64(floats, ", ")
// // result will be "1.23, 4.56, 7.89"
func JoinFloat64(a []float64, sep string) string {
	return JoinFloat64WithFormatAndPrecision(a, 'G', 32, sep)
}

// JoinFloat64WithFormatAndPrecision concatenates a slice of float64 values into a single string,
// allowing for custom formatting and precision during the conversion.
//
// This function converts each float64 in the input slice `a` to its string representation
// using strconv.FormatFloat with the specified format `fmt` and precision `precision`,
// and then joins them together using the specified separator `sep`.
//
// Parameters:
//   - `a`: A slice of float64 values to be joined.
//   - `fmt`: A byte that specifies the format for float conversion ('b', 'e', 'E', 'f', 'g', 'G').
//   - `precision`: An integer specifying the precision for float conversion (bit size).
//   - `sep`: A string used to separate the float64 values in the output.
//
// Returns:
//   - A string that contains the float64 values joined together, separated by `sep`.
//
// Example:
//
// floats := []float64{1.2345, 6.7890}
// result := JoinFloat64WithFormatAndPrecision(floats, 'f', 2, ", ")
// // result will be "1.23, 6.79"
func JoinFloat64WithFormatAndPrecision(a []float64, fmt byte, precision int, sep string) string {
	strs := make([]string, len(a))
	for idx, i := range a {
		strs[idx] = strconv.FormatFloat(i, fmt, -1, precision)
	}
	return JoinUnary(strs, sep)
}

// JoinInt concatenates a slice of integers into a single string,
// with each integer separated by the specified separator.
//
// This function converts each integer in the input slice `a` to its string representation
// using strconv.Itoa, and then joins them together using the specified separator `sep`.
//
// Parameters:
//   - `a`: A slice of integers to be joined.
//   - `sep`: A string used to separate the integer values in the output.
//
// Returns:
//   - A string that contains the integer values joined together, separated by `sep`.
//
// Example:
//
// ints := []int{1, 2, 3}
// result := JoinInt(ints, ", ")
// // result will be "1, 2, 3"
func JoinInt(a []int, sep string) string {
	strs := make([]string, len(a))
	for idx, i := range a {
		strs[idx] = strconv.Itoa(i)
	}
	return JoinUnary(strs, sep)
}

// JoinInt64 concatenates a slice of int64 values into a single string,
// with each int64 separated by the specified separator.
//
// This function converts each int64 in the input slice `a` to its string representation
// using strconv.FormatInt, and then joins them together using the specified separator `sep`.
//
// Parameters:
//   - `a`: A slice of int64 values to be joined.
//   - `sep`: A string used to separate the int64 values in the output.
//
// Returns:
//   - A string that contains the int64 values joined together, separated by `sep`.
//
// Example:
//
// int64s := []int64{100, 200, 300}
// result := JoinInt64(int64s, ", ")
// // result will be "100, 200, 300"
func JoinInt64(a []int64, sep string) string {
	strs := make([]string, len(a))
	for idx, i := range a {
		strs[idx] = strconv.FormatInt(i, 10)
	}
	return JoinUnary(strs, sep)
}

// JoinUint64 concatenates a slice of uint64 values into a single string,
// with each uint64 separated by the specified separator.
//
// This function converts each uint64 in the input slice `ints` to its string representation
// using strconv.FormatUint, and then joins them together using the specified separator `sep`.
//
// Parameters:
//   - `ints`: A slice of uint64 values to be joined.
//   - `sep`: A string used to separate the uint64 values in the output.
//
// Returns:
//   - A string that contains the uint64 values joined together, separated by `sep`.
//
// Example:
//
// uint64s := []uint64{1000, 2000, 3000}
// result := JoinUint64(uint64s, ", ")
// // result will be "1000, 2000, 3000"
func JoinUint64(ints []uint64, sep string) string {
	strs := make([]string, len(ints))
	for idx, i := range ints {
		strs[idx] = strconv.FormatUint(i, 10)
	}
	return JoinUnary(strs, sep)
}

// Mid extracts a substring of a specified size from the middle of a given string, starting at a specified position.
//
// This function returns a substring from the input string `str`, starting from the specified `pos`
// and extending for the specified number of `size` characters. If the `size` is negative, it returns an empty string.
// If the starting position is outside the bounds of the string, or if the input string is empty,
// an empty string is returned. If the specified `size` extends beyond the end of the string,
// the substring from the starting position to the end of the string is returned.
//
// Parameters:
//   - `str`: The input string from which to extract the substring.
//   - `pos`: The starting position in the string to begin extraction.
//   - `size`: The number of characters to extract from the string.
//
// Returns:
//   - A substring of the input string starting from `pos` and having length `size`.
//     If the starting position is invalid or the size is negative, returns an empty string.
//
// Example:
//
// input := "Hello, World!"
// result := Mid(input, 7, 5)
// // result will be "World"
func Mid(str string, pos int, size int) string {
	if str == "" || size < 0 || pos > len(str) {
		return ""
	}
	if pos < 0 {
		pos = 0
	}
	if len(str) <= pos+size {
		return str[pos:]
	}
	return str[pos : pos+size]
}

// Overlay replaces a portion of the input string with another string, starting from a specified start index
// and ending at a specified end index.
//
// This function takes the input string `str` and overlays it with the string `overlay`, replacing the
// portion of `str` that starts at index `start` and ends at index `end`. If the start index is out of bounds,
// it is adjusted to the nearest valid position. If the end index is out of bounds, it is set to the length
// of the string. If the start index is greater than the end index, the indices are swapped.
//
// Parameters:
//   - `str`: The original string to be modified.
//   - `overlay`: The string that will replace the specified portion of `str`.
//   - `start`: The starting index from which to begin the overlay.
//   - `end`: The ending index up to which the overlay will occur.
//
// Returns:
//   - A new string that results from overlaying `overlay` onto `str` between the specified indices.
//     If both `start` and `end` are out of bounds, returns the original string without modification.
//
// Example:
//
// original := "Hello, World!"
// newString := Overlay(original, "Gopher", 7, 12)
// // newString will be "Hello, Gopher!"
func Overlay(str string, overlay string, start int, end int) string {
	strLen := len(str)
	// guards
	if start < 0 {
		start = 0
	}
	if start > strLen {
		start = strLen
	}
	if end < 0 {
		end = 0
	}
	if end > strLen {
		end = strLen
	}
	if start > end {
		start, end = end, start
	}
	return str[:start] + overlay + str[end:]
}

// Remove removes all occurrences of a specified substring from the source string.
//
// This function searches for all instances of the substring `remove` within the input string `str`
// and removes them. If the input string is empty, the function returns it unchanged. The function
// utilizes a loop to find and remove each occurrence of the specified substring until none remain.
//
// Parameters:
//   - `str`: The source string from which the substring will be removed.
//   - `remove`: The substring to be removed from the source string.
//
// Returns:
//   - A new string with all occurrences of `remove` removed from `str`.
//     If the `remove` substring is not found, or if `str` is empty, the original string is returned unchanged.
//
// Example:
//
// input := "Hello, world! Goodbye, world!"
// result := Remove(input, "world")
// // result will be "Hello, ! Goodbye, !"
func Remove(str string, remove string) string {
	if IsEmpty(str) {
		return str
	}
	index := strings.Index(str, remove)
	for index > -1 {
		str = str[:index] + str[index+len(remove):]
		index = strings.Index(str, remove)
	}
	return str
}

// RemoveEnd removes a specified substring from the end of a source string,
// if the substring is found there. If the substring is not present at the end,
// the original source string is returned unchanged.
//
// Parameters:
//   - `str`: The source string from which the substring will be removed.
//   - `remove`: The substring to be removed from the end of the source string.
//
// Returns:
//   - A new string with the specified `remove` substring removed from the end
//     of `str`. If `str` is empty or `remove` is empty, the original string is returned.
//     If the substring is not found at the end, the original string is returned unchanged.
//
// Example:
//
// input := "example.txt"
// result := RemoveEnd(input, ".txt")
// // result will be "example" since ".txt" is at the end of the string.
func RemoveEnd(str string, remove string) string {
	if IsEmpty(str) || IsEmpty(remove) {
		return str
	}
	if EndsWith(str, remove) {
		return str[:len(str)-len(remove)]
	}
	return str
}

// RemoveEndIgnoreCase removes a specified substring from the end of a source string
// in a case-insensitive manner. If the substring is found at the end, it is removed;
// otherwise, the original string is returned unchanged.
//
// Parameters:
//   - `str`: The source string from which the substring will be removed.
//   - `remove`: The substring to be removed from the end of the source string.
//
// Returns:
//   - A new string with the specified `remove` substring removed from the end
//     of `str`. If `str` is empty or `remove` is empty, the original string is returned.
//     If the substring is not found at the end (case insensitive), the original string is returned unchanged.
//
// Example:
//
// input := "example.TXT"
// result := RemoveEndIgnoreCase(input, ".txt")
// // result will be "example" since ".txt" is found at the end of the string ignoring case.
func RemoveEndIgnoreCase(str string, remove string) string {
	if IsEmpty(str) || IsEmpty(remove) {
		return str
	}
	if EndsWithIgnoreCase(str, remove) {
		return str[:len(str)-len(remove)]
	}
	return str
}

// RemovePattern removes all substrings from the source string that match the specified
// regular expression pattern. If no matches are found, the original string is returned unchanged.
//
// Parameters:
//   - `str`: The source string from which substrings matching the pattern will be removed.
//   - `pattern`: A regular expression pattern that defines the substrings to be removed.
//
// Returns:
//   - A new string with all occurrences of substrings matching the given `pattern` removed.
//     If no matches are found, the original string is returned unchanged.
//
// Example:
//
// input := "abc123xyz"
// pattern := "[0-9]+" // matches all digits
// result := RemovePattern(input, pattern)
// // result will be "abcxyz" since all digits are removed from the string.
func RemovePattern(str string, pattern string) string {
	return regexp.MustCompile(pattern).ReplaceAllString(str, "")
}

// RemoveStart removes a substring only if it is at the beginning of a source string,
// otherwise returns the source string
func RemoveStart(str string, remove string) string {
	if IsEmpty(str) || IsEmpty(remove) {
		return str
	}
	if StartsWith(str, remove) {
		return str[len(remove)+1:]
	}
	return str
}

// RemoveStartIgnoreCase removes a specified substring from the start of a string,
// ignoring case differences.
//
// This function checks if the input string `str` starts with the specified `remove`
// substring in a case-insensitive manner. If `str` starts with `remove`, it removes
// that substring from the start of `str` and returns the modified string. If `str`
// does not start with `remove`, or if either `str` or `remove` is empty, the function
// returns the original string unchanged.
//
// Parameters:
//   - `str`: The source string from which to remove the specified substring.
//   - `remove`: The substring to be removed from the start of `str`.
//
// Returns:
//   - The modified string with the `remove` substring removed from the start if it
//     was found; otherwise, returns the original string unchanged.
//
// Example:
//
// input := "Hello, World!"
// remove := "hello"
// result := RemoveStartIgnoreCase(input, remove)
// // result will be "Hello, World!" since input does not start with "hello" (case-insensitive).
//
// input2 := "Hello, World!"
// remove2 := "Hello"
// result2 := RemoveStartIgnoreCase(input2, remove2)
// // result2 will be ", World!" since input2 starts with "Hello" (case-insensitive).
//
// Notes:
//   - This function is case-insensitive, so it will remove the substring regardless of
//     the casing in `str` and `remove`.
func RemoveStartIgnoreCase(str string, remove string) string {
	if IsEmpty(str) || IsEmpty(remove) {
		return str
	}
	if StartsWithIgnoreCase(str, remove) {
		return str[len(remove)+1:]
	}
	return str
}

// StartsWith checks if a string starts with a specified prefix.
//
// This function determines whether the input string `str` begins with the specified
// `prefix` in a case-sensitive manner.
//
// Parameters:
//   - `str`: The source string to be checked for the specified prefix.
//   - `prefix`: The prefix string to look for at the start of `str`.
//
// Returns:
//   - `true` if `str` starts with `prefix` (case-sensitive).
//   - `false` if `str` does not start with `prefix`, or if either `str` or `prefix` is empty
//     (in which case it returns `true` only if both are empty).
//
// Example:
//
// str := "Hello, World!"
// prefix := "Hello"
// result := StartsWith(str, prefix)
// // result will be true since str starts with the prefix "Hello".
//
// Notes:
//   - This function is case-sensitive; for case-insensitive checks, use a different
//     function that leverages the `internalStartsWith` with the `ignoreCase` flag set to true.
func StartsWith(str string, prefix string) bool {
	return internalStartsWith(str, prefix, false)
}

// StartsWithIgnoreCase checks if a string starts with a specified prefix,
// ignoring case differences.
//
// This function determines whether the input string `str` begins with the specified
// `prefix` in a case-insensitive manner.
//
// Parameters:
//   - `str`: The source string to be checked for the specified prefix.
//   - `prefix`: The prefix string to look for at the start of `str`.
//
// Returns:
//   - `true` if `str` starts with `prefix`, ignoring case differences.
//   - `false` if `str` does not start with `prefix`, or if either `str` or `prefix` is empty
//     (in which case it returns `true` only if both are empty).
//
// Example:
//
// str := "Hello, World!"
// prefix := "hello"
// result := StartsWithIgnoreCase(str, prefix)
// // result will be true since str starts with the prefix "hello", ignoring case.
//
// Notes:
//   - This function is case-insensitive; if a case-sensitive check is required,
//     use the `StartsWith` function instead.
func StartsWithIgnoreCase(str string, prefix string) bool {
	return internalStartsWith(str, prefix, true)
}

// Repeat returns a new string consisting of the specified string repeated a given number of times.
//
// This function takes an input string `str` and an integer `repeat`, and constructs a new
// string that contains `str` concatenated `repeat` times. If `repeat` is less than or equal to
// zero, an empty string is returned.
//
// Parameters:
//   - `str`: The input string to be repeated.
//   - `repeat`: The number of times to repeat the string.
//
// Returns:
//   - A new string that contains `str` repeated `repeat` times. If `repeat` is 0 or negative,
//     an empty string is returned.
//
// Example:
//
// input := "abc"
// result := Repeat(input, 3)
// // result will be "abcabcabc".
//
// input2 := "xyz"
// result2 := Repeat(input2, 0)
// // result2 will be "" (empty string).
func Repeat(str string, repeat int) string {
	buff := ""
	for repeat > 0 {
		repeat = repeat - 1
		buff += str
	}
	return buff
}

// RepeatWithSeparator returns a new string consisting of the specified string repeated a given number of times,
// with a separator placed between each repetition.
//
// This function takes an input string `str`, a separator string `sep`, and an integer `repeat`,
// and constructs a new string that contains `str` concatenated `repeat` times, separated by `sep`.
// If `repeat` is less than or equal to zero, an empty string is returned.
//
// Parameters:
//   - `str`: The input string to be repeated.
//   - `sep`: The separator string to be placed between repetitions of `str`.
//   - `repeat`: The number of times to repeat the string.
//
// Returns:
//   - A new string that contains `str` repeated `repeat` times, with `sep` separating each occurrence.
//     If `repeat` is 0 or negative, an empty string is returned.
//
// Example:
//
// input := "abc"
// sep := "-"
// result := RepeatWithSeparator(input, sep, 3)
// // result will be "abc-abc-abc".
//
// input2 := "hello"
// sep2 := ", "
// result2 := RepeatWithSeparator(input2, sep2, 0)
// // result2 will be "" (empty string).
func RepeatWithSeparator(str string, sep string, repeat int) string {
	buff := ""
	for repeat > 0 {
		repeat = repeat - 1
		buff += str
		if repeat > 0 {
			buff += sep
		}
	}
	return buff
}

// ReverseDelimited reverses the order of substrings in a string that are separated by a specified delimiter.
//
// This function takes an input string `str` and a delimiter `del`, splits the string into
// substrings using the delimiter, reverses the order of these substrings, and then joins them
// back together using the same delimiter. If the input string is empty or does not contain the
// delimiter, it returns the original string unchanged.
//
// Parameters:
//   - `str`: The input string to be reversed, containing substrings separated by the specified delimiter.
//   - `del`: The delimiter used to split the input string into substrings.
//
// Returns:
//   - A new string with the order of substrings reversed. If the input string is empty or the delimiter
//     is not found, the original string is returned unchanged.
//
// Example:
//
// input := "apple,banana,cherry"
// delimiter := ","
// result := ReverseDelimited(input, delimiter)
// // result will be "cherry,banana,apple".
//
// input2 := "one|two|three"
// delimiter2 := "|"
// result2 := ReverseDelimited(input2, delimiter2)
// // result2 will be "three|two|one".
//
// Notes:
//   - The function modifies the order of the substrings but retains the original delimiter.
func ReverseDelimited(str string, del string) string {
	s := strings.Split(str, del)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return JoinUnary(s, del)
}

// Strip removes whitespace from both the start and end of a string.
//
// This function takes an input string `str` and uses a regular expression to
// remove any whitespace characters (spaces, tabs, newlines, etc.) that appear
// at the beginning and the end of the string. If the input string is empty,
// it returns the string unchanged.
//
// Parameters:
//   - `str`: The input string from which leading and trailing whitespace will be removed.
//
// Returns:
//   - A new string with leading and trailing whitespace removed. If there is no
//     whitespace or if the string is empty, the original string is returned unchanged.
//
// Example:
//
// input := "   Hello, World!   "
// result := Strip(input)
// // result will be "Hello, World!".
func Strip(str string) string {
	return regexp.MustCompile(`^\s+|\s+$`).ReplaceAllString(str, "")
}

// StripEnd removes whitespace from the end of a string.
//
// This function takes an input string `str` and uses a regular expression to
// remove any whitespace characters that appear at the end of the string.
// If the input string is empty, it returns the string unchanged.
//
// Parameters:
//   - `str`: The input string from which trailing whitespace will be removed.
//
// Returns:
//   - A new string with trailing whitespace removed. If there is no trailing
//     whitespace or if the string is empty, the original string is returned unchanged.
//
// Example:
//
// input := "Hello, World!   "
// result := StripEnd(input)
// // result will be "Hello, World!".
func StripEnd(str string) string {
	return regexp.MustCompile(`\s+$`).ReplaceAllString(str, "")
}

// StripStart removes whitespace from the start of a string.
//
// This function takes an input string `str` and uses a regular expression to
// remove any whitespace characters that appear at the beginning of the string.
// If the input string is empty, it returns the string unchanged.
//
// Parameters:
//   - `str`: The input string from which leading whitespace will be removed.
//
// Returns:
//   - A new string with leading whitespace removed. If there is no leading
//     whitespace or if the string is empty, the original string is returned unchanged.
//
// Example:
//
// input := "   Hello, World!"
// result := StripStart(input)
// // result will be "Hello, World!".
func StripStart(str string) string {
	return regexp.MustCompile(`^\s+`).ReplaceAllString(str, "")
}

// SwapCase returns a new string with all uppercase letters converted to lowercase
// and all lowercase letters converted to uppercase.
//
// This function iterates through each character in the input string `str` and checks
// whether the character is uppercase or lowercase. If the character is lowercase,
// it converts it to uppercase; if it is uppercase, it converts it to lowercase.
// All non-alphabetic characters remain unchanged.
//
// Parameters:
//   - `str`: The input string whose letter case will be swapped.
//
// Returns:
//   - A new string with the cases of the letters swapped. If the input string is
//     empty or contains no alphabetic characters, it returns the original string
//     unchanged.
//
// Example:
//
// input := "Hello, World!"
// result := SwapCase(input)
// // result will be "hELLO, wORLD!".
//
// Notes:
//   - This function utilizes the `unicode` package to check the case of each character
//     and assumes that the functions `UpperCase` and `LowerCase` are defined to
//     convert characters to their respective cases.
func SwapCase(str string) string {
	buff := ""
	for _, c := range str {
		cs := (string)(c)
		if unicode.IsLower(c) {
			buff += strings.ToUpper(cs)
		} else if unicode.IsUpper(c) {
			buff += strings.ToLower(cs)
		} else {
			buff += cs
		}
	}
	return buff
}

// UnCapitalize changes the first letter of a string to lowercase.
//
// This function checks if the input string `str` is empty. If it is, the function
// returns the original string unchanged. If the first character of the string is
// uppercase, it converts that character to lowercase while leaving the rest of the
// string intact.
//
// Parameters:
//   - `str`: The input string to un-capitalize.
//
// Returns:
//   - A new string with the first character converted to lowercase if it was
//     originally uppercase. If the string is empty or if the first character is
//     already lowercase, the original string is returned unchanged.
//
// Example:
//
// input := "Hello"
// result := UnCapitalize(input)
// // result will be "hello".
//
// Notes:
//   - This function assumes that the input string contains at least one character
//     if it is not empty.
func UnCapitalize(str string) string {
	if IsEmpty(str) {
		return str
	}
	r := []rune(str)
	if unicode.IsUpper(r[0]) {
		return string(unicode.ToLower(r[0])) + string(r[1:])
	}
	return str
}

// Wrap wraps a given string with another string.
//
// This function checks if the input string `str` is empty. If it is, the function
// returns the original string unchanged. If not, it concatenates the `wrapWith`
// string to both the beginning and the end of `str`, effectively wrapping it.
//
// Parameters:
//   - `str`: The input string to be wrapped.
//   - `wrapWith`: The string to wrap around `str`.

// Returns:
//   - A new string that consists of `wrapWith` followed by `str` and then
//     another `wrapWith`. If the input string is empty, the original string is
//     returned unchanged.
//
// Example:
//
// input := "Hello"
// wrapWith := "***"
// result := Wrap(input, wrapWith)
// // result will be "***Hello***".
//
// Notes:
//   - This function does not modify the original string; it creates and returns
//     a new string instead.
func Wrap(str string, wrapWith string) string {
	if IsEmpty(str) {
		return str
	}
	return wrapWith + str + wrapWith
}

// PrependIfMissing prepends a specified prefix to the start of a string if the string does not already start with that prefix
// or any additional prefixes provided.
//
// This function checks if the input string `str` starts with the specified `prefix`. If it does not, it will check
// against any prefixes provided in the variadic `prefixes` parameter. If `str` does not start with any of the
// specified prefixes, the function prepends the `prefix` to `str` and returns the modified string. If `prefix`
// is empty or if `str` already starts with `prefix` or any of the provided prefixes, the original string is returned.
//
// Parameters:
//   - `str`: The input string to check and potentially modify.
//   - `prefix`: The main prefix to prepend if `str` does not already start with it.
//   - `prefixes`: An optional variadic parameter specifying additional prefixes to check against.
//
// Returns:
//   - The original string if it already starts with the specified prefix (or one of the provided prefixes).
//     Otherwise, it returns the string with the `prefix` prepended.
//
// Example:
//
// input := "example"
// prefix := "test_"
// result := PrependIfMissing(input, prefix)
// // result will be "test_example" if input does not already start with "test_".
//
// Notes:
//   - This function is case-sensitive.
func PrependIfMissing(str string, prefix string, prefixes ...string) string {
	return prependIfMissing(str, prefix, false, prefixes...)
}

// PrependIfMissingIgnoreCase prepends a specified prefix to the start of a string if the string does not already start
// (case-insensitively) with that prefix or any additional prefixes provided.
//
// This function performs the same checks as `PrependIfMissing`, but it ignores case when checking for the presence
// of `prefix` and the additional prefixes. If `str` does not start with `prefix` or any of the provided prefixes,
// the function prepends `prefix` to `str` and returns the modified string. If `prefix` is empty or if `str`
// already starts with `prefix` or any of the provided prefixes (ignoring case), the original string is returned.
//
// Parameters:
//   - `str`: The input string to check and potentially modify.
//   - `prefix`: The main prefix to prepend if `str` does not already start with it.
//   - `prefixes`: An optional variadic parameter specifying additional prefixes to check against.
//
// Returns:
//   - The original string if it already starts with the specified prefix (or one of the provided prefixes,
//     ignoring case). Otherwise, it returns the string with the `prefix` prepended.
//
// Example:
//
// input := "example"
// prefix := "Test_"
// result := PrependIfMissingIgnoreCase(input, prefix)
// // result will be "Test_example" if input does not already start with "Test_" or any prefix provided.
//
// Notes:
//   - This function is case-insensitive, so "Test_" and "test_" will be considered equivalent.
func PrependIfMissingIgnoreCase(str string, prefix string, prefixes ...string) string {
	return prependIfMissing(str, prefix, true, prefixes...)
}

// prependIfMissing is an internal function that performs the logic of prepending a prefix to a string
// if the string does not already start with the specified prefix or any of the additional prefixes,
// with an option for case-insensitive comparison.
//
// Parameters:
//   - `str`: The input string to check and potentially modify.
//   - `prefix`: The main prefix to prepend if `str` does not already start with it.
//   - `ignoreCase`: A boolean flag indicating whether to ignore case when checking for prefixes.
//   - `prefixes`: An optional variadic parameter specifying additional prefixes to check against.
//
// Returns:
//   - The original string if it already starts with the specified prefix (or one of the provided prefixes).
//     Otherwise, it returns the string with the `prefix` prepended.
func prependIfMissing(str string, prefix string, ignoreCase bool, prefixes ...string) string {
	if IsEmpty(prefix) || internalStartsWith(str, prefix, ignoreCase) {
		return str
	}
	for _, pref := range prefixes {
		if pref == "" || internalStartsWith(str, pref, ignoreCase) {
			return str
		}
	}
	return prefix + str
}

// internalStartsWith checks if a given string starts with a specified prefix.
//
// This function determines whether the input string `str` begins with the specified
// `prefix`. It can perform the check in a case-sensitive or case-insensitive manner
// based on the `ignoreCase` parameter.
//
// Parameters:
//   - `str`: The source string to be checked for the specified prefix.
//   - `prefix`: The prefix string to look for at the start of `str`.
//   - `ignoreCase`: A boolean indicating whether the comparison should be case-sensitive
//     (false) or case-insensitive (true).
//
// Returns:
//   - `true` if `str` starts with `prefix` (case-sensitive or case-insensitive depending
//     on the `ignoreCase` flag).
//   - `false` if `str` does not start with `prefix`, or if either `str` or `prefix` is empty
//     (in which case it returns `true` only if both are empty).
//
// Example:
//
// str := "Hello, World!"
// prefix := "Hello"
// result := internalStartsWith(str, prefix, false)
// // result will be true since str starts with the prefix "Hello".
//
// Notes:
//   - If both `str` and `prefix` are empty, the function returns true.
//   - If `prefix` is longer than `str`, the function will return false.
func internalStartsWith(str string, prefix string, ignoreCase bool) bool {
	if str == "" || prefix == "" {
		return (str == "" && prefix == "")
	}
	if utf8.RuneCountInString(prefix) > utf8.RuneCountInString(str) {
		return false
	}
	if ignoreCase {
		return strings.HasPrefix(strings.ToLower(str), strings.ToLower(prefix))
	}
	return strings.HasPrefix(str, prefix)
}

// internalEndsWith is a helper function that checks if the given string `str` ends with the specified `suffix`.
//
// This function provides both case-sensitive and case-insensitive checks based on the `ignoreCase` flag.
// It first checks if either the string or suffix is empty, returning true if both are empty. Then, it ensures
// that the `suffix` is not longer than `str`. If `ignoreCase` is true, it performs a case-insensitive comparison;
// otherwise, it uses a case-sensitive comparison.
//
// Parameters:
//   - `str`: The input string to check.
//   - `suffix`: The suffix to check for at the end of the string.
//   - `ignoreCase`: A boolean flag indicating whether the comparison should be case-insensitive.
//
// Returns:
//   - `true` if `str` ends with `suffix` based on the case-sensitivity specified by `ignoreCase`, `false` otherwise.
func internalEndsWith(str string, suffix string, ignoreCase bool) bool {
	if str == "" || suffix == "" {
		return (str == "" && suffix == "")
	}
	if utf8.RuneCountInString(suffix) > utf8.RuneCountInString(str) {
		return false
	}
	if ignoreCase {
		return strings.HasSuffix(strings.ToLower(str), strings.ToLower(suffix))
	}
	return strings.HasSuffix(str, suffix)
}

// internalAppendIfMissing appends a specified suffix to a string if it is not already present.
//
// This function checks if the input string `str` ends with a specified `suffix` or any of the additional
// suffixes provided in the variadic `suffixes` parameter. If none of the specified suffixes are present,
// it appends the first suffix to `str`. The comparison can be case-sensitive or case-insensitive based on
// the `ignoreCase` flag.
//
// Parameters:
//   - `str`: The input string to check and potentially modify.
//   - `suffix`: The main suffix to append if `str` does not already end with it.
//   - `ignoreCase`: A boolean flag that determines whether the comparison should be case-insensitive.
//   - `suffixes`: An optional variadic parameter specifying additional suffixes to check against.
//
// Returns:
//   - The original string if it already ends with the specified suffix (or one of the provided suffixes).
//     Otherwise, it returns the string with the `suffix` appended.
//
// Example:
//
// input := "example"
// suffix := "txt"
// result := internalAppendIfMissing(input, suffix, false)
// // result will be "exampletxt" if input does not already end with "txt".
//
// Notes:
//   - This function is case-insensitive if `ignoreCase` is set to true.
//   - It checks both the `suffix` and any additional suffixes in the order they are provided.
func internalAppendIfMissing(str string, suffix string, ignoreCase bool, suffixes ...string) string {
	if str == "" || IsEmpty(str) {
		return str
	}
	if ignoreCase {
		if EndsWithIgnoreCase(str, suffix) {
			return str
		}

		for _, suffix := range suffixes {
			if EndsWithIgnoreCase(str, (string)(suffix)) {
				return str
			}
		}
	} else {
		if EndsWith(str, suffix) {
			return str
		}

		for _, suffix := range suffixes {
			if EndsWith(str, (string)(suffix)) {
				return str
			}
		}
	}
	return str + suffix
}

// stringInSlice checks if a given string is present in a slice of strings.
//
// This function iterates through the provided slice `list` and compares each
// element with the specified string `a`. If a match is found, it returns
// `true`. If the iteration completes without finding a match, it returns `false`.
//
// Parameters:
//   - `a`: The string to search for within the slice.
//   - `list`: A slice of strings to search through.
//
// Returns:
//   - `true` if the string `a` is found in the `list`, `false` otherwise.
//
// Example:
//
// inputStr := "apple"
// strList := []string{"banana", "orange", "apple", "grape"}
// result := stringInSlice(inputStr, strList)
// // result will be true because "apple" is present in the strList.
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
