package models

import (
	"time"
)

type (
	TahunAjaran struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		SekolahID uint      `json:"sekolah_id"`
		Nama      string    `json:"nama"`
		Semester  string    `json:"semester"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
