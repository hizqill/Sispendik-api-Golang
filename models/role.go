package models

type (
	// Role User
	Roles struct {
		ID   uint   `gorm:"primary_key" json:"id"`
		Nama string `json:"nama"`
	}
)
