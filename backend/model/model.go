package model

import (
	"time"
)

type Model struct {
	ID        int       `db:"id" json:"id"`
	UID       string    `db:"uid" json:"uid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
