package services

import (
	"crypto/rand" //always we us this for keys not math/rand
	"encoding/hex"
)

func GenerateAPIKey() (string, error) {
	bytes := make([]byte, 32) //256 bits

	_, err := rand.Read(bytes) //fills with completely random data
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil //encodes in hex to which 64 char key only consists 0-9 and a-f
}
