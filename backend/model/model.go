package model

import (
	"github.com/cameronbrill/fig-issue/backend/crypto/aes256"
)

// UID implements and encoder and decoder for encrypting external uids (such as supabase user uid)
type UID struct {
	encryptionStore aes256.EncryptionStore
}

func (u *UID) Decode() (string, error) {
	decrypted, err := u.encryptionStore.Decrypt()
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

func (u *UID) Encode(uid string) error {
	return u.encryptionStore.Encrypt([]byte(uid))
}

type Comment struct {
	Message  string
	Mentions []string
}
