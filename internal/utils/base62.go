package utils

import (
	"math/rand"
)

func EncodeBase62(n int) string {
	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if n == 0 {
		return "0"
	}
	result := ""

	for n > 0 {
		rem := n % 62
		result = string(chars[rem]) + result
		n = n / 62
	}

	return result
}

func GenerateRandomCode(length int) string {
	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// rand.Seed(time.Now().UnixNano())

	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result)
}
