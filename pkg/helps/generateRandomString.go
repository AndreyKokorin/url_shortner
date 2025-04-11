package helps

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func GenerateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const charsetLen = int64(len(charset))

	if length <= 0 {
		return "", errors.New("length must be positive")
	}

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		// Генерируем случайное число от 0 до длины charset
		n, err := rand.Int(rand.Reader, big.NewInt(charsetLen))
		if err != nil {
			return "", err
		}
		result[i] = charset[n.Int64()]
	}

	return string(result), nil
}
