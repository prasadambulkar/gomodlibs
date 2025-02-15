package gomodlibs

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Call Following functions from your app by importing this module.

// GenerateRandomString generates a random string of the given length
// using characters from the allowed charset.

func GenerateRandomString(length int) (string, error) {
	result := make([]byte, length)

	for i := range result {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[randomIndex.Int64()]
	}

	return string(result), nil
}
