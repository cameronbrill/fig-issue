package model

import (
	"time"

	"github.com/cameronbrill/fig-issue/backend/crypto/aes256"
)

type Model struct {
	ID        int       `json:"id"`
	UID       string    `json:"uid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

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
	Model
	Message  string
	Mentions []string
}

type User struct {
	Model
	SupaBaseUID UID `json:"supa_base_uid"`
	Projects    []Project
}

type Project struct {
	Model
	Users []User
}
