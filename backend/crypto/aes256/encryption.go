package aes256

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"os"
	"strings"
	"unicode"
)

// credit to author of https://goplay.space/#F_vCnIbq8wY

type EncryptionStore struct {
	IV         []byte `json:"iv"`
	CipherText []byte `json:"cipher_text"`
}

// Encrypt uses an encryption key from the callers environment to encrypt content
func (e *EncryptionStore) Encrypt(content []byte) (err error) {
	var encryptionKey []byte = []byte(os.Getenv("AES_ENCRYPTION_KEY"))
	bPlaintext := pkcs5Padding(content, aes.BlockSize)
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return err
	}

	e.IV, _ = generateRandomBytes(block.BlockSize())

	e.CipherText = make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, e.IV)
	mode.CryptBlocks(e.CipherText, bPlaintext)

	return err
}

// Decrypt uses an encryption key from the callers environment
// alongside previously stored initialization vector and encrypted cipher text
// to decrypt content
func (e *EncryptionStore) Decrypt() (decryptedContent []byte, err error) {
	var encryptionKey []byte = []byte(os.Getenv("AES_ENCRYPTION_KEY"))
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, e.IV)
	mode.CryptBlocks(e.CipherText, e.CipherText)

	cutTrailingSpaces := []byte(extraTrim(strings.TrimSpace(string(e.CipherText))))
	return cutTrailingSpaces, err
}

func extraTrim(s string) string {
	clean := strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, s)
	return clean
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipherText, padText...)
}
