package models

import (
	"time"
)

type Sekolah struct {
	ID                                uint      `gorm:"primary_key" json:"id"`
	Nama                              string    `json:"nama"`
	UserID                            uint      `json:"user_id"`
	StatusMenuSiswa                   string    `json:"status_menu_siswa"`
	StatusMenuPendidik                string    `json:"status_menu_pendidik"`
	PendaftaranSiswa                  string    `json:"pendaftaran_siswa"`
	KodeVerifikasiPendaftaranSiswa    string    `json:"kode_verifikasi_pendaftaran_siswa" gorm:"default:null"`
	PendaftaranPendidik               string    `json:"pendaftaran_pendidik"`
	KodeVerifikasiPendaftaranPendidik string    `json:"kode_verifikasi_pendaftaran_pendidik" gorm:"default:null"`
	CreatedAt                         time.Time `json:"created_at"`
	UpdatedAt                         time.Time `json:"updated_at"`
}
