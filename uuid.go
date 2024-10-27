package unify4g

import (
	"fmt"
	"os"
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
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, nil
}
