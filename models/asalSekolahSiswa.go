package models

import (
	"time"
)

type (
	// Sekolah
	AsalsekolahSiswas struct {
		ID               uint      `gorm:"primary_key" json:"id"`
		SiswaID          uint      `json:"siswa_id"`
		JenisSekolah     string    `json:"jenis_sekolah"  gorm:"default:null"`
		StatusSekolah    string    `json:"status_sekolah"  gorm:"default:null"`
		LokasiSekolah    string    `json:"lokasi_sekolah"  gorm:"default:null"`
		NoPesertaUN      string    `json:"no_peserta_un"  gorm:"default:null"`
		NpsnSekolah      string    `json:"npsn_sekolah"  gorm:"default:null"`
		NoBlankoSKHUN    string    `json:"no_blanko_skhun"  gorm:"default:null"`
		NoSeriIjazah     string    `json:"no_seri_ijazah"  gorm:"default:null"`
		TotalNilaiUN     string    `json:"total_nilai_un"  gorm:"default:null"`
		TanggalKelulusan time.Time `json:"tanggal_kelulusan"  gorm:"default:null"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
	}
)
