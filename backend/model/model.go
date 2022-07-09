package model

import (
	"time"
)

type Model struct {
	ID        int       `json:"id"`
	UID       string    `json:"uid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
