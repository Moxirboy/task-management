package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func Hash(data []byte) string {
	hash := sha256.New()
	hash.Write(data)

	hashed := hash.Sum(nil)
	return fmt.Sprintf("%x", hashed)
}

func Hash2(data []byte) string {
	hash := sha256.New()
	hash.Write(data)

	hashed := hash.Sum(nil)
	// Encode to base64 to get a shorter representation
	encoded := base64.RawURLEncoding.EncodeToString(hashed)

	// Truncate to 7 characters
	return encoded[:7]
}
