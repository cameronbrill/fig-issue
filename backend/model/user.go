package model

type User struct {
	Model
	SupaBaseUID UID `json:"supa_base_uid"`
	Projects    []Project
}
