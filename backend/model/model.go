package model

import (
	"time"
)

type Model struct {
	UID       string    `json:"uid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var Tables = []interface{}{
	&User{},
	&Project{},
	&Comment{},
}
