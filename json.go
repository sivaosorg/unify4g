package unify4g

import (
	jsonI "github.com/json-iterator/go"
)

var _json = jsonI.ConfigCompatibleWithStandardLibrary

// Marshal converts a Go value into its JSON byte representation.
//
// This function marshals the input value `v` using the json-iterator library.
// The resulting JSON data is returned as a byte slice. If there is an error
// during marshalling, it returns the error.
//
// Parameters:
// - `v`: The Go value to be marshalled into JSON.
//
// Returns:
// - A byte slice containing the JSON representation of the input value.
// - An error if the marshalling fails.
//
// Example:
// jsonData, err := Marshal(myStruct)
func Marshal(v interface{}) ([]byte, error) {
	return _json.Marshal(v)
}

// MarshalIndent converts a Go value to its JSON string representation with indentation.
//
// This function marshals the input value `v` into a formatted JSON string,
// allowing for easy readability by including a specified prefix and indentation.
// It returns the resulting JSON byte slice or an error if marshalling fails.
//
// Parameters:
// - `v`: The Go value to be marshalled into JSON.
// - `prefix`: A string that will be prefixed to each line of the output JSON.
// - `indent`: A string used for indentation (typically a series of spaces or a tab).
//
// Returns:
// - A byte slice containing the formatted JSON representation of the input value.
// - An error if the marshalling fails.
//
// Example:
// jsonIndented, err := MarshalIndent(myStruct, "", "    ")
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return _json.MarshalIndent(v, prefix, indent)
}

// MarshalToString converts a Go value to its JSON string representation.
//
// This function utilizes the json-iterator library to marshal the input value `v`
// into a JSON string. If the marshalling is successful, it returns the resulting
// JSON string. If an error occurs during the process, it returns an error.
//
// Parameters:
// - `v`: The Go value to be marshalled into JSON.
//
// Returns:
// - A string containing the JSON representation of the input value.
// - An error if the marshalling fails.
//
// Example:
// jsonString, err := MarshalToString(myStruct)
func MarshalToString(v interface{}) (string, error) {
	return _json.MarshalToString(v)
}

// Unmarshal parses JSON-encoded data and stores the result in the value pointed to by `v`.
//
// This function uses the json-iterator library to unmarshal JSON data
// (given as a byte slice) into the specified Go value `v`. If the unmarshalling
// is successful, it populates the value `v`. If an error occurs, it returns the error.
//
// Parameters:
// - `data`: A byte slice containing JSON data to be unmarshalled.
// - `v`: A pointer to the Go value where the unmarshalled data will be stored.
//
// Returns:
// - An error if the unmarshalling fails.
//
// Example:
// err := Unmarshal(jsonData, &myStruct)
func Unmarshal(data []byte, v interface{}) error {
	return _json.Unmarshal(data, v)
}

// UnmarshalFromString parses JSON-encoded string and stores the result in the value pointed to by `v`.
//
// This function utilizes the json-iterator library to unmarshal JSON data
// from a string into the specified Go value `v`. If the unmarshalling is
// successful, it populates the value `v`. If an error occurs, it returns the error.
//
// Parameters:
// - `str`: A string containing JSON data to be unmarshalled.
// - `v`: A pointer to the Go value where the unmarshalled data will be stored.
//
// Returns:
// - An error if the unmarshalling fails.
//
// Example:
// err := UnmarshalFromString(jsonString, &myStruct)
func UnmarshalFromString(str string, v interface{}) error {
	return _json.UnmarshalFromString(str, v)
}

// Json converts a Go value to its JSON string representation or returns the value directly if it is already a string.
//
// This function checks if the input data is a string; if so, it returns it directly.
// Otherwise, it marshals the input value `data` into a JSON string using the
// MarshalToString function. If an error occurs during marshalling, it returns an empty string.
//
// Parameters:
// - `data`: The Go value to be converted to JSON, or a string to be returned directly.
//
// Returns:
// - A string containing the JSON representation of the input value, or an empty string if an error occurs.
//
// Example:
// jsonStr := Json(myStruct)
func Json(data interface{}) string {
	s, ok := data.(string)
	if ok {
		return s
	}
	result, err := MarshalToString(data)
	if err != nil {
		return ""
	}
	return string(result)
}

// JsonPretty converts a Go value to its pretty-printed JSON string representation or returns the value directly if it is already a string.
//
// This function checks if the input data is a string; if so, it returns it directly.
// Otherwise, it marshals the input value `data` into a formatted JSON string using
// the MarshalIndent function. If an error occurs during marshalling, it returns an empty string.
//
// Parameters:
// - `data`: The Go value to be converted to pretty-printed JSON, or a string to be returned directly.
//
// Returns:
// - A string containing the pretty-printed JSON representation of the input value, or an empty string if an error occurs.
//
// Example:
// jsonPrettyStr := JsonPretty(myStruct)
func JsonPretty(data interface{}) string {
	s, ok := data.(string)
	if ok {
		return s
	}
	result, err := MarshalIndent(data, "", "    ")
	if err != nil {
		return ""
	}
	return string(result)
}
