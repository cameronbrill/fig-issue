package aes256

import (
	"crypto/aes"
	"errors"
	"os"
	"testing"
)

func TestEncryption(t *testing.T) {
	emptyStr := ""
	tests := []struct {
		name             string
		input            string
		encryptionKey    string
		hostileEnvAttack *string
		expectedError    error
	}{
		{
			name:          "test valid",
			encryptionKey: "1234567890123456",
			input:         "Hello World",
		},
		{
			name:          "test valid different key and input",
			encryptionKey: "2234567890123456",
			input:         "ello World",
		},
		{
			name:          "test valid different key and input longer than 16 characters",
			encryptionKey: "2234567890123456",
			input:         "Hello World sixteen characters",
		},
		{
			name:          "test invalid encryption key",
			encryptionKey: "",
			input:         "string",
			expectedError: aes.KeySizeError(0),
		},
		{
			name:             "test hostile take over",
			encryptionKey:    "2234567890123456",
			hostileEnvAttack: &emptyStr,
			input:            "string",
			expectedError:    aes.KeySizeError(0),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("AES_ENCRYPTION_KEY", tc.encryptionKey)
			var encryptionStore EncryptionStore
			err := encryptionStore.Encrypt([]byte(tc.input))
			if err != nil {
				if errors.Is(err, tc.expectedError) {
					return
				}
				t.Fatalf("%+v is not equal to %+v", err, tc.expectedError)
			}
			if tc.hostileEnvAttack != nil {
				os.Setenv("AES_ENCRYPTION_KEY", *tc.hostileEnvAttack)
			}
			decryptedValue, err := encryptionStore.Decrypt()
			if err != nil {
				if errors.Is(err, tc.expectedError) {
					return
				}
				t.Fatalf("%+v is not equal to %+v", err, tc.expectedError)
			}
			if string(decryptedValue) != tc.input {
				t.Fatalf("decrypted value does not match input: len(%d) %s != len(%d) %s", len(string(decryptedValue)), string(decryptedValue), len(tc.input), tc.input)
			}
		})
	}
}
