package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func EncryptPost(value, key string) (string, error) {
	plaintext := []byte(value)
	keyBytes := []byte(key)
	if block, err := aes.NewCipher(keyBytes); err != nil {
		return "", err
	} else {
		iv := make([]byte, aes.BlockSize)
		stream := cipher.NewCTR(block, iv)

		cipherText := make([]byte, len(plaintext))
		stream.XORKeyStream(cipherText, plaintext)
		return base64.StdEncoding.EncodeToString(cipherText), nil
	}
}

func DecryptPost(value, key string) (string, error) {
	if encoded, err := base64.StdEncoding.DecodeString(value); err != nil {
		return "", err
	} else {
		keyBytes := []byte(key)

		block, err := aes.NewCipher(keyBytes)
		if err != nil {
			return "", err
		}
		iv := make([]byte, aes.BlockSize)
		stream := cipher.NewCTR(block, iv)

		plaintext := make([]byte, len(encoded))
		stream.XORKeyStream(plaintext, encoded)

		return string(plaintext), nil
	}

}
