package unify4g

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"time"
)

// ContainsN checks if a specified item is present within a given slice.
//
// This function iterates over a slice of any type that supports comparison
// and checks if the specified `item` exists within it. It returns `true`
// if the `item` is found and `false` otherwise.
//
// The function is generic, so it can be used with any comparable type,
// including strings, integers, floats, or custom types that implement the
// comparable interface.
//
// Parameters:
//   - `array`: The slice of elements to search through. This slice can
//     contain any type `T` that supports comparison (e.g., int, string).
//   - `item`: The item to search for within `array`. It should be of the
//     same type `T` as the elements in `array`.
//
// Returns:
//   - `true` if `item` is found within `array`, `false` otherwise.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	isPresent := ContainsN(numbers, 3) // isPresent will be true as 3 is in the slice
//
//	names := []string{"Alice", "Bob", "Charlie"}
//	isPresent := ContainsN(names, "Eve") // isPresent will be false as "Eve" is not in the slice
func ContainsN[T comparable](array []T, item T) bool {
	if len(array) == 0 {
		return false
	}
	for _, v := range array {
		if v == item {
			return true
		}
	}
	return false
}

// MapContainsKey checks if a specified key is present within a given map.
//
// This function takes a map with keys of any comparable type `K` and values of
// any type `V`. It checks if the specified `key` exists in the map `m`. If the key
// is found, it returns `true`; otherwise, it returns `false`.
//
// The function is generic and can be used with maps that have keys of any type
// that supports comparison (e.g., int, string). The value type `V` can be any type.
//
// Parameters:
//   - `m`: The map in which to search for the key. The map has keys of type `K`
//     and values of type `V`.
//   - `key`: The key to search for within `m`. It should be of the same type `K` as
//     the keys in `m`.
//
// Returns:
//   - `true` if `key` is found in `m`, `false` otherwise.
//
// Example:
//
//	ages := map[string]int{"Alice": 30, "Bob": 25}
//	isPresent := MapContainsKey(ages, "Alice") // isPresent will be true as "Alice" is a key in the map
//
//	prices := map[int]float64{1: 9.99, 2: 19.99}
//	isPresent := MapContainsKey(prices, 3) // isPresent will be false as 3 is not a key in the map
func MapContainsKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

// Filter returns a new slice containing only the elements from the input slice
// that satisfy a specified condition.
//
// This function iterates over each element in the input slice `list` and applies
// the provided `condition` function to it. If the `condition` function returns `true`
// for an element, that element is added to the `filtered` slice. At the end, the
// function returns the `filtered` slice containing only the elements that met the
// condition.
//
// The function is generic, allowing it to work with slices of any type `T` and
// any condition function that takes a `T` and returns a boolean.
//
// Parameters:
//   - `list`: The slice of elements to filter. It can contain elements of any type `T`.
//   - `condition`: A function that defines the filtering criteria. It takes an element
//     of type `T` as input and returns `true` if the element should be included in
//     the result, or `false` if it should be excluded.
//
// Returns:
//   - A new slice of type `[]T` containing only the elements from `list` for which
//     the `condition` function returned `true`.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	oddNumbers := Filter(numbers, func(n int) bool { return n%2 != 0 })
//	// oddNumbers will be []int{1, 3, 5} as only the odd numbers satisfy the condition
//
//	words := []string{"apple", "banana", "cherry"}
//	longWords := Filter(words, func(word string) bool { return len(word) > 5 })
//	// longWords will be []string{"banana", "cherry"} as they are longer than 5 characters
func Filter[T any](list []T, condition func(T) bool) []T {
	filtered := make([]T, 0)
	for _, item := range list {
		if condition(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

// Map returns a new slice where each element is the result of applying a specified
// transformation function to each element in the input slice.
//
// This function iterates over each element in the input slice `list`, applies the
// provided transformation function `f` to it, and stores the result in the new slice
// `result`. The length of the resulting slice is the same as the input slice, and
// each element in `result` corresponds to a transformed element from `list`.
//
// The function is generic, allowing it to work with slices of any type `T` and
// apply a transformation function that converts each element of type `T` to a
// new type `U`.
//
// Parameters:
//   - `list`: The slice of elements to transform. It can contain elements of any type `T`.
//   - `f`: A function that defines the transformation. It takes an element of type `T`
//     as input and returns a transformed value of type `U`.
//
// Returns:
//   - A new slice of type `[]U` where each element is the result of applying `f`
//     to the corresponding element in `list`.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	squaredNumbers := Map(numbers, func(n int) int { return n * n })
//	// squaredNumbers will be []int{1, 4, 9, 16, 25} as each number is squared
//
//	words := []string{"apple", "banana", "cherry"}
//	wordLengths := Map(words, func(word string) int { return len(word) })
//	// wordLengths will be []int{5, 6, 6} as each word's length is calculated
func Map[T any, U any](list []T, f func(T) U) []U {
	result := make([]U, len(list))
	for i, item := range list {
		result[i] = f(item)
	}
	return result
}

// Concat returns a new slice that is the result of concatenating multiple input slices
// into a single slice.
//
// This function takes a variable number of slices as input and combines them into
// one contiguous slice. It first calculates the total length needed for the resulting
// slice, then copies each input slice into the appropriate position within the
// resulting slice.
//
// The function is generic, allowing it to concatenate slices of any type `T`.
//
// Parameters:
//   - `slices`: A variadic parameter representing the slices to concatenate. Each slice
//     can contain elements of any type `T`, and they will be concatenated in the order
//     they are provided.
//
// Returns:
//   - A new slice of type `[]T` containing all elements from each input slice in sequence.
//
// Example:
//
//	// Concatenating integer slices
//	a := []int{1, 2}
//	b := []int{3, 4}
//	c := []int{5, 6}
//	combined := Concat(a, b, c)
//	// combined will be []int{1, 2, 3, 4, 5, 6}
//
//	// Concatenating string slices
//	words1 := []string{"hello", "world"}
//	words2 := []string{"go", "lang"}
//	concatenatedWords := Concat(words1, words2)
//	// concatenatedWords will be []string{"hello", "world", "go", "lang"}
func Concat[T any](slices ...[]T) []T {
	totalLen := 0
	for _, s := range slices {
		totalLen += len(s)
	}
	result := make([]T, totalLen)
	i := 0
	for _, s := range slices {
		copy(result[i:], s)
		i += len(s)
	}
	return result
}

// Sum calculates the sum of elements in a slice after transforming each element to a float64.
//
// This function iterates over each element in the input slice `slice`, applies a transformation
// function `transformer` to convert the element to a float64, and adds the result to a running
// total. The final sum is returned as a float64.
//
// The function is generic, allowing it to operate on slices of any type `T`. The `transformer`
// function is used to convert each element to a float64, enabling flexible summation of
// different types (e.g., integers, custom types with numeric properties).
//
// Parameters:
//   - `slice`: The slice of elements to sum. It can contain elements of any type `T`.
//   - `transformer`: A function that takes an element of type `T` and returns a float64
//     representation, which will be used in the summation.
//
// Returns:
//   - A float64 representing the sum of the transformed elements.
//
// Example:
//
//	// Summing integer slice values
//	numbers := []int{1, 2, 3, 4}
//	total := Sum(numbers, func(n int) float64 { return float64(n) })
//	// total will be 10.0 as each integer is converted to float64 and summed
//
//	// Summing custom struct values
//	type Product struct {
//	    Price float64
//	}
//	products := []Product{{Price: 9.99}, {Price: 19.99}, {Price: 5.0}}
//	totalPrice := Sum(products, func(p Product) float64 { return p.Price })
//	// totalPrice will be 34.98 as the prices are summed
func Sum[T any](slice []T, transformer func(T) float64) float64 {
	sum := 0.0
	for _, item := range slice {
		sum += transformer(item)
	}
	return sum
}

// Equal checks if two slices are equal in both length and elements.
//
// This function compares two slices `a` and `b` of any comparable type `T`. It first
// checks if the lengths of the two slices are the same. If they are not, it returns `false`.
// If the lengths match, it then iterates through each element in `a` and `b` to check
// if corresponding elements are equal. If all elements are equal, the function returns `true`;
// otherwise, it returns `false`.
//
// The function is generic and can be used with slices of any comparable type, such as
// integers, strings, or other types that support equality comparison.
//
// Parameters:
//   - `a`: The first slice to compare. It should contain elements of a comparable type `T`.
//   - `b`: The second slice to compare. It should also contain elements of type `T`.
//
// Returns:
//   - `true` if both slices have the same length and identical elements at each position;
//     `false` otherwise.
//
// Example:
//
//	// Comparing integer slices
//	a := []int{1, 2, 3}
//	b := []int{1, 2, 3}
//	isEqual := Equal(a, b)
//	// isEqual will be true as both slices contain the same elements in the same order
//
//	c := []int{1, 2, 4}
//	isEqual = Equal(a, c)
//	// isEqual will be false as the elements differ
//
//	// Comparing string slices
//	names1 := []string{"Alice", "Bob"}
//	names2 := []string{"Alice", "Bob"}
//	isEqual = Equal(names1, names2)
//	// isEqual will be true since the slices have identical elements
func Equal[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// SliceToMap converts a slice into a map, using a specified function to generate keys
// for each element in the slice.
//
// This function iterates over each element in the input slice `slice`, applies the
// provided `keyFunc` function to generate a key for each element, and then inserts the
// element into the resulting map `result` using that key. This allows the creation of
// a map from a slice, where each element is accessible via a unique key.
//
// The function is generic, allowing it to operate on slices of any type `T` and
// generate keys of any comparable type `K`. The resulting map will have keys of
// type `K` and values of type `T`.
//
// Parameters:
//   - `slice`: The slice of elements to convert to a map. It can contain elements of any type `T`.
//   - `keyFunc`: A function that takes an element of type `T` and returns a key of type `K`,
//     which is used as the key for each element in the resulting map.
//
// Returns:
//   - A map of type `map[K]T`, where each element in `slice` is inserted using the key
//     generated by `keyFunc`. If `keyFunc` generates the same key for multiple elements,
//     the last one will overwrite the previous entry in the map.
//
// Example:
//
//	// Converting a slice of strings to a map with string lengths as keys
//	words := []string{"apple", "banana", "cherry"}
//	wordMap := SliceToMap(words, func(word string) int { return len(word) })
//	// wordMap will be map[int]string{5: "apple", 6: "cherry"}
//	// Note: "banana" is overwritten by "cherry" as they have the same key 6
//
//	// Converting a slice of structs to a map using a struct field as the key
//	type Person struct {
//	    ID   int
//	    Name string
//	}
//	people := []Person{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}}
//	personMap := SliceToMap(people, func(p Person) int { return p.ID })
//	// personMap will be map[int]Person{1: {ID: 1, Name: "Alice"}, 2: {ID: 2, Name: "Bob"}}
func SliceToMap[T any, K comparable](slice []T, keyFunc func(T) K) map[K]T {
	result := make(map[K]T)
	for _, item := range slice {
		result[keyFunc(item)] = item
	}
	return result
}

// Reduce applies an accumulator function over a slice, producing a single accumulated result.
//
// This function iterates over each element in the input slice `slice`, applying the
// `accumulator` function to combine each element with an accumulated result. It starts
// with an initial value `initialValue` and successively updates the result by applying
// `accumulator` to each element in `slice`. The final accumulated result is returned
// once all elements have been processed.
//
// The function is generic, allowing it to operate on slices of any type `T` and produce
// an output of any type `U`. This enables flexible aggregation operations such as
// summing, counting, or accumulating data into more complex structures.
//
// Parameters:
//   - `slice`: The slice of elements to reduce. It can contain elements of any type `T`.
//   - `accumulator`: A function that takes the current accumulated result of type `U`
//     and an element of type `T`, then returns the updated accumulated result of type `U`.
//   - `initialValue`: The initial value for the accumulator, of type `U`. This is the
//     starting point for the reduction process.
//
// Returns:
//   - The final accumulated result of type `U` after applying `accumulator` to each element
//     in `slice`.
//
// Example:
//
//	// Summing integer values in a slice
//	numbers := []int{1, 2, 3, 4}
//	sum := Reduce(numbers, func(acc, n int) int { return acc + n }, 0)
//	// sum will be 10 as each integer is added to the accumulated result
//
//	// Concatenating strings in a slice
//	words := []string{"go", "is", "fun"}
//	sentence := Reduce(words, func(acc, word string) string { return acc + " " + word }, "")
//	// sentence will be " go is fun" (note leading space due to initial value being "")
//
//	// Using a custom struct and custom accumulator
//	type Product struct {
//	    Name  string
//	    Price float64
//	}
//	products := []Product{{Name: "apple", Price: 0.99}, {Name: "banana", Price: 1.29}}
//	totalPrice := Reduce(products, func(total float64, p Product) float64 { return total + p.Price }, 0.0)
//	// totalPrice will be 2.28 as each product's price is added to the accumulated total
func Reduce[T any, U any](slice []T, accumulator func(U, T) U, initialValue U) U {
	result := initialValue
	for _, item := range slice {
		result = accumulator(result, item)
	}
	return result
}

// IndexOf searches for a specific element in a slice and returns its index if found.
//
// This function iterates over each element in the input slice `slice` to find the first
// occurrence of the specified `item`. If `item` is found, the function returns the index
// of `item` within `slice`. If `item` is not present in the slice, it returns -1.
//
// The function is generic, allowing it to operate on slices of any comparable type `T`
// (e.g., int, string, or other types that support equality comparison).
//
// Parameters:
//   - `slice`: The slice in which to search for `item`. It can contain elements of any
//     comparable type `T`.
//   - `item`: The item to search for within `slice`. It should be of the same type `T`
//     as the elements in `slice`.
//
// Returns:
//   - The zero-based index of `item` in `slice` if it exists; otherwise, -1.
//
// Example:
//
//	// Searching for an integer in a slice
//	numbers := []int{1, 2, 3, 4}
//	index := IndexOf(numbers, 3)
//	// index will be 2, as 3 is located at index 2 in the slice
//
//	// Searching for a string in a slice
//	words := []string{"apple", "banana", "cherry"}
//	index = IndexOf(words, "banana")
//	// index will be 1, as "banana" is at index 1 in the slice
//
//	// Item not found in the slice
//	index = IndexOf(words, "date")
//	// index will be -1, as "date" is not in the slice
func IndexOf[T comparable](slice []T, item T) int {
	for i, value := range slice {
		if value == item {
			return i
		}
	}
	return -1
}

// Unique returns a new slice containing only the unique elements from the input slice,
// preserving their original order.
//
// This function iterates over each element in the input slice `slice` and uses a map
// `uniqueMap` to track elements that have already been encountered. If an element has
// not been seen before, it is added to both the `uniqueValues` result slice and the
// map. This ensures that only the first occurrence of each unique element is kept in
// the final slice, while duplicates are ignored.
//
// The function is generic, allowing it to operate on slices of any comparable type `T`.
// The elements must be of a comparable type to allow them to be used as keys in the map.
//
// Parameters:
//   - `slice`: The input slice from which unique elements are extracted. It can contain
//     elements of any comparable type `T`.
//
// Returns:
//   - A new slice of type `[]T` containing only the unique elements from `slice` in the
//     order of their first appearance.
//
// Example:
//
//	// Extracting unique integers from a slice
//	numbers := []int{1, 2, 2, 3, 4, 4, 5}
//	uniqueNumbers := Unique(numbers)
//	// uniqueNumbers will be []int{1, 2, 3, 4, 5}
//
//	// Extracting unique strings from a slice
//	words := []string{"apple", "banana", "apple", "cherry"}
//	uniqueWords := Unique(words)
//	// uniqueWords will be []string{"apple", "banana", "cherry"}
//
//	// An empty slice will return an empty result
//	empty := []int{}
//	uniqueEmpty := Unique(empty)
//	// uniqueEmpty will be []int{}
func Unique[T comparable](slice []T) []T {
	uniqueMap := make(map[T]bool)
	uniqueValues := make([]T, 0)
	for _, value := range slice {
		if _, found := uniqueMap[value]; !found {
			uniqueValues = append(uniqueValues, value)
			uniqueMap[value] = true
		}
	}
	return uniqueValues
}

// Flatten takes a slice of potentially nested elements and returns a new slice
// containing all elements of type `T` in a flat structure.
//
// This function recursively processes each element in the input slice `s`, checking if
// it is a nested slice (`[]interface{}`). If a nested slice is found, `Flatten` is called
// recursively to flatten it and append its elements to the `result` slice. If an element
// is of type `T`, it is directly appended to `result`. Elements that are neither `[]interface{}`
// nor of type `T` are ignored.
//
// The function is generic, allowing it to work with any element type `T`, which must be
// specified when calling the function. This makes `Flatten` useful for flattening slices
// with nested structures while filtering only the elements of a specified type.
//
// Parameters:
//   - `s`: A slice of `interface{}`, which can contain nested slices (`[]interface{}`) or
//     elements of any type. Nested slices may contain more nested slices at arbitrary depths.
//
// Returns:
//   - A new slice of type `[]T` containing all elements of type `T` from `s`, flattened
//     into a single level.
//
// Example:
//
//	// Flattening a nested slice of integers
//	nestedInts := []interface{}{1, []interface{}{2, 3}, []interface{}{[]interface{}{4, 5}}}
//	flatInts := Flatten[int](nestedInts)
//	// flatInts will be []int{1, 2, 3, 4, 5}
//
//	// Flattening a nested slice with mixed types, extracting only strings
//	mixedNested := []interface{}{"apple", []interface{}{"banana", 1, []interface{}{"cherry"}}}
//	flatStrings := Flatten[string](mixedNested)
//	// flatStrings will be []string{"apple", "banana", "cherry"}
//
//	// Flattening an empty slice
//	empty := []interface{}{}
//	flatEmpty := Flatten[int](empty)
//	// flatEmpty will be []int{}
func Flatten[T any](s []interface{}) []T {
	result := make([]T, 0)
	for _, v := range s {
		switch val := v.(type) {
		case []interface{}:
			result = append(result, Flatten[T](val)...)
		default:
			if _, ok := val.(T); ok {
				result = append(result, val.(T))
			}
		}
	}
	return result
}

// DeepEqual compares two values of any comparable type to determine if they are deeply equal.
//
// This function uses the `reflect.DeepEqual` function from the `reflect` package to compare
// two values `a` and `b`. It checks for deep equality, meaning it considers nested structures,
// such as slices, maps, or structs, and compares them element-by-element or field-by-field.
// If the values are deeply equal, the function returns `true`; otherwise, it returns `false`.
//
// The function is generic, allowing it to work with any type `T` that is comparable, including
// basic types (e.g., integers, strings) as well as complex types with nested structures.
//
// Parameters:
//   - `a`: The first value to compare. It can be of any comparable type `T`.
//   - `b`: The second value to compare. It must be of the same type `T` as `a`.
//
// Returns:
//   - `true` if `a` and `b` are deeply equal; `false` otherwise.
//
// Example:
//
//	// Comparing two integer values
//	isEqual := DeepEqual(5, 5)
//	// isEqual will be true as both integers are equal
//
//	// Comparing two slices with the same elements
//	sliceA := []int{1, 2, 3}
//	sliceB := []int{1, 2, 3}
//	isEqual = DeepEqual(sliceA, sliceB)
//	// isEqual will be true as both slices have identical elements in the same order
//
//	// Comparing two different maps
//	mapA := map[string]int{"a": 1, "b": 2}
//	mapB := map[string]int{"a": 1, "b": 3}
//	isEqual = DeepEqual(mapA, mapB)
//	// isEqual will be false as the values for key "b" differ between the maps
func DeepEqual[T comparable](a, b T) bool {
	return reflect.DeepEqual(a, b)
}

// GroupBy groups elements of a slice into a map based on a specified key.
//
// This function iterates over each element in the input slice `slice`, applies the
// provided `getKey` function to extract a key for each element, and groups elements
// that share the same key into a slice. The function then returns a map where each
// key maps to a slice of elements that correspond to that key.
//
// The function is generic, allowing it to work with slices of any type `T` and to
// generate keys of any comparable type `K`. This makes `GroupBy` useful for organizing
// data based on shared attributes, such as grouping items by category or organizing
// records by a specific field.
//
// Parameters:
//   - `slice`: The input slice containing elements to be grouped. It can be of any type `T`.
//   - `getKey`: A function that takes an element of type `T` and returns a key of type `K`,
//     which is used to group the element in the resulting map.
//
// Returns:
//   - A map of type `map[K][]T`, where each key is associated with a slice of elements
//     that share that key.
//
// Example:
//
//	// Grouping integers by even and odd
//	numbers := []int{1, 2, 3, 4, 5}
//	grouped := GroupBy(numbers, func(n int) string {
//		if n%2 == 0 {
//			return "even"
//		}
//		return "odd"
//	})
//	// grouped will be map[string][]int{"even": {2, 4}, "odd": {1, 3, 5}}
//
//	// Grouping people by age
//	type Person struct {
//	    Name string
//	    Age  int
//	}
//	people := []Person{{Name: "Alice", Age: 30}, {Name: "Bob", Age: 25}, {Name: "Charlie", Age: 30}}
//	groupedByAge := GroupBy(people, func(p Person) int { return p.Age })
//	// groupedByAge will be map[int][]Person{30: {{Name: "Alice", Age: 30}, {Name: "Charlie", Age: 30}}, 25: {{Name: "Bob", Age: 25}}}
//
//	// Grouping strings by their length
//	words := []string{"apple", "pear", "banana", "peach"}
//	groupedByLength := GroupBy(words, func(word string) int { return len(word) })
//	// groupedByLength will be map[int][]string{5: {"apple", "peach"}, 4: {"pear"}, 6: {"banana"}}
func GroupBy[T any, K comparable](slice []T, getKey func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, item := range slice {
		key := getKey(item)
		result[key] = append(result[key], item)
	}
	return result
}

// FlattenDeep takes a nested structure of arbitrary depth and returns a flat slice
// containing all elements in a single level.
//
// This function recursively processes each element in `arr`. If an element is itself a
// slice (`[]interface{}`), `FlattenDeep` calls itself to flatten that nested slice and
// appends its elements to the `result` slice. If the element is not a slice, it is directly
// added to `result`. The function allows flattening of complex nested structures while
// maintaining all elements in a single-level output.
//
// This function operates with values of type `interface{}`, making it flexible enough
// to handle mixed types in the input. It returns a slice of `interface{}`, which may
// contain elements of varying types from the original nested structure.
//
// Parameters:
//   - `arr`: The input slice, which can contain nested slices of arbitrary depth and elements
//     of any type.
//
// Returns:
//   - A slice of `[]interface{}` containing all elements from `arr` flattened into a single level.
//
// Example:
//
//	// Flattening a nested structure of mixed values
//	nested := []interface{}{1, []interface{}{2, 3, []interface{}{4, []interface{}{5}}}}
//	flat := FlattenDeep(nested)
//	// flat will be []interface{}{1, 2, 3, 4, 5}
//
//	// Flattening a deeply nested structure with varied types
//	mixedNested := []interface{}{"apple", []interface{}{"banana", 1, []interface{}{"cherry"}}}
//	flatMixed := FlattenDeep(mixedNested)
//	// flatMixed will be []interface{}{"apple", "banana", 1, "cherry"}
//
//	// Flattening a non-nested input returns the input as-is
//	nonNested := 5
//	flatNonNested := FlattenDeep(nonNested)
//	// flatNonNested will be []interface{}{5}
func FlattenDeep(arr interface{}) []interface{} {
	result := make([]interface{}, 0)
	switch v := arr.(type) {
	case []interface{}:
		for _, val := range v {
			result = append(result, FlattenDeep(val)...)
		}
	case interface{}:
		result = append(result, v)
	}
	return result
}

// Join concatenates the string representation of each element in a slice into a single
// string, with a specified separator between each element.
//
// This function iterates over each element in the input slice `slice`, converts each
// element to a string using `fmt.Sprintf` with the `%v` format, and appends it to the
// `result` string. A separator string `separator` is inserted between elements in the
// final concatenated result. If the slice has only one element, no separator is added.
// The function is generic and can work with slices containing elements of any type `T`.
//
// Parameters:
//   - `slice`: The input slice containing elements to be joined. It can contain elements
//     of any type `T`.
//   - `separator`: A string that will be inserted between each element in the final result.
//
// Returns:
//   - A single string that is the result of concatenating all elements in `slice` with the
//     specified `separator` in between.
//
// Example:
//
//	// Joining integers with a comma separator
//	numbers := []int{1, 2, 3}
//	joinedNumbers := Join(numbers, ", ")
//	// joinedNumbers will be "1, 2, 3"
//
//	// Joining strings with a space separator
//	words := []string{"Go", "is", "awesome"}
//	joinedWords := Join(words, " ")
//	// joinedWords will be "Go is awesome"
//
//	// Joining an empty slice returns an empty string
//	emptySlice := []int{}
//	joinedEmpty := Join(emptySlice, ",")
//	// joinedEmpty will be ""
func Join[T any](slice []T, separator string) string {
	result := ""
	for i, item := range slice {
		if i > 0 {
			result += separator
		}
		result += fmt.Sprintf("%v", item)
	}
	return result
}

// ReverseN reverses the order of elements in the input slice and returns a new slice
// containing the elements in reverse order.
//
// This function creates a new slice `reversed` with the same length as the input slice `slice`.
// It then uses a two-pointer approach to swap the elements of `slice` from both ends toward the center,
// effectively reversing the slice. The result is a new slice with elements in the opposite order.
//
// The function is generic, allowing it to work with slices of any type `T`.
//
// Parameters:
//   - `slice`: The input slice whose elements are to be reversed. It can contain elements of any type `T`.
//
// Returns:
//   - A new slice of type `[]T` containing the elements of `slice` in reverse order.
//
// Example:
//
//	// Reversing a slice of integers
//	numbers := []int{1, 2, 3, 4}
//	reversedNumbers := ReverseN(numbers)
//	// reversedNumbers will be []int{4, 3, 2, 1}
//
//	// Reversing a slice of strings
//	words := []string{"apple", "banana", "cherry"}
//	reversedWords := ReverseN(words)
//	// reversedWords will be []string{"cherry", "banana", "apple"}
//
//	// Reversing an empty slice returns an empty slice
//	empty := []int{}
//	reversedEmpty := ReverseN(empty)
//	// reversedEmpty will be []int{}
func ReverseN[T any](slice []T) []T {
	reversed := make([]T, len(slice))
	for i, j := 0, len(slice)-1; i <= j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = slice[j], slice[i]
	}
	return reversed
}

// FindIndex searches for the first occurrence of a target element in a slice
// and returns its index. If the element is not found, it returns -1.
//
// This function iterates over each element in the input slice `slice` and compares
// each element to the specified `target`. When the first occurrence of `target` is found,
// the function returns the index of that element. If the element is not found, the function
// returns -1 to indicate that the target is not present in the slice.
//
// The function is generic, allowing it to work with slices of any comparable type `T`,
// such as integers, strings, or other types that support equality comparison.
//
// Parameters:
//   - `slice`: The input slice in which to search for the target element. It can contain
//     elements of any comparable type `T`.
//   - `target`: The element to search for within the slice. It should be of the same type `T`
//     as the elements in `slice`.
//
// Returns:
//   - The zero-based index of the first occurrence of `target` in the slice if it exists;
//     otherwise, -1 if the target is not found.
//
// Example:
//
//	// Searching for an integer in a slice
//	numbers := []int{1, 2, 3, 4}
//	index := FindIndex(numbers, 3)
//	// index will be 2, as 3 is located at index 2 in the slice
//
//	// Searching for a string in a slice
//	words := []string{"apple", "banana", "cherry"}
//	index = FindIndex(words, "banana")
//	// index will be 1, as "banana" is at index 1 in the slice
//
//	// Item not found in the slice
//	index = FindIndex(words, "date")
//	// index will be -1, as "date" is not in the slice
func FindIndex[T comparable](slice []T, target T) int {
	for i, item := range slice {
		if item == target {
			return i
		}
	}
	return -1
}

// MapToSlice applies a mapping function to each element in the input slice
// and returns a new slice containing the results of the mapping.
//
// This function iterates over each element in the input slice `slice` and applies
// the provided `mapper` function to each element. The result of applying the mapping
// function to each element is stored in a new slice `mappedSlice`. This allows for
// transforming the elements of the input slice into a new slice of a different type.
//
// The function is generic, allowing it to work with slices of any type `T` as input,
// and the result can be a slice of any type `U`.
//
// Parameters:
//   - `slice`: The input slice containing elements of type `T` to be mapped.
//   - `mapper`: A function that takes an element of type `T` and returns a transformed
//     element of type `U`.
//
// Returns:
//   - A new slice of type `[]U` containing the mapped elements, with the same length
//     as the input slice, but with elements transformed according to the `mapper` function.
//
// Example:
//
//	// Mapping a slice of integers to their string representations
//	numbers := []int{1, 2, 3}
//	mappedStrings := MapToSlice(numbers, func(n int) string {
//		return fmt.Sprintf("Number %d", n)
//	})
//	// mappedStrings will be []string{"Number 1", "Number 2", "Number 3"}
//
//	// Mapping a slice of strings to their lengths
//	words := []string{"apple", "banana", "cherry"}
//	mappedLengths := MapToSlice(words, func(word string) int {
//		return len(word)
//	})
//	// mappedLengths will be []int{5, 6, 6}
//
//	// Mapping an empty slice returns an empty slice
//	empty := []int{}
//	mappedEmpty := MapToSlice(empty, func(n int) string {
//		return fmt.Sprintf("Number %d", n)
//	})
//	// mappedEmpty will be []string{}
func MapToSlice[T any, U any](slice []T, mapper func(T) U) []U {
	mappedSlice := make([]U, len(slice))
	for i, item := range slice {
		mappedSlice[i] = mapper(item)
	}
	return mappedSlice
}

// MergeMaps combines multiple maps into a single map. If there are any key conflicts,
// the value from the last map will be used.
//
// This function accepts a variable number of maps of type `map[interface{}]V` and merges
// them into a single map. It iterates through each input map, adding all key-value pairs
// to the `mergedMap`. If a key already exists in `mergedMap`, the corresponding value
// from the current map will overwrite the existing value. The function returns the merged map.
//
// The function is generic, allowing it to work with maps where the key is of any type `K`
// and the value is of any type `V`. It uses `interface{}` as the key type, enabling it to
// handle a variety of key types, though this may require careful handling of the key types
// to ensure they are comparable if needed.
//
// Parameters:
//   - `maps`: A variadic parameter representing multiple maps to be merged. Each map has keys
//     of type `interface{}` and values of type `V`.
//
// Returns:
//   - A new map of type `map[interface{}]V` containing the merged key-value pairs. If there
//     are key conflicts, the last map's value will be used.
//
// Example:
//
//	// Merging two maps with integer keys and string values
//	map1 := map[interface{}]string{"a": "apple", "b": "banana"}
//	map2 := map[interface{}]string{"b": "blueberry", "c": "cherry"}
//	merged := MergeMaps(map1, map2)
//	// merged will be map[interface{}]string{"a": "apple", "b": "blueberry", "c": "cherry"}
//
//	// Merging maps with different value types (e.g., int and string)
//	map3 := map[interface{}]int{"x": 10, "y": 20}
//	map4 := map[interface{}]string{"y": "yellow", "z": "zebra"}
//	mergedMixed := MergeMaps(map3, map4)
//	// mergedMixed will be map[interface{}]string{"x": "10", "y": "yellow", "z": "zebra"}
//
//	// Merging an empty slice of maps returns an empty map
//	mergedEmpty := MergeMaps()
//	// mergedEmpty will be an empty map
func MergeMaps[K any, V any](maps ...map[interface{}]V) map[interface{}]V {
	mergedMap := make(map[interface{}]V)
	for _, m := range maps {
		for k, v := range m {
			mergedMap[k] = v
		}
	}
	return mergedMap
}

// FilterMap filters the key-value pairs of a map based on a condition provided by the filter function.
//
// This function iterates over each key-value pair in the input map `m` and applies the provided
// `filter` function to the value. If the `filter` function returns `true` for a value, that key-value
// pair is added to the `filteredMap`. Otherwise, the pair is excluded. The function returns a new map
// containing only the key-value pairs that satisfy the condition specified in the `filter` function.
//
// The function is generic, allowing it to work with maps where the keys and values can be of any type `K`
// and `V`, respectively.
//
// Parameters:
//   - `m`: The input map to be filtered, with keys of type `any` and values of type `V`.
//   - `filter`: A function that takes a value of type `V` and returns a boolean. It determines
//     whether the corresponding key-value pair should be included in the result map.
//
// Returns:
//   - A new map of type `map[any]V`, containing only the key-value pairs for which the `filter`
//     function returned `true`.
//
// Example:
//
//	// Filtering a map of integers, keeping only values greater than 10
//	map1 := map[any]int{"a": 5, "b": 15, "c": 20}
//	filtered := FilterMap(map1, func(v int) bool {
//		return v > 10
//	})
//	// filtered will be map[any]int{"b": 15, "c": 20}
//
//	// Filtering a map of strings, keeping only values with length greater than 3
//	map2 := map[any]string{"a": "apple", "b": "banana", "c": "cat"}
//	filteredStrings := FilterMap(map2, func(v string) bool {
//		return len(v) > 3
//	})
//	// filteredStrings will be map[any]string{"a": "apple", "b": "banana"}
//
//	// Filtering an empty map returns an empty map
//	emptyMap := map[any]int{}
//	filteredEmpty := FilterMap(emptyMap, func(v int) bool {
//		return v > 10
//	})
//	// filteredEmpty will be an empty map
func FilterMap[K any, V any](m map[any]V, filter func(V) bool) map[any]V {
	filteredMap := make(map[any]V)
	for k, v := range m {
		if filter(v) {
			filteredMap[k] = v
		}
	}
	return filteredMap
}

// Chunk splits a slice into smaller slices (chunks) of the specified size.
//
// This function takes an input slice `slice` and a `chunkSize` and splits the input slice into
// smaller slices, each containing up to `chunkSize` elements. The function returns a slice of slices
// containing the chunked elements. If the `chunkSize` is greater than the length of the input slice,
// the entire slice will be returned as a single chunk. If the `chunkSize` is less than or equal to 0,
// the function returns `nil`.
//
// The function is generic, allowing it to work with slices of any type `T`.
//
// Parameters:
//   - `slice`: The input slice to be split into chunks. It can contain elements of any type `T`.
//   - `chunkSize`: The size of each chunk. If this value is less than or equal to 0, the function returns `nil`.
//
// Returns:
//   - A slice of slices (`[][]T`), where each inner slice contains up to `chunkSize` elements from
//     the original slice. If the slice cannot be split into even chunks, the last chunk may contain
//     fewer elements than `chunkSize`.
//
// Example:
//
//	// Chunking a slice of integers into chunks of size 2
//	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//	chunks := Chunk(numbers, 2)
//	// chunks will be [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9}}
//
//	// Chunking a slice of strings into chunks of size 3
//	words := []string{"apple", "banana", "cherry", "date", "elderberry", "fig"}
//	chunksWords := Chunk(words, 3)
//	// chunksWords will be [][]string{{"apple", "banana", "cherry"}, {"date", "elderberry", "fig"}}
//
//	// Chunking an empty slice returns an empty slice of slices
//	empty := []int{}
//	chunksEmpty := Chunk(empty, 3)
//	// chunksEmpty will be [][]int{}
//
//	// If chunkSize is 0 or negative, return nil
//	chunksInvalid := Chunk(numbers, -1)
//	// chunksInvalid will be nil
func Chunk[T any](slice []T, chunkSize int) [][]T {
	if chunkSize <= 0 {
		return nil
	}
	var chunks [][]T
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// Values extracts and returns the values from a map as a slice.
//
// This function takes a map `m` with keys of type `any` and values of type `V`, and creates
// a new slice containing all the values from the map. The function iterates over the map and
// appends each value to the `values` slice. The resulting slice will have the same number of
// elements as the map has key-value pairs, and the order of values will correspond to the
// order in which they were iterated over (which is not guaranteed to be in any particular order).
//
// The function is generic, allowing it to work with maps of any key type `K` and value type `V`.
//
// Parameters:
//   - `m`: The input map from which to extract the values. The keys are of type `any`
//     and the values are of type `V`.
//
// Returns:
//   - A slice of type `[]V` containing all the values from the map `m`.
//
// Example:
//
//	// Extracting values from a map of strings to integers
//	map1 := map[any]int{"a": 1, "b": 2, "c": 3}
//	values := Values(map1)
//	// values will be []int{1, 2, 3}
//
//	// Extracting values from a map of strings to strings
//	map2 := map[any]string{"x": "apple", "y": "banana", "z": "cherry"}
//	valuesStrings := Values(map2)
//	// valuesStrings will be []string{"apple", "banana", "cherry"}
//
//	// Extracting values from an empty map returns an empty slice
//	emptyMap := map[any]int{}
//	emptyValues := Values(emptyMap)
//	// emptyValues will be []int{}
func Values[K any, V any](m map[any]V) []V {
	values := make([]V, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

// Shuffle randomly shuffles the elements of a slice and returns a new slice with the shuffled elements.
//
// This function takes an input slice `slice` and shuffles its elements using a random permutation.
// It creates a new slice `shuffledSlice` and populates it by selecting elements from the original slice
// according to a random permutation of indices. The resulting `shuffledSlice` contains the same elements
// as the input slice, but in a random order. The function uses a seeded random generator to ensure different
// results each time it is called.
//
// The function is generic, allowing it to work with slices of any type `T`.
//
// Parameters:
//   - `slice`: The input slice to be shuffled. It can contain elements of any type `T`.
//
// Returns:
//   - A new slice of type `[]T`, containing the shuffled elements of the input slice.
//
// Example:
//
//	// Shuffling a slice of integers
//	numbers := []int{1, 2, 3, 4, 5}
//	shuffledNumbers := Shuffle(numbers)
//	// shuffledNumbers will be a random permutation of [1, 2, 3, 4, 5]
//
//	// Shuffling a slice of strings
//	words := []string{"apple", "banana", "cherry"}
//	shuffledWords := Shuffle(words)
//	// shuffledWords will be a random permutation of ["apple", "banana", "cherry"]
//
//	// Shuffling an empty slice returns an empty slice
//	empty := []int{}
//	shuffledEmpty := Shuffle(empty)
//	// shuffledEmpty will be []int{}
func Shuffle[T any](slice []T) []T {
	shuffledSlice := make([]T, len(slice))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(slice))
	for i, randIndex := range perm {
		shuffledSlice[i] = slice[randIndex]
	}
	return shuffledSlice
}

// CartesianProduct computes the Cartesian product of multiple slices and returns the result as a slice of slices.
//
// This function takes multiple slices of type `[]T` and computes their Cartesian product. The Cartesian
// product of two or more sets (or slices in this case) is the set of all possible combinations where each
// combination consists of one element from each slice. The function recursively computes the product of the
// slices, starting from the second slice and combining it with each element of the first slice. The result is
// a slice of slices, where each inner slice is a combination of elements from the input slices.
//
// The function is generic, allowing it to work with slices of any type `T`.
//
// Parameters:
//   - `slices`: A variadic parameter that represents multiple slices to compute the Cartesian product of.
//     Each slice can contain elements of any type `T`.
//
// Returns:
//   - A slice of slices (`[][]T`), where each inner slice represents a unique combination of elements
//     from the input slices.
//
// Example:
//
//	// Cartesian product of two slices of integers
//	slice1 := []int{1, 2}
//	slice2 := []int{3, 4}
//	product := CartesianProduct(slice1, slice2)
//	// product will be [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}}
//
//	// Cartesian product of three slices of strings
//	slice3 := []string{"a", "b"}
//	slice4 := []string{"x", "y"}
//	slice5 := []string{"1", "2"}
//	productStrings := CartesianProduct(slice3, slice4, slice5)
//	// productStrings will be [][]string{
//	//   {"a", "x", "1"}, {"a", "x", "2"},
//	//   {"a", "y", "1"}, {"a", "y", "2"},
//	//   {"b", "x", "1"}, {"b", "x", "2"},
//	//   {"b", "y", "1"}, {"b", "y", "2"}
//	// }
//
//	// Cartesian product of an empty slice returns an empty slice
//	empty := []int{}
//	productEmpty := CartesianProduct(empty)
//	// productEmpty will be [][]int{{}}
func CartesianProduct[T any](slices ...[]T) [][]T {
	n := len(slices)
	if n == 0 {
		return [][]T{{}}
	}
	if n == 1 {
		product := make([][]T, len(slices[0]))
		for i, item := range slices[0] {
			product[i] = []T{item}
		}
		return product
	}
	tailProduct := CartesianProduct(slices[1:]...)
	product := make([][]T, 0, len(slices[0])*len(tailProduct))
	for _, head := range slices[0] {
		for _, tail := range tailProduct {
			product = append(product, append([]T{head}, tail...))
		}
	}
	return product
}

// Sort sorts a slice based on a custom comparison function and returns a new sorted slice.
//
// This function takes an input slice `slice` and a comparison function `comparer` that defines the
// sorting order. The comparison function takes two elements of type `T` and returns `true` if the
// first element should come before the second one (i.e., if it should be sorted earlier). The function
// creates a new slice `sortedSlice` by copying the elements of the original slice, then sorts it in place
// using the provided `comparer`. The resulting `sortedSlice` will be a new slice containing the elements
// from the original slice, arranged in the order specified by the comparison function.
//
// The function is generic, allowing it to work with slices of any type `T`.
//
// Parameters:
//   - `slice`: The input slice to be sorted. It can contain elements of any type `T`.
//   - `comparer`: A comparison function that takes two elements of type `T` and returns a boolean value.
//     It determines the order of the elements: it should return `true` if the first element should come
//     before the second element in the sorted order.
//
// Returns:
//   - A new slice of type `[]T`, containing the elements from the original slice sorted according to
//     the provided comparison function.
//
// Example:
//
//	// Sorting a slice of integers in ascending order
//	numbers := []int{5, 3, 8, 1, 2}
//	sortedNumbers := Sort(numbers, func(a, b int) bool {
//		return a < b
//	})
//	// sortedNumbers will be []int{1, 2, 3, 5, 8}
//
//	// Sorting a slice of strings in descending order
//	words := []string{"apple", "banana", "cherry"}
//	sortedWords := Sort(words, func(a, b string) bool {
//		return a > b
//	})
//	// sortedWords will be []string{"cherry", "banana", "apple"}
//
//	// Sorting an empty slice returns an empty slice
//	empty := []int{}
//	sortedEmpty := Sort(empty, func(a, b int) bool {
//		return a < b
//	})
//	// sortedEmpty will be []int{}
func Sort[T any](slice []T, comparer func(T, T) bool) []T {
	sortedSlice := make([]T, len(slice))
	copy(sortedSlice, slice)
	sort.Slice(sortedSlice, func(i, j int) bool {
		return comparer(sortedSlice[i], sortedSlice[j])
	})
	return sortedSlice
}

// AllMatch checks if all elements in a slice satisfy a given condition and returns a boolean result.
//
// This function takes an input slice `slice` and a predicate function `predicate`. It iterates over
// each element in the slice, applying the predicate function to each one. If any element does not
// satisfy the predicate (i.e., the predicate returns `false`), the function immediately returns `false`.
// If all elements satisfy the predicate, the function returns `true`.
//
// The function is generic, allowing it to work with slices of any type `T`.
//
// Parameters:
//   - `slice`: The input slice whose elements will be checked. It can contain elements of any type `T`.
//   - `predicate`: A function that takes an element of type `T` and returns a boolean. This function
//     represents the condition that each element must satisfy. If the predicate returns `true` for an
//     element, the element meets the condition.
//
// Returns:
//   - `true` if all elements in the slice satisfy the predicate; `false` if any element does not.
//
// Example:
//
//	// Checking if all integers in a slice are positive
//	numbers := []int{2, 4, 6, 8}
//	allPositive := AllMatch(numbers, func(n int) bool {
//		return n > 0
//	})
//	// allPositive will be true
//
//	// Checking if all strings in a slice have a length greater than 3
//	words := []string{"apple", "banana", "pear"}
//	allLong := AllMatch(words, func(s string) bool {
//		return len(s) > 3
//	})
//	// allLong will be true
//
//	// If the slice is empty, returns true since no elements violate the predicate
//	empty := []int{}
//	allMatchEmpty := AllMatch(empty, func(n int) bool {
//		return n > 0
//	})
//	// allMatchEmpty will be true
func AllMatch[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// AnyMatch checks if any element in a slice satisfies a given condition and returns a boolean result.
//
// This function takes an input slice `slice` and a predicate function `predicate`. It iterates over
// each element in the slice, applying the predicate function to each one. If the predicate returns
// `true` for any element, the function immediately returns `true`. If no elements satisfy the predicate,
// the function returns `false`.
//
// The function is generic, allowing it to work with slices of any type `T`.
//
// Parameters:
//   - `slice`: The input slice whose elements will be checked. It can contain elements of any type `T`.
//   - `predicate`: A function that takes an element of type `T` and returns a boolean. This function
//     represents the condition that an element must satisfy. If the predicate returns `true` for an
//     element, the element meets the condition.
//
// Returns:
//   - `true` if at least one element in the slice satisfies the predicate; `false` if no elements do.
//
// Example:
//
//	// Checking if any integers in a slice are even
//	numbers := []int{1, 3, 5, 6}
//	anyEven := AnyMatch(numbers, func(n int) bool {
//		return n%2 == 0
//	})
//	// anyEven will be true because 6 is even
//
//	// Checking if any strings in a slice contain the letter "a"
//	words := []string{"apple", "banana", "cherry"}
//	containsA := AnyMatch(words, func(s string) bool {
//		return strings.Contains(s, "a")
//	})
//	// containsA will be true because "apple" and "banana" contain "a"
//
//	// Checking an empty slice returns false since no elements satisfy the predicate
//	empty := []int{}
//	anyMatchEmpty := AnyMatch(empty, func(n int) bool {
//		return n > 0
//	})
//	// anyMatchEmpty will be false
func AnyMatch[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if predicate(item) {
			return true
		}
	}
	return false
}

// Push appends an element to the end of a slice and returns the resulting slice.
//
// This function takes an input slice `slice` and an element `element`, and appends
// the element to the end of the slice using the built-in `append` function. The
// function returns a new slice containing the original elements followed by the new
// element. This function is useful for adding elements dynamically to a slice.
//
// The function is generic, allowing it to work with slices of any type `T`.
//
// Parameters:
//   - `slice`: The input slice to which the element will be appended. It can contain elements of any type `T`.
//   - `element`: The element to be appended to the end of the slice. It is of type `T`.
//
// Returns:
//   - A new slice of type `[]T`, containing the original elements in `slice` with `element` appended at the end.
//
// Example:
//
//	// Appending an integer to a slice of integers
//	numbers := []int{1, 2, 3}
//	updatedNumbers := Push(numbers, 4)
//	// updatedNumbers will be []int{1, 2, 3, 4}
//
//	// Appending a string to a slice of strings
//	words := []string{"apple", "banana"}
//	updatedWords := Push(words, "cherry")
//	// updatedWords will be []string{"apple", "banana", "cherry"}
//
//	// Appending to an empty slice
//	var emptySlice []int
//	newSlice := Push(emptySlice, 1)
//	// newSlice will be []int{1}
func Push[T any](slice []T, element T) []T {
	return append(slice, element)
}

// Pop removes the last element from a slice and returns the resulting slice.
//
// This function takes an input slice `slice` and removes its last element by creating
// a new slice that excludes the last element. The function uses slicing to return a
// portion of the original slice that ends before the last element. If the input slice
// is empty, calling this function will result in a runtime panic.
//
// The function is generic, allowing it to work with slices of any type `T`.
//
// Parameters:
//   - `slice`: The input slice from which the last element will be removed. It can contain elements of any type `T`.
//
// Returns:
//   - A new slice of type `[]T`, containing all elements from the original slice except the last one.
//
// Example:
//
//	// Removing the last element from a slice of integers
//	numbers := []int{1, 2, 3, 4}
//	updatedNumbers := Pop(numbers)
//	// updatedNumbers will be []int{1, 2, 3}
//
//	// Removing the last element from a slice of strings
//	words := []string{"apple", "banana", "cherry"}
//	updatedWords := Pop(words)
//	// updatedWords will be []string{"apple", "banana"}
//
//	// Attempting to pop from an empty slice will cause a runtime panic
//	var emptySlice []int
//	// updatedEmpty := Pop(emptySlice) // This will panic
func Pop[T any](slice []T) []T {
	if len(slice) == 0 {
		return slice
	}
	return slice[:len(slice)-1]
}

// Unshift inserts an element at the beginning of a slice and returns the resulting slice.
//
// This function takes an input slice `slice` and an element `element`, then creates
// a new slice by appending the `element` at the start, followed by the elements of
// the original slice. The function uses the built-in `append` function to combine a
// new slice containing just the `element` with the original slice, effectively adding
// the element to the beginning.
//
// The function is generic, allowing it to work with slices of any type `T`.
//
// Parameters:
//   - `slice`: The input slice to which the element will be prepended. It can contain elements of any type `T`.
//   - `element`: The element to be inserted at the beginning of the slice. It is of type `T`.
//
// Returns:
//   - A new slice of type `[]T`, containing `element` followed by all the elements of the original `slice`.
//
// Example:
//
//	// Adding an integer to the beginning of a slice of integers
//	numbers := []int{2, 3, 4}
//	updatedNumbers := Unshift(numbers, 1)
//	// updatedNumbers will be []int{1, 2, 3, 4}
//
//	// Adding a string to the beginning of a slice of strings
//	words := []string{"banana", "cherry"}
//	updatedWords := Unshift(words, "apple")
//	// updatedWords will be []string{"apple", "banana", "cherry"}
//
//	// Adding to an empty slice
//	var emptySlice []int
//	newSlice := Unshift(emptySlice, 1)
//	// newSlice will be []int{1}
func Unshift[T any](slice []T, element T) []T {
	return append([]T{element}, slice...)
}

// Shift removes the first element from a slice and returns the resulting slice.
//
// This function takes an input slice `slice` and removes its first element by creating
// a new slice that starts from the second element of the original slice. The function
// achieves this using slicing, effectively returning a view of the original slice that
// excludes the first element. If the input slice is empty, calling this function will
// result in a runtime panic due to out-of-bounds access.
//
// The function is generic, allowing it to work with slices of any type `T`.
//
// Parameters:
//   - `slice`: The input slice from which the first element will be removed. It can contain elements of any type `T`.
//
// Returns:
//   - A new slice of type `[]T`, containing all elements from the original slice except the first one.
//
// Example:
//
//	// Removing the first element from a slice of integers
//	numbers := []int{1, 2, 3, 4}
//	updatedNumbers := Shift(numbers)
//	// updatedNumbers will be []int{2, 3, 4}
//
//	// Removing the first element from a slice of strings
//	words := []string{"apple", "banana", "cherry"}
//	updatedWords := Shift(words)
//	// updatedWords will be []string{"banana", "cherry"}
func Shift[T any](slice []T) []T {
	if len(slice) == 0 {
		return slice
	}
	return slice[1:]
}

// AppendIfMissingN appends an element to a slice if it is not already present.
//
// This function takes an input slice `slice` and an element `element`. It first checks
// if the element is already in the slice by calling the helper function `ContainsN`.
// If the element is not found in the slice, the function appends it to the end of the slice.
// If the element is already present, the original slice is returned unchanged.
//
// The function is generic and requires that the type `T` be `comparable`, allowing the
// function to use the `==` operator in `ContainsN` to check for equality.
//
// Parameters:
//   - `slice`: The input slice to which the element might be appended. It can contain elements of any comparable type `T`.
//   - `element`: The element to be appended if it is not already in `slice`. It is of type `T`.
//
// Returns:
//   - A new slice of type `[]T` containing the original elements and, if missing, the appended `element`.
//
// Example:
//
//	// Adding a missing integer to a slice
//	numbers := []int{1, 2, 3}
//	updatedNumbers := AppendIfMissingN(numbers, 4)
//	// updatedNumbers will be []int{1, 2, 3, 4}
//
//	// Trying to add an existing integer to a slice
//	updatedNumbers = AppendIfMissingN(numbers, 3)
//	// updatedNumbers will be []int{1, 2, 3} (unchanged)
//
//	// Adding a missing string to a slice
//	words := []string{"apple", "banana"}
//	updatedWords := AppendIfMissingN(words, "cherry")
//	// updatedWords will be []string{"apple", "banana", "cherry"}
func AppendIfMissingN[T comparable](slice []T, element T) []T {
	if !ContainsN(slice, element) {
		return append(slice, element)
	}
	return slice
}

// Intersect returns a new slice containing elements that are present in both input slices.
//
// This function takes two input slices, `slice1` and `slice2`, and identifies elements
// that are present in both slices. It uses a map to track the elements of `slice1`,
// then iterates over `slice2` to find common elements. If an element from `slice2` is
// found in the map (indicating it exists in `slice1`), it is added to the result slice.
//
// The function is generic, allowing it to work with slices of any `comparable` type `T`.
//
// Parameters:
//   - `slice1`: The first input slice containing elements of any comparable type `T`.
//   - `slice2`: The second input slice containing elements of any comparable type `T`.
//
// Returns:
//   - A new slice of type `[]T` that contains the elements found in both `slice1` and `slice2`.
//     Each element in the result slice will appear only once, even if it is duplicated in the input slices.
//
// Example:
//
//	// Finding common integers between two slices
//	numbers1 := []int{1, 2, 3, 4}
//	numbers2 := []int{3, 4, 5, 6}
//	commonNumbers := Intersect(numbers1, numbers2)
//	// commonNumbers will be []int{3, 4}
//
//	// Finding common strings between two slices
//	words1 := []string{"apple", "banana", "cherry"}
//	words2 := []string{"banana", "cherry", "date"}
//	commonWords := Intersect(words1, words2)
//	// commonWords will be []string{"banana", "cherry"}
//
//	// Intersecting with an empty slice results in an empty slice
//	empty := []int{}
//	intersectEmpty := Intersect(numbers1, empty)
//	// intersectEmpty will be []int{}
func Intersect[T comparable](slice1, slice2 []T) []T {
	set := make(map[T]bool)
	result := []T{}
	for _, item := range slice1 {
		set[item] = true
	}
	for _, item := range slice2 {
		if set[item] {
			result = append(result, item)
		}
	}
	return result
}

// Difference returns a new slice containing elements that are unique to each of the two input slices.
//
// This function takes two input slices, `slice1` and `slice2`, and identifies elements
// that are present in either slice but not both. It creates a map to track the elements
// of `slice1`, then checks for unique elements in `slice2` by confirming that they are
// not present in `slice1`. Finally, it appends any unique elements from `slice1` to ensure
// that the result includes all elements unique to either slice.
//
// The function is generic, allowing it to work with slices of any `comparable` type `T`.
//
// Parameters:
//   - `slice1`: The first input slice containing elements of any comparable type `T`.
//   - `slice2`: The second input slice containing elements of any comparable type `T`.
//
// Returns:
//   - A new slice of type `[]T` that contains elements unique to either `slice1` or `slice2`.
//     If an element appears in both slices, it will not appear in the result.
//
// Example:
//
//	// Finding unique integers between two slices
//	numbers1 := []int{1, 2, 3, 4}
//	numbers2 := []int{3, 4, 5, 6}
//	uniqueNumbers := Difference(numbers1, numbers2)
//	// uniqueNumbers will be []int{1, 2, 5, 6}
//
//	// Finding unique strings between two slices
//	words1 := []string{"apple", "banana", "cherry"}
//	words2 := []string{"banana", "date"}
//	uniqueWords := Difference(words1, words2)
//	// uniqueWords will be []string{"apple", "cherry", "date"}
//
//	// Difference with an empty slice results in the original slice
//	empty := []int{}
//	uniqueFromEmpty := Difference(numbers1, empty)
//	// uniqueFromEmpty will be []int{1, 2, 3, 4}
func Difference[T comparable](slice1, slice2 []T) []T {
	set := make(map[T]bool)
	result := []T{}
	for _, item := range slice1 {
		set[item] = true
	}
	for _, item := range slice2 {
		if !set[item] {
			result = append(result, item)
		}
	}
	for _, item := range slice1 {
		if !set[item] {
			result = append(result, item)
		}
	}
	return result
}

// JoinMapKeys concatenates the keys of a map into a single string, with each key separated by a specified separator.
//
// This function takes a map `m` with string keys and any type of values `V`, and a `separator` string.
// It collects all the keys of the map into a slice, then joins them into a single string using the provided separator.
//
// This function is generic, allowing it to work with maps that have values of any type `V`.
//
// Parameters:
//   - `m`: A map with string keys and values of any type `V`. Only the keys are used for concatenation.
//   - `separator`: A string used to separate each key in the resulting string.
//
// Returns:
//   - A string containing all the keys in the map `m`, separated by the specified `separator`. If the map has no keys,
//     an empty string is returned.
//
// Example:
//
//	// Concatenating the keys of a map with a comma separator
//	m := map[string]int{"apple": 1, "banana": 2, "cherry": 3}
//	joinedKeys := JoinMapKeys(m, ", ")
//	// joinedKeys will be "apple, banana, cherry"
//
//	// Using a different separator
//	m = map[string]bool{"cat": true, "dog": true}
//	joinedKeys = JoinMapKeys(m, " | ")
//	// joinedKeys will be "cat | dog"
//
//	// With an empty map
//	emptyMap := map[string]int{}
//	joinedKeys = JoinMapKeys(emptyMap, ",")
//	// joinedKeys will be ""
func JoinMapKeys[V any](m map[string]V, separator string) string {
	joined_keys := []string{}
	for key := range m {
		joined_keys = append(joined_keys, key)
	}
	return strings.Join(joined_keys, separator)
}

// DeepMergeMap merges two maps, deeply combining values from the source map into the target map.
//
// This function takes two maps: `target` and `source`, both with string keys and interface{} values. It recursively merges
// the values from the `source` map into the `target` map. If a key exists in both maps, the function checks if the values
// associated with the key are themselves maps. If so, it recursively merges the nested maps. Otherwise, it directly overwrites
// the target map's value with the value from the source map. This function allows for deep merging of nested maps.
//
// The function modifies the `target` map in place and does not return anything.
//
// Parameters:
//   - `target`: The map that will be updated with values from the `source`. It is modified in place.
//   - `source`: The map whose values will be merged into the `target`.
//
// Example:
//
//	// Merging two maps with nested maps
//	target := map[string]interface{}{
//		"fruit": map[string]interface{}{"apple": 5, "banana": 10},
//		"vegetable": map[string]interface{}{"carrot": 3},
//	}
//	source := map[string]interface{}{
//		"fruit": map[string]interface{}{"banana": 7, "orange": 2},
//		"vegetable": map[string]interface{}{"spinach": 5},
//		"grain": 100,
//	}
//	DeepMergeMap(target, source)
//	// target will now be:
//	// map[string]interface{}{
//	//		"fruit": map[string]interface{}{"apple": 5, "banana": 7, "orange": 2},
//	//		"vegetable": map[string]interface{}{"carrot": 3, "spinach": 5},
//	//		"grain": 100,
//	//	}
//
//	// If there is no conflict, the value from the source is added as is.
//	// If the source value is a nested map, the function will perform a deep merge.
func DeepMergeMap(target, source map[string]interface{}) {
	for key, sourceValue := range source {
		if targetValue, exists := target[key]; exists {
			if sourceMap, sourceIsMap := sourceValue.(map[string]interface{}); sourceIsMap {
				if targetMap, targetIsMap := targetValue.(map[string]interface{}); targetIsMap {
					DeepMergeMap(targetMap, sourceMap)
				}
			} else {
				target[key] = sourceValue
			}
		} else {
			target[key] = sourceValue
		}
	}
}

// MergeMapString merges multiple maps of type map[string]string into a single map.
//
// This function takes a variadic number of maps of type map[string]string and combines their key-value pairs into a
// single resulting map. If there are duplicate keys across the input maps, the value from the last map in the
// variadic list will overwrite the earlier ones.
//
// The function creates a new map to hold the merged results, iterating over each input map and adding its
// key-value pairs to the result map.
//
// Parameters:
//   - `maps`: A variadic parameter that takes one or more maps of type map[string]string to be merged.
//
// Returns:
//   - A new map of type map[string]string that contains all the key-value pairs from the input maps.
//     If a key appears in multiple maps, the value from the last map will be used.
//
// Example:
//
//	// Merging two maps
//	map1 := map[string]string{"a": "apple", "b": "banana"}
//	map2 := map[string]string{"b": "blueberry", "c": "cherry"}
//	merged := MergeMapString(map1, map2)
//	// merged will be map[string]string{"a": "apple", "b": "blueberry", "c": "cherry"}
//
//	// Merging more than two maps
//	map3 := map[string]string{"d": "date"}
//	mergedAll := MergeMapString(map1, map2, map3)
//	// mergedAll will be map[string]string{"a": "apple", "b": "blueberry", "c": "cherry", "d": "date"}
//
//	// If duplicate keys exist, the value from the last map wins
//	map4 := map[string]string{"a": "apricot"}
//	mergedWithDup := MergeMapString(map1, map4)
//	// mergedWithDup will be map[string]string{"a": "apricot", "b": "banana"}
func MergeMapString(maps ...map[string]string) map[string]string {
	result := make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// MapString2Tb formats a map of type map[string]string into a string representation suitable for a table.
//
// This function takes a map with string keys and string values, and formats it into a neatly aligned
// string where each key and value are presented in two columns. The key column is aligned to the left
// with the longest key in the map, and the values are displayed next to their corresponding keys. The
// function uses a `strings.Builder` for efficient string concatenation.
//
// The function computes the maximum key length to ensure that all keys align properly. Then, it iterates
// over each key-value pair in the map, appending them in a formatted manner to the builder.
//
// Parameters:
//   - `data`: A map of type `map[string]string` that contains key-value pairs to be formatted.
//
// Returns:
//   - A string representing the map in a table-like format with keys and values aligned in columns.
//
// Example:
//
//	// Formatting a map into a table-like string
//	data := map[string]string{
//		"apple":  "fruit",
//		"carrot": "vegetable",
//		"banana": "fruit",
//	}
//	formattedTable := MapString2Tb(data)
//	// formattedTable will be:
//	// apple   fruit
//	// carrot  vegetable
//	// banana  fruit
//
//	// If the keys have different lengths, the function ensures that all values are aligned properly.
//	// The longest key will define the column width for the keys.
func MapString2Tb(data map[string]string) string {
	var builder strings.Builder
	maxKeyLength := 0
	for key := range data {
		if len(key) > maxKeyLength {
			maxKeyLength = len(key)
		}
	}
	for key, value := range data {
		fmt.Fprintf(&builder, "%-*s   %s\n", maxKeyLength, key, value)
	}
	return builder.String()
}

// Map2Table formats a map of type map[string]interface{} into a string representation suitable for a table.
//
// This function takes a map with string keys and values of any type (interface{}), and formats it into a neatly
// aligned string where each key and value are presented in two columns. The key column is aligned to the left
// with the longest key in the map, and the corresponding values are displayed next to their keys. The function
// uses a `strings.Builder` for efficient string concatenation.
//
// The function computes the maximum key length to ensure that all keys align properly. It then iterates over each
// key-value pair in the map, formatting the value by serializing it into a JSON string using a helper function `Json()`.
// This ensures that even complex values (such as structs or maps) are represented in a readable format.
//
// Parameters:
//   - `data`: A map of type `map[string]interface{}` that contains key-value pairs to be formatted. The values can
//     be of any type, and they will be serialized to JSON for proper display.
//
// Returns:
//   - A string representing the map in a table-like format with keys and their serialized values aligned in columns.
//
// Example:
//
//	// Formatting a map with mixed value types into a table-like string
//	data := map[string]interface{}{
//		"fruit":   "apple",
//		"vegetable": map[string]interface{}{"type": "root", "name": "carrot"},
//		"quantity": 10,
//	}
//	formattedTable := Map2Table(data)
//	// formattedTable will be:
//	// fruit       apple
//	// vegetable   {"type":"root","name":"carrot"}
//	// quantity    10
//
//	// The Json function will serialize the `vegetable` map, ensuring that it is displayed in JSON format.
func Map2Table(data map[string]interface{}) string {
	var builder strings.Builder
	maxKeyLength := 0
	for key := range data {
		if len(key) > maxKeyLength {
			maxKeyLength = len(key)
		}
	}
	for key, value := range data {
		fmt.Fprintf(&builder, "%-*s   %s\n", maxKeyLength, key, Json(value))
	}
	return builder.String()
}

// IndexExists checks whether a given index is valid for the provided slice.
//
// This function takes a slice `slice` of any type `T` and an integer `index`, and returns a boolean indicating
// whether the specified index is within the valid range for the slice. A valid index is one that is greater than
// or equal to 0 and less than the length of the slice.
//
// Parameters:
//   - `slice`: The input slice of any type `T` that is being checked.
//   - `index`: The index to check for validity in the slice.
//
// Returns:
//   - `true` if the index is within the bounds of the slice (i.e., 0 <= index < len(slice)).
//   - `false` otherwise, such as when the index is negative or greater than or equal to the length of the slice.
//
// Example:
//
//	// Checking if an index exists in a slice
//	numbers := []int{1, 2, 3, 4}
//	exists := IndexExists(numbers, 2)
//	// exists will be true, as index 2 is valid for the slice
//
//	// Checking an invalid index
//	exists = IndexExists(numbers, 5)
//	// exists will be false, as index 5 is out of bounds for the slice
func IndexExists[T any](slice []T, index int) bool {
	return index >= 0 && index < len(slice)
}

// Iterate iterates over a collection (slice, array, or map) and applies a callback function on each element.
//
// This function takes a collection of any type (using an empty `interface{}`), which can be a slice, array, or map,
// and a callback function. The callback function is executed for each element in the collection. For slices and arrays,
// the callback receives the index and the corresponding value. For maps, the callback receives each key and value, with
// the key being passed first followed by the value. For slices and arrays, the index is passed, while for maps, it is
// passed as -1 (since maps are unordered).
//
// The function uses reflection to handle different types of collections and ensures that the correct value is passed
// to the callback.
//
// Parameters:
//   - `collection`: The collection (slice, array, or map) to iterate over. It can be of any type.
//   - `callback`: A function that takes two arguments: the index (for slices and arrays, -1 for maps) and the value
//     from the collection. The callback is executed for each element in the collection.
//
// Example:
//
//	// Iterating over a slice
//	numbers := []int{1, 2, 3, 4}
//	Iterate(numbers, func(index int, value interface{}) {
//		fmt.Printf("Index: %d, Value: %v\n", index, value)
//	})
//	// Output:
//	// Index: 0, Value: 1
//	// Index: 1, Value: 2
//	// Index: 2, Value: 3
//	// Index: 3, Value: 4
//
//	// Iterating over a map
//	colors := map[string]string{"red": "FF0000", "green": "00FF00", "blue": "0000FF"}
//	Iterate(colors, func(index int, value interface{}) {
//		fmt.Printf("Value: %v\n", value)
//	})
//	// Output:
//	// Value: red
//	// Value: FF0000
//	// Value: green
//	// Value: 00FF00
//	// Value: blue
//	// Value: 0000FF
//
// Notes:
//   - For slices and arrays, the callback will receive the index and the value from the collection.
//   - For maps, the callback will be executed twice per key-value pair: once with the key and once with the value,
//     since maps are unordered and the order of key-value pairs cannot be guaranteed.
func Iterate(collection interface{}, callback func(index int, value interface{})) {
	v := reflect.ValueOf(collection)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			callback(i, v.Index(i).Interface())
		}
	} else if v.Kind() == reflect.Map {
		keys := v.MapKeys()
		for _, key := range keys {
			callback(-1, key.Interface())
			callback(-1, v.MapIndex(key).Interface())
		}
	}
}

// MapN applies a transformation function to each element of a collection (slice, array, or map) and returns
// a new collection with the transformed values.
//
// This function takes a collection of any type (using an empty `interface{}`), which can be a slice, array, or map,
// and a mapping function (`mapper`). The function applies the `mapper` function to each value in the collection,
// transforming it. For slices and arrays, the callback is applied to each element. For maps, the callback is applied
// to each key and value. The transformed elements (or key-value pairs) are collected into a new result, which is returned.
//
// The function uses reflection to handle different types of collections and constructs a new collection with the transformed values.
//
// Parameters:
//   - `collection`: The collection (slice, array, or map) to iterate over and transform. It can be of any type.
//   - `mapper`: A function that takes a value from the collection and transforms it. The function is applied to each
//     element of the collection (or key-value pair for maps).
//
// Returns:
//   - A new collection of the same type as the original collection, where each element has been transformed using the
//     provided `mapper` function.
//
// Example:
//
//	// Mapping a slice of integers to their squares
//	numbers := []int{1, 2, 3, 4}
//	squared := MapN(numbers, func(value interface{}) interface{} {
//		return value.(int) * value.(int)
//	})
//	// squared will be []int{1, 4, 9, 16}
//
//	// Mapping a map of strings to their lengths
//	words := map[string]string{"apple": "fruit", "carrot": "vegetable"}
//	lengths := MapN(words, func(value interface{}) interface{} {
//		return len(value.(string))
//	})
//	// lengths will be a new collection containing key-value pairs, where values represent string lengths.
//
// Notes:
//   - For slices and arrays, the `mapper` function is applied to each element.
//   - For maps, the `mapper` function is applied to both the key and the value. The resulting collection will include
//     both transformed keys and values.
//
// Limitations:
//   - The function creates a new collection based on the results of the `mapper` function, so it does not modify the
//     original collection.
func MapN(collection interface{}, mapper func(value interface{}) interface{}) interface{} {
	v := reflect.ValueOf(collection)
	result := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(mapper(v.Index(0).Interface()))), 0, 0)

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			mappedValue := mapper(v.Index(i).Interface())
			result = reflect.Append(result, reflect.ValueOf(mappedValue))
		}
	} else if v.Kind() == reflect.Map {
		keys := v.MapKeys()
		for _, key := range keys {
			mappedKey := mapper(key.Interface())
			mappedValue := mapper(v.MapIndex(key).Interface())
			result = reflect.Append(result, reflect.ValueOf(mappedKey))
			result = reflect.Append(result, reflect.ValueOf(mappedValue))
		}
	}
	return result.Interface()
}

// FilterN filters a collection (slice or array) based on a predicate function and returns a new collection
// containing only the elements that satisfy the condition specified by the predicate.
//
// This function takes a collection of any type (using an empty `interface{}`), which can be a slice or array,
// and a filtering predicate function (`predicate`). The function applies the `predicate` function to each element in
// the collection, and if the predicate returns true, the element is included in the new collection. The function only
// supports slices and arrays as input collections.
//
// The function uses reflection to handle slices and arrays and constructs a new collection containing only the
// elements that pass the filter condition.
//
// Parameters:
//   - `collection`: The collection (slice or array) to filter. It can be of any type, but only slices and arrays
//     are supported.
//   - `predicate`: A function that takes a value from the collection and returns a boolean indicating whether the
//     element should be included in the resulting collection. If it returns true, the element is included.
//
// Returns:
//   - A new collection of the same type as the original collection, containing only the elements that satisfy the
//     condition defined by the `predicate` function.
//
// Example:
//
//	// Filtering a slice of integers to get only even numbers
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	evens := FilterN(numbers, func(value interface{}) bool {
//		return value.(int)%2 == 0
//	})
//	// evens will be []int{2, 4, 6}
//
// Notes:
//   - This function only works with slices or arrays, and will return an empty collection if the input is of another type.
//
// Limitations:
//   - The function creates a new collection based on the results of the `predicate` function, so it does not modify
//     the original collection.
func FilterN(collection interface{}, predicate func(value interface{}) bool) interface{} {
	v := reflect.ValueOf(collection)
	result := reflect.MakeSlice(v.Type(), 0, 0)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i).Interface()
			if predicate(item) {
				result = reflect.Append(result, reflect.ValueOf(item))
			}
		}
	}
	return result.Interface()
}

// ReduceN reduces a collection (slice or array) to a single value by applying a reducer function
// to each element, combining them into a single result.
//
// This function takes a collection of any type (using an empty `interface{}`), which can be a slice or array,
// and a reducer function (`reducer`). The reducer function is applied to each element of the collection, along with an
// accumulator value, and the result of the reducer is passed as the accumulator to the next iteration. This process
// continues until all elements are processed, resulting in a single accumulated value.
//
// The function uses reflection to handle slices and arrays and iterates over the collection to apply the reducer function.
//
// Parameters:
//   - `collection`: The collection (slice or array) to reduce. It can be of any type, but only slices and arrays
//     are supported.
//   - `reducer`: A function that takes two arguments: an accumulator and an element from the collection. It returns
//     a new accumulator value after combining the accumulator and the element.
//   - `initialValue`: The initial value of the accumulator, used as the starting point for the reduction process.
//
// Returns:
//   - A single value, which is the result of reducing the entire collection using the `reducer` function. The return
//     type is the same as the type of the `initialValue`.
//
// Example:
//
//	// Reducing a slice of integers by summing them
//	numbers := []int{1, 2, 3, 4, 5}
//	sum := ReduceN(numbers, func(acc interface{}, value interface{}) interface{} {
//		return acc.(int) + value.(int)
//	}, 0)
//	// sum will be 15
//
// Notes:
//   - This function only works with slices or arrays, and will return the initial value if the input collection is empty.
//
// Limitations:
//   - The function creates a single accumulated result by repeatedly applying the `reducer` function to each element,
//     so it does not modify the original collection.
func ReduceN(collection interface{}, reducer func(acc interface{}, value interface{}) interface{}, initialValue interface{}) interface{} {
	v := reflect.ValueOf(collection)
	accumulator := initialValue
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			accumulator = reducer(accumulator, v.Index(i).Interface())
		}
	}
	return accumulator
}

// Find searches for the first element in a collection (slice or array) that satisfies a given predicate
// function and returns it.
//
// This function takes a collection of any type (using an empty `interface{}`), which can be a slice or array,
// and a predicate function (`predicate`). The predicate function is applied to each element of the collection,
// and if it returns true for an element, that element is returned as the result. The function will return the first
// element that satisfies the condition, and if no elements satisfy the condition, it returns `nil`.
//
// The function uses reflection to handle slices and arrays, iterating over the collection to check each element
// with the provided predicate.
//
// Parameters:
//   - `collection`: The collection (slice or array) to search through. It can be of any type, but only slices and arrays
//     are supported.
//   - `predicate`: A function that takes a value from the collection and returns a boolean indicating whether the element
//     satisfies the condition. If it returns true, the element is returned.
//
// Returns:
//   - The first element from the collection that satisfies the `predicate` function. If no elements satisfy the condition,
//     it returns `nil`.
//
// Example:
//
//	// Finding the first even number in a slice of integers
//	numbers := []int{1, 2, 3, 4, 5}
//	even := Find(numbers, func(value interface{}) bool {
//		return value.(int)%2 == 0
//	})
//	// even will be 2 (the first even number)
//
// Notes:
//   - This function only works with slices or arrays. If the collection is of another type, it will return `nil`.
//   - The function returns the first element that matches the predicate and stops searching after finding it.
//
// Limitations:
//   - The function works only with slices and arrays. If no elements satisfy the predicate, the function will return `nil`,
//     even if the collection is non-empty.
func Find(collection interface{}, predicate func(value interface{}) bool) interface{} {
	v := reflect.ValueOf(collection)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i).Interface()
			if predicate(item) {
				return item
			}
		}
	}
	return nil
}

// All checks whether all elements in a collection (slice or array) satisfy a given condition.
//
// This function takes a collection of any type (using an empty `interface{}`), which can be a slice or array,
// and a condition function (`condition`). The condition function is applied to each element of the collection, and
// the function returns `true` if all elements satisfy the condition. If any element does not satisfy the condition,
// the function returns `false`. If the collection is empty, the function returns `true` (since the condition is trivially satisfied).
//
// The function uses reflection to handle slices and arrays, iterating over the collection to check each element
// with the provided condition.
//
// Parameters:
//   - `collection`: The collection (slice or array) to check. It can be of any type, but only slices and arrays
//     are supported.
//   - `condition`: A function that takes a value from the collection and returns a boolean indicating whether the
//     element satisfies the condition. If it returns `false` for any element, the function immediately returns `false`.
//
// Returns:
//   - `true` if all elements in the collection satisfy the condition, otherwise `false`.
//   - If the collection is empty, the function returns `true`.
//
// Example:
//
//	// Checking if all elements in a slice of integers are positive
//	numbers := []int{1, 2, 3, 4, 5}
//	allPositive := All(numbers, func(value interface{}) bool {
//		return value.(int) > 0
//	})
//	// allPositive will be true
//
// Notes:
//   - This function only works with slices or arrays. If the collection is of another type, it will return `false`.
//   - The function returns `false` as soon as it finds an element that does not satisfy the condition, making it more
//     efficient for early termination.
//
// Limitations:
//   - The function works only with slices and arrays. If no elements satisfy the condition, the function will return `false`,
//     but if all elements are valid, it will return `true`. An empty collection is considered to trivially satisfy the condition.
func All(collection interface{}, condition func(value interface{}) bool) bool {
	v := reflect.ValueOf(collection)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			if !condition(v.Index(i).Interface()) {
				return false
			}
		}
		return true
	}
	return false
}

// Any checks whether any element in a collection (slice or array) satisfies a given condition.
//
// This function accepts a collection (slice or array) and a condition function. It returns `true` if at least one element in the collection satisfies the condition,
// and `false` if none do. The function stops iterating as soon as a matching element is found, making it more efficient for early termination.
//
// Parameters:
//   - `collection`: A slice or array of any type to check.
//   - `condition`: A function that takes a value from the collection and returns a boolean indicating whether the element satisfies the condition.
//
// Returns:
//   - `true` if any element satisfies the condition, `false` otherwise.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	anyNegative := Any(numbers, func(value interface{}) bool {
//	  return value.(int) < 0
//	})
//	// anyNegative will be false because no element is negative.
func Any(collection interface{}, condition func(value interface{}) bool) bool {
	v := reflect.ValueOf(collection)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			if condition(v.Index(i).Interface()) {
				return true
			}
		}
		return false
	}
	return false
}

// Count returns the number of elements in a collection (slice or array) that satisfy a given condition.
//
// This function takes a collection (slice or array) and a condition function. It iterates through the collection, applying the condition to each element.
// It returns the total count of elements that satisfy the condition.
//
// Parameters:
//   - `collection`: A slice or array of any type to check.
//   - `condition`: A function that checks if an element satisfies a condition. The function returns `true` for elements that match the condition, and `false` otherwise.
//
// Returns:
//   - The count of elements in the collection that satisfy the condition.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	countNegative := Count(numbers, func(value interface{}) bool {
//	  return value.(int) < 0
//	})
//	// countNegative will be 0, since no element is negative.
func Count(collection interface{}, condition func(value interface{}) bool) int {
	v := reflect.ValueOf(collection)
	count := 0
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			if condition(v.Index(i).Interface()) {
				count++
			}
		}
	}
	return count
}

// RemoveN returns a new collection (slice or array) where all elements that satisfy a given condition are removed.
//
// This function takes a collection (slice or array) and a condition function. It iterates through the collection, and for each element that does not satisfy the condition,
// it is added to a new result collection. The function returns a new collection that contains only the elements that do not match the condition.
//
// Parameters:
//   - `collection`: A slice or array of any type to process.
//   - `condition`: A function that checks if an element satisfies a condition. If the element matches the condition, it is removed from the result collection.
//
// Returns:
//   - A new collection (slice or array) with elements that do not satisfy the condition.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	result := RemoveN(numbers, func(value interface{}) bool {
//	  return value.(int) % 2 == 0 // Removes even numbers
//	})
//	// result will be []int{1, 3, 5}
func RemoveN(collection interface{}, condition func(value interface{}) bool) interface{} {
	v := reflect.ValueOf(collection)
	result := reflect.MakeSlice(v.Type(), 0, 0)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i).Interface()
			if !condition(item) {
				result = reflect.Append(result, reflect.ValueOf(item))
			}
		}
	}
	return result.Interface()
}

// SortN sorts a collection (slice or array) in-place according to a custom comparison function.
//
// This function takes a collection (slice or array) and a `less` function. The `less` function should return `true` if the element at index `i` should come before the element at index `j`.
// The collection is sorted in-place, meaning the original collection is modified.
//
// Parameters:
//   - `collection`: A slice or array of any type to sort.
//   - `less`: A comparison function that takes two indices `i` and `j` and returns `true` if the element at index `i` should come before the element at index `j`.
//
// Returns:
//   - The collection is sorted in place. No value is returned.
//
// Example:
//
//	numbers := []int{5, 3, 1, 4, 2}
//	SortN(numbers, func(i, j int) bool {
//	  return numbers[i] < numbers[j] // Sort in ascending order
//	})
//	// numbers will be sorted to []int{1, 2, 3, 4, 5}
func SortN(collection interface{}, less func(i, j int) bool) {
	v := reflect.ValueOf(collection)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		sort.SliceStable(collection, func(i, j int) bool {
			return less(i, j)
		})
	}
}

// Reverse_N reverses the order of elements in a collection (slice or array) in-place.
//
// This function takes a collection (slice or array) and reverses the elements by swapping elements at corresponding positions (i and length-i-1) until the middle of the collection is reached.
// The collection is modified in place, meaning no new collection is created.
//
// Parameters:
//   - `collection`: A slice or array of any type to reverse.
//
// Returns:
//   - The collection is reversed in place. No value is returned.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	Reverse_N(numbers)
//	// numbers will be reversed to []int{5, 4, 3, 2, 1}
func Reverse_N(collection interface{}) {
	v := reflect.ValueOf(collection)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		length := v.Len()
		for i := 0; i < length/2; i++ {
			j := length - i - 1
			vi := v.Index(i).Interface()
			vj := v.Index(j).Interface()
			v.Index(i).Set(reflect.ValueOf(vj))
			v.Index(j).Set(reflect.ValueOf(vi))
		}
	}
}

// UniqueN returns a new collection (slice or array) containing only unique elements from the original collection.
//
// This function takes a collection (slice or array) and removes duplicate elements, returning a new collection with only the unique elements.
// The function uses a map to track elements that have already been encountered, ensuring that only the first occurrence of each element is included in the result.
//
// Parameters:
//   - `collection`: A slice or array of any type from which duplicates should be removed.
//
// Returns:
//   - A new collection (slice or array) containing only the unique elements from the original collection.
//
// Example:
//
//	numbers := []int{1, 2, 2, 3, 4, 4, 5}
//	result := UniqueN(numbers)
//	// result will be []int{1, 2, 3, 4, 5}
func UniqueN(collection interface{}) interface{} {
	v := reflect.ValueOf(collection)
	uniqueMap := make(map[interface{}]struct{})
	result := reflect.MakeSlice(v.Type(), 0, 0)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i).Interface()
			if _, found := uniqueMap[item]; !found {
				uniqueMap[item] = struct{}{}
				result = reflect.Append(result, reflect.ValueOf(item))
			}
		}
	}
	return result.Interface()
}

// Contains_N checks if a given element exists within a collection (slice or array).
//
// This function takes a collection (slice or array) and an element, then iterates through the collection to see if any element matches the provided element.
// It uses `reflect.DeepEqual` to compare elements, which ensures that even complex types (like structs or slices) are compared correctly.
//
// Parameters:
//   - `collection`: A slice or array of any type to search within.
//   - `element`: The element to search for within the collection.
//
// Returns:
//   - `true` if the element is found within the collection, `false` otherwise.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	containsThree := Contains_N(numbers, 3)
//	// containsThree will be true because 3 is in the slice.
func Contains_N(collection interface{}, element interface{}) bool {
	v := reflect.ValueOf(collection)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			if reflect.DeepEqual(v.Index(i).Interface(), element) {
				return true
			}
		}
	}
	return false
}

// Difference_N returns a new collection (slice or array) containing elements from the first collection
// that are not present in the second collection.
//
// This function takes two collections (slices or arrays) and compares the elements of the first collection
// against the elements of the second collection. It returns a new collection with elements that appear in the
// first collection but are absent in the second collection. The function uses `Contains_N` to check for membership
// of each element of the first collection in the second collection.
//
// Parameters:
//   - `collection1`: The first slice or array of any type to compare.
//   - `collection2`: The second slice or array of any type to compare against.
//
// Returns:
//   - A new collection (slice or array) containing the elements from `collection1` that are not in `collection2`.
//
// Example:
//
//	numbers1 := []int{1, 2, 3, 4, 5}
//	numbers2 := []int{3, 4, 6}
//	result := Difference_N(numbers1, numbers2)
//	// result will be []int{1, 2, 5}, as these are the elements in numbers1 that are not in numbers2.
func Difference_N(collection1 interface{}, collection2 interface{}) interface{} {
	v1 := reflect.ValueOf(collection1)
	result := reflect.MakeSlice(v1.Type(), 0, 0)
	if v1.Kind() == reflect.Slice || v1.Kind() == reflect.Array {
		for i := 0; i < v1.Len(); i++ {
			item := v1.Index(i).Interface()
			if !Contains_N(collection2, item) {
				result = reflect.Append(result, v1.Index(i))
			}
		}
	}
	return result.Interface()
}

// Intersection returns a new collection (slice or array) containing elements that are present in both collections.
//
// This function takes two collections (slices or arrays) and compares the elements of both collections.
// It returns a new collection with elements that appear in both the first and the second collection.
// The function uses `Contains_N` to check if an element from the first collection is also present in the second collection.
//
// Parameters:
//   - `collection1`: The first slice or array of any type to compare.
//   - `collection2`: The second slice or array of any type to compare against.
//
// Returns:
//   - A new collection (slice or array) containing the elements that are found in both `collection1` and `collection2`.
//
// Example:
//
//	numbers1 := []int{1, 2, 3, 4, 5}
//	numbers2 := []int{3, 4, 6}
//	result := Intersection(numbers1, numbers2)
//	// result will be []int{3, 4}, as these are the elements common to both numbers1 and numbers2.
func Intersection(collection1 interface{}, collection2 interface{}) interface{} {
	v1 := reflect.ValueOf(collection1)
	result := reflect.MakeSlice(v1.Type(), 0, 0)
	if v1.Kind() == reflect.Slice || v1.Kind() == reflect.Array {
		for i := 0; i < v1.Len(); i++ {
			item := v1.Index(i).Interface()
			if Contains_N(collection2, item) {
				result = reflect.Append(result, v1.Index(i))
			}
		}
	}
	return result.Interface()
}

// Slice returns a new collection (slice or array) that is a subrange of the input collection,
// starting from the specified `start` index and ending at the `end` index (exclusive).
//
// This function creates a sub-slice or sub-array from the given collection by copying elements from the
// original collection starting at the `start` index and ending just before the `end` index. If `start` or `end`
// is out of bounds, the function adjusts them to ensure they stay within the valid range for the collection.
// The returned collection will contain the elements between the adjusted `start` and `end` indices.
//
// Parameters:
//   - `collection`: A slice or array of any type to extract a subrange from.
//   - `start`: The starting index (inclusive) for the sub-slice or sub-array.
//   - `end`: The ending index (exclusive) for the sub-slice or sub-array.
//
// Returns:
//   - A new collection (slice or array) containing the elements from `start` to `end` (exclusive).
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5, 6, 7}
//	result := Slice(numbers, 2, 5)
//	// result will be []int{3, 4, 5}, as it extracts elements from index 2 to 4 (inclusive of 2, exclusive of 5).
//
// Notes:
//   - If `start` is less than 0, it is adjusted to 0 (beginning of the collection).
//   - If `end` is greater than the length of the collection, it is adjusted to the collection's length.
func Slice(collection interface{}, start, end int) interface{} {
	v := reflect.ValueOf(collection)
	result := reflect.MakeSlice(v.Type(), 0, 0)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		if start < 0 {
			start = 0
		}
		if end > v.Len() {
			end = v.Len()
		}
		for i := start; i < end; i++ {
			result = reflect.Append(result, v.Index(i))
		}
	}
	return result.Interface()
}

// SliceWithIndices returns a new collection (slice or array) containing elements from the input collection
// specified by the provided indices.
//
// This function creates a new collection by selecting elements from the original collection based on a list
// of indices provided in the `indices` slice. It ensures that only valid indices (within the bounds of the
// collection) are used to build the result collection.
//
// Parameters:
//   - `collection`: A slice or array of any type to extract elements from.
//   - `indices`: A slice of integers representing the indices of elements to include in the result collection.
//     Only valid indices (within the bounds of the collection) are considered.
//
// Returns:
//   - A new collection (slice or array) containing the elements from `collection` at the specified `indices`.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5, 6, 7}
//	indices := []int{1, 3, 5}
//	result := SliceWithIndices(numbers, indices)
//	// result will be []int{2, 4, 6}, as these are the elements at indices 1, 3, and 5 of the original slice.
//
// Notes:
//   - If an index in `indices` is out of bounds (less than 0 or greater than or equal to the collection's length),
//     it is ignored.
func SliceWithIndices(collection interface{}, indices []int) interface{} {
	v := reflect.ValueOf(collection)
	result := reflect.MakeSlice(v.Type(), 0, 0)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for _, index := range indices {
			if index >= 0 && index < v.Len() {
				result = reflect.Append(result, v.Index(index))
			}
		}
	}
	return result.Interface()
}

// Partition splits a collection (slice or array) into two parts based on a condition function.
//
// This function iterates through the elements of the input collection, applying the provided `condition` function to each element.
// It creates two separate collections: one containing elements for which the condition returns `true`, and the other containing elements
// for which the condition returns `false`. The function returns both collections.
//
// Parameters:
//   - `collection`: A slice or array of any type to partition.
//   - `condition`: A function that checks whether an element should go into the "true" partition or the "false" partition.
//     It returns `true` if the element should go into the "true" partition, and `false` otherwise.
//
// Returns:
//   - A tuple containing two new collections:
//   - The first collection contains elements for which the condition is `true`.
//   - The second collection contains elements for which the condition is `false`.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	truePartition, falsePartition := Partition(numbers, func(value interface{}) bool {
//		return value.(int) % 2 == 0 // Partition into even and odd numbers
//	})
//	// truePartition will be []int{2, 4, 6} (even numbers)
//	// falsePartition will be []int{1, 3, 5} (odd numbers)
//
// Notes:
//   - The condition function is applied to each element in the collection.
//   - The returned collections are of the same type as the input collection (slice or array).
func Partition(collection interface{}, condition func(value interface{}) bool) (interface{}, interface{}) {
	v := reflect.ValueOf(collection)
	truePartition := reflect.MakeSlice(v.Type(), 0, 0)
	falsePartition := reflect.MakeSlice(v.Type(), 0, 0)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i).Interface()
			if condition(item) {
				truePartition = reflect.Append(truePartition, reflect.ValueOf(item))
			} else {
				falsePartition = reflect.Append(falsePartition, reflect.ValueOf(item))
			}
		}
	}
	return truePartition.Interface(), falsePartition.Interface()
}

// Zip combines multiple collections (slices or arrays) element-wise into a new collection of tuples.
//
// This function takes multiple collections (slices or arrays) as arguments and combines their elements
// into tuples. Each tuple consists of elements from the same index of each collection. The function returns
// a new collection (a slice of tuples), where each tuple contains the elements from the input collections
// at the same position. The length of the resulting collection is determined by the shortest input collection.
//
// Parameters:
//   - `collections`: A variadic list of slices or arrays of any type to combine. All collections must be
//     of the same length or shorter, and they will be combined element-wise until the shortest collection's length is reached.
//
// Returns:
//   - A slice of tuples (slices), where each tuple contains the elements from the input collections at the same index.
//     The length of the returned slice is the same as the shortest input collection.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	strings := []string{"a", "b", "c"}
//	result := Zip(numbers, strings)
//	// result will be [][]interface{}{{1, "a"}, {2, "b"}, {3, "c"}}
//
// Notes:
//   - If any of the input collections is not a slice or array, the function returns `nil`.
//   - If the collections have different lengths, the function will combine elements up to the length of the shortest collection.
func Zip(collections ...interface{}) []interface{} {
	minLength := -1
	for _, collection := range collections {
		v := reflect.ValueOf(collection)
		if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
			return nil
		}
		if minLength == -1 || v.Len() < minLength {
			minLength = v.Len()
		}
	}
	result := make([]interface{}, minLength)
	for i := 0; i < minLength; i++ {
		tuple := make([]interface{}, len(collections))
		for j, collection := range collections {
			v := reflect.ValueOf(collection)
			tuple[j] = v.Index(i).Interface()
		}
		result[i] = tuple
	}
	return result
}

// ReduceRight performs a right-to-left reduction on a collection (slice or array) using a reducer function.
//
// This function takes a collection (slice or array), a reducer function, and an initial accumulator value.
// It iterates through the collection from right to left, applying the reducer function to each element and
// accumulating the result. The reducer function is called with two arguments: the current accumulator value
// and the current element in the collection. The final result is returned after processing all elements.
//
// Parameters:
//   - `collection`: A slice or array of any type to reduce.
//   - `reducer`: A function that takes the current accumulator value and the current element as input, and returns
//     the updated accumulator value. This function is applied from right to left on the collection.
//   - `initialValue`: The initial value for the accumulator, which will be passed as the first argument to the
//     reducer function during the first iteration.
//
// Returns:
//   - The final accumulated value after reducing the collection from right to left.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4}
//	result := ReduceRight(numbers, func(acc, value interface{}) interface{} {
//		return acc.(int) + value.(int) // Sum of elements from right to left
//	}, 0)
//	// result will be 10, as the reduction is (0 + 4) + (4 + 3) + (7 + 2) + (9 + 1) = 10
//
// Notes:
//   - The reduction starts from the rightmost element of the collection and proceeds towards the left.
//   - The function uses reflection to support collections of any type.
func ReduceRight(collection interface{}, reducer func(acc, value interface{}) interface{}, initialValue interface{}) interface{} {
	v := reflect.ValueOf(collection)
	accumulator := initialValue
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := v.Len() - 1; i >= 0; i-- {
			accumulator = reducer(accumulator, v.Index(i).Interface())
		}
	}
	return accumulator
}

// RotateLeft rotates the elements of a collection (slice or array) to the left by a specified number of positions.
//
// This function takes a collection (slice or array) and rotates its elements to the left by the given number
// of positions. Elements at the beginning of the collection are moved to the end. The rotation is done in-place
// (modifying the collection), and if the number of positions is negative, the rotation will be adjusted accordingly.
//
// Parameters:
//   - `collection`: A slice or array of any type to rotate.
//   - `positions`: The number of positions to rotate the collection to the left. A positive number rotates
//     elements left, while a negative number rotates right. The positions are normalized to be within the valid range
//     of the collection's length.
//
// Returns:
//   - A new collection (slice or array) where the elements have been rotated left by the specified number of positions.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	result := RotateLeft(numbers, 2)
//	// result will be []int{3, 4, 5, 1, 2}, as the collection is rotated left by 2 positions.
//
// Notes:
//   - If the number of positions is larger than the length of the collection, it is normalized using modulo
//     to ensure it rotates only the necessary number of positions.
//   - If the collection is not a slice or array, the original collection is returned unchanged.
func RotateLeft(collection interface{}, positions int) interface{} {
	v := reflect.ValueOf(collection)
	length := v.Len()
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		if positions < 0 {
			positions = (positions%length + length) % length
		} else {
			positions = positions % length
		}
		result := reflect.MakeSlice(v.Type(), length, length)
		for i := 0; i < length; i++ {
			result.Index((i - positions + length) % length).Set(v.Index(i))
		}
		return result.Interface()
	}
	return collection
}

// RotateRight rotates the elements of a collection (slice or array) to the right by a specified number of positions.
//
// This function takes a collection (slice or array) and rotates its elements to the right by the given number
// of positions. Elements at the end of the collection are moved to the beginning. The rotation is done in-place
// (modifying the collection), and if the number of positions is negative, the rotation will be adjusted accordingly.
//
// Parameters:
//   - `collection`: A slice or array of any type to rotate.
//   - `positions`: The number of positions to rotate the collection to the right. A positive number rotates
//     elements right, while a negative number rotates left. The positions are normalized to be within the valid range
//     of the collection's length.
//
// Returns:
//   - A new collection (slice or array) where the elements have been rotated right by the specified number of positions.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	result := RotateRight(numbers, 2)
//	// result will be []int{4, 5, 1, 2, 3}, as the collection is rotated right by 2 positions.
//
// Notes:
//   - If the number of positions is larger than the length of the collection, it is normalized using modulo
//     to ensure it rotates only the necessary number of positions.
//   - If the collection is not a slice or array, the original collection is returned unchanged.
func RotateRight(collection interface{}, positions int) interface{} {
	v := reflect.ValueOf(collection)
	length := v.Len()
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		if positions < 0 {
			positions = (-positions%length + length) % length
		} else {
			positions = positions % length
		}
		result := reflect.MakeSlice(v.Type(), length, length)
		for i := 0; i < length; i++ {
			result.Index((i + positions) % length).Set(v.Index(i))
		}
		return result.Interface()
	}
	return collection
}
