package unify4go

import (
	"math/rand"
	"time"
)

var r *rand.Rand // Package-level random generator

func init() {
	// Initialize the package-level random generator with a seed
	src := rand.NewSource(time.Now().UTC().UnixNano())
	r = rand.New(src)
}

// NextInt generates a random integer within the specified range, inclusive.
//
// The function takes two parameters, `min` and `max`, which define the lower and upper bounds of the range.
// It checks if `min` is greater than or equal to `max`, in which case it returns `min` as a default value.
// This behavior may need to be adjusted based on how you want to handle such cases (e.g., returning an error or panic).
//
// The function uses a package-level random generator to generate a random integer in the range from `min` to `max`,
// ensuring that both bounds are included in the result. The formula used is:
// r.Intn(max-min+1) + min, where `r` is a custom random generator that has been initialized elsewhere in the code.
//
// Parameters:
// - `min`: The lower bound of the random number range (inclusive).
// - `max`: The upper bound of the random number range (inclusive).
//
// Returns:
// - A random integer between `min` and `max`, including both bounds.
//
// Example:
//
// randomNum := NextInt(1, 10)
// fmt.Println("Random number between 1 and 10:", randomNum)
func NextInt(min, max int) int {
	if min >= max {
		return min // or handle the error appropriately
	}
	return r.Intn(max-min+1) + min // Ensure the range includes max
}

// NextReseed generates a random integer within the specified range, inclusive,
// and reseeds the random number generator with a new value each time it is called.
//
// The function first creates a new seed by combining the current UTC time in nanoseconds
// with a random integer from the custom random generator `r`. This helps to ensure that
// each call to `NextReseed` produces a different sequence of random numbers.
//
// If the provided `min` is greater than or equal to `max`, the function returns `min`
// as a default value. This behavior should be considered when using this function,
// as it may not be intuitive to return `min` in cases where the range is invalid.
//
// The function uses the reseeded random generator to generate a random integer in the
// inclusive range from `min` to `max`. The calculation ensures that both bounds are included
// in the result by using the formula: r.Intn(max-min+1) + min.
//
// Parameters:
// - `min`: The lower bound of the random number range (inclusive).
// - `max`: The upper bound of the random number range (inclusive).
//
// Returns:
// - A random integer between `min` and `max`, including both bounds.
//
// Example:
//
// randomNum := NextReseed(1, 10)
// fmt.Println("Random number between 1 and 10 after reseeding:", randomNum)
func NextReseed(min, max int) int {
	// Reseed the custom random generator with a new seed
	x := time.Now().UTC().UnixNano() + int64(r.Int())
	r.Seed(x)

	// Ensure max is included in the range
	if min >= max {
		return min
	}
	return r.Intn(max-min+1) + min
}

// NextUUID generates and returns a new UUID using the GenerateUUID function.
//
// If an error occurs during UUID generation (for example, if there is an issue reading from /dev/urandom),
// the function returns an empty string.
//
// This function is useful when you want a simple UUID generation without handling errors directly.
// It abstracts away the error handling by returning an empty string in case of failure.
//
// Returns:
// - A string representing the newly generated UUID.
// - An empty string if an error occurs during UUID generation.
//
// Example:
//
// uuid := NextUUID()
//
//	if uuid == "" {
//	    fmt.Println("Failed to generate UUID")
//	} else {
//
//	    fmt.Println("Generated UUID:", uuid)
//	}
func NextUUID() string {
	uuid, err := GenerateUUID()
	if err != nil {
		return ""
	}
	return uuid
}
