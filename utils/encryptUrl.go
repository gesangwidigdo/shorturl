package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
)

func EncryptURL(plainURL string) (string, error) {
	secretKey := os.Getenv("AES_KEY")
	keyBytes := []byte(secretKey)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, []byte(plainURL), nil)

	encryptedUrl := base64.URLEncoding.EncodeToString(cipherText)

	return encryptedUrl[:6], nil
}
