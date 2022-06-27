package models

import (
	"time"
)

type (
	Angkatan struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		SekolahID uint      `json:"sekolah_id"`
		Nama      string    `json:"nama"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
