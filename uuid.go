package unify4g

import (
	cr "crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// GenerateUUID generates a new universally unique identifier (UUID) using random data from /dev/urandom (Unix-based systems).
//
// This function opens the special file /dev/urandom to read 16 random bytes, which are then used to construct a UUID
// in the standard format (8-4-4-4-12 hex characters). It ensures that the file is properly closed after reading, even
// if an error occurs. If there's an error opening or reading from /dev/urandom, the function returns an appropriate error.
//
// UUID Format: The generated UUID is formatted as a string in the following structure:
// XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX, where X is a hexadecimal digit.
//
// Returns:
//   - A string representing the newly generated UUID.
//   - An error if there is an issue opening or reading from /dev/urandom.
//
// Example:
//
//	 uuid, err := GenerateUUID()
//		if err != nil {
//		    log.Fatalf("Failed to generate UUID: %v", err)
//		}
//	 fmt.Println("Generated UUID:", uuid)
//
// Notes:
//   - This function is designed for Unix-based systems. On non-Unix systems, this may not work because /dev/urandom
//     may not be available.
func GenerateUUID() (string, error) {
	dash := "-"
	return GenerateUUIDDelimiter(dash)
}

// GenerateUUIDDelimiter generates a new universally unique identifier (UUID) using random data from /dev/urandom
// (Unix-based systems) with a customizable delimiter.
//
// This function is similar to GenerateUUID but allows the user to specify a custom delimiter to separate
// different sections of the UUID. It opens the special file /dev/urandom to read 16 random bytes,
// which are then used to construct a UUID. The UUID is returned as a string in the format:
// XXXXXXXX<delimiter>XXXX<delimiter>XXXX<delimiter>XXXX<delimiter>XXXXXXXXXXXX, where X is a hexadecimal digit.
//
// Parameters:
//   - delimiter: A string used to separate sections of the UUID. Common choices are "-" or "" (no delimiter).
//
// Returns:
//   - A string representing the newly generated UUID with the specified delimiter.
//   - An error if there is an issue opening or reading from /dev/urandom.
//
// Example:
//
//	uuid, err := GenerateUUIDDelimiter("-")
//	if err != nil {
//	    log.Fatalf("Failed to generate UUID: %v", err)
//	}
//	fmt.Println("Generated UUID:", uuid)
//
// Notes:
//   - This function is designed for Unix-based systems. On non-Unix systems, it may not work because /dev/urandom
//     may not be available.
func GenerateUUIDDelimiter(delimiter string) (string, error) {
	file, err := os.Open("/dev/urandom")
	if err != nil {
		return "", fmt.Errorf("open /dev/urandom error:[%v]", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}()
	b := make([]byte, 16)
	_, err = file.Read(b)
	if err != nil {
		return "", err
	}
	// Format the bytes as a UUID string with the specified delimiter.
	// The UUID is structured as XXXXXXXX<delimiter>XXXX<delimiter>XXXX<delimiter>XXXX<delimiter>XXXXXXXXXXXX.
	uuid := fmt.Sprintf("%x%s%x%s%x%s%x%s%x", b[0:4], delimiter, b[4:6], delimiter, b[6:8], delimiter, b[8:10], delimiter, b[10:])
	return uuid, nil
}

// GenerateRandomID generates a random alphanumeric string of the specified length.
// This string includes uppercase letters, lowercase letters, and numbers, making it
// suitable for use as unique IDs or tokens.
//
// Parameters:
//   - length: The length of the random ID to generate. Must be a positive integer.
//
// Returns:
//   - A string of random alphanumeric characters with the specified length.
//
// The function uses a custom random source seeded with the current Unix timestamp
// in nanoseconds to ensure that each call produces a unique sequence.
// This function is intended to generate random strings quickly and is not
// cryptographically secure.
//
// Example:
//
//	id := GenerateRandomID(16)
//	fmt.Println("Generated Random ID:", id)
//
// Notes:
//   - This function is suitable for use cases where simple random IDs are needed.
//     However, for cryptographic purposes, consider using more secure random generation.
func GenerateRandomID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano())) // Create a seeded random generator for unique results each call
	// Allocate a byte slice for the generated ID and populate it with random characters
	id := make([]byte, length)
	for i := range id {
		id[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(id)
}

// GenerateCryptoID generates a cryptographically secure random ID as a hexadecimal string.
// It uses 16 random bytes, which are then encoded to a hexadecimal string for easy representation.
//
// Returns:
//   - A string representing a secure random hexadecimal ID of 32 characters (since 16 bytes are used, and each byte
//     is represented by two hexadecimal characters).
//
// The function uses crypto/rand.Read to ensure cryptographic security in the generated ID, making it suitable for
// sensitive use cases such as API keys, session tokens, or any security-critical identifiers.
//
// Example:
//
//	id := GenerateCryptoID()
//	fmt.Println("Generated Crypto ID:", id)
//
// Notes:
//   - This function is suitable for use cases where high security is required in the generated ID.
//   - It is not recommended for use cases where deterministic or non-cryptographic IDs are preferred.
func GenerateCryptoID() string {
	bytes := make([]byte, 16)
	// Use crypto/rand.Read for cryptographically secure random byte generation.
	if _, err := cr.Read(bytes); err != nil {
		log.Fatalf("Failed to generate secure random bytes: %v", err)
		return ""
	}
	return hex.EncodeToString(bytes)
}

// GenerateTimestampID generates a unique identifier based on the current Unix timestamp in nanoseconds,
// with an additional random integer to enhance uniqueness.
//
// This function captures the current time in nanoseconds since the Unix epoch and appends a random integer
// to ensure additional randomness and uniqueness, even if called in rapid succession. The result is returned
// as a string. This type of ID is well-suited for time-based ordering and can be useful for generating
// unique identifiers for logs, events, or non-cryptographic applications.
//
// Returns:
//   - A string representing the current Unix timestamp in nanoseconds, concatenated with a random integer.
//
// Example:
//
//	id := GenerateTimestampID()
//	fmt.Println("Generated Timestamp ID:", id)
//
// Notes:
//   - This function provides a unique, time-ordered identifier, but it is not suitable for cryptographic use.
//   - The combination of the current time and a random integer is best suited for applications requiring
//     uniqueness and ordering, rather than secure identifiers.
func GenerateTimestampID() string {
	return fmt.Sprintf("%d%d", time.Now().UnixNano(), nextInt())
}
