package stringutil

import (
	"crypto/rand"

	"github.com/rotisserie/eris"
)

const characterSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateRandomString creates a random string of length 'n'
// using the characterSet and the cryptographically secure rand reader.
func GenerateRandomString(n int) (string, error) {
	// 1. Create a byte slice to hold the random indices
	bytes := make([]byte, n)

	// 2. Read 'n' cryptographically random bytes.
	// This ensures the randomness is high quality.
	if _, err := rand.Read(bytes); err != nil {
		return "", eris.Wrap(err, "error generating random string")
	}

	// 3. Map the random bytes to indices within the characterSet
	for i, b := range bytes {
		// Use the modulo operator (%) to constrain the random byte value
		// to the length of our characterSet.
		bytes[i] = characterSet[b%byte(len(characterSet))]
	}

	// 4. Convert the resulting byte slice to a string
	return string(bytes), nil
}
