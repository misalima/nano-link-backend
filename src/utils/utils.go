package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func HashURLWithRandom(url string) (string, error) {

	randomBytes := make([]byte, 4)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	randomNumber := fmt.Sprintf("%d", randomBytes[0])
	data := url + randomNumber

	hash := sha256.Sum256([]byte(data))
	return base64.URLEncoding.EncodeToString(hash[:])[:8], nil
}

func IsValidCustomShortID(id string) bool {
	for _, char := range id {
		if !(('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9') || char == '-' || char == '_') {
			return false
		}
	}
	return true
}
