package utils

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/google/uuid"
)

func GenerateSecretToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return "clxtkn_" + hex.EncodeToString(bytes), nil
}

func GenerateAccountID() string {
	uuidStr := "acc_" + uuid.New().String()
	shortID := uuidStr[:5]
	return "acc_" + shortID
}
