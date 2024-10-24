package unify4g

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

// NextFloat64 returns the next random float64 value in the range [0.0, 1.0).
//
// This function uses the rand package to generate a random float64 value.
// The generated value is uniformly distributed over the interval [0.0, 1.0).
//
// Returns:
//   - A random float64 value between 0.0 and 1.0.
func NextFloat64() float64 {
	return rand.Float64()
}

// NextFloat64Bounded returns the next random float64 value bounded by the specified range.
//
// Parameters:
//   - `start`: The lower bound of the random float64 value (inclusive).
//   - `end`: The upper bound of the random float64 value (exclusive).
//
// Returns:
//   - A random float64 value uniformly distributed between `start` and `end`.
func NextFloat64Bounded(start float64, end float64) float64 {
	return rand.Float64()*(end-start) + start
}

// NextFloat32 returns the next random float32 value in the range [0.0, 1.0).
//
// This function uses the rand package to generate a random float32 value.
// The generated value is uniformly distributed over the interval [0.0, 1.0).
//
// Returns:
//   - A random float32 value between 0.0 and 1.0.
func NextFloat32() float32 {
	return rand.Float32()
}

// NextFloat32Bounded returns the next random float32 value bounded by the specified range.
//
// Parameters:
//   - `start`: The lower bound of the random float32 value (inclusive).
//   - `end`: The upper bound of the random float32 value (exclusive).
//
// Returns:
//   - A random float32 value uniformly distributed between `start` and `end`.
func NextFloat32Bounded(start float32, end float32) float32 {
	return rand.Float32()*(end-start) + start
}

// NextIntBounded returns the next random int value bounded by the specified range.
//
// Parameters:
//   - `start`: The lower bound of the random int value (inclusive).
//   - `end`: The upper bound of the random int value (exclusive).
//
// Returns:
//   - A random int value uniformly distributed between `start` and `end`.
func NextIntBounded(start int, end int) int {
	return start + rand.Intn(end-start)
}

// NextIntUpperBounded returns the next random int value bounded by a maximum value.
//
// Parameters:
//   - `end`: The exclusive upper bound for the random int value.
//
// Returns:
//   - A random int value uniformly distributed in the range [0, end).
func NextIntUpperBounded(end int) int {
	return rand.Intn(end)
}

// NextBytes creates an array of random bytes with the specified length.
//
// Parameters:
//   - `count`: The number of random bytes to generate.
//
// Returns:
//   - A slice of random bytes with the specified length.
func NextBytes(count int) []byte {
	a := make([]byte, count)
	for i := range a {
		a[i] = (byte)(nextInt())
	}
	return a
}

// nextInt returns the next random int value.
//
// This helper function generates a random int value using the rand package.
// It is called by NextBytes to populate the byte array with random values.
//
// Returns:
//   - A random int value.
func nextInt() int {
	return rand.Int()
}
