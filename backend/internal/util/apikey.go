package util

import (
	"crypto/rand"
	"encoding/hex"
)

// lm-api-{random-hex-string} format.
// total length : 64 (lm-api- + 57 random hex characters)
func GenerateRandomApiKey() string {
	
	const desiredHexLength = 57
	numBytes := (desiredHexLength + 1) / 2
	randomBytes := make([]byte, numBytes)
	rand.Read(randomBytes)

	hexStr := hex.EncodeToString(randomBytes)
	if len(hexStr) > desiredHexLength {
		hexStr = hexStr[:desiredHexLength]
	}

	return "lm-api-" + hexStr
}
