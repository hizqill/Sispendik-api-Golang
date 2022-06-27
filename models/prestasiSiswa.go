package models

import (
	"time"
)

type (
	// Sekolah
	PrestasiSiswas struct {
		ID                  uint      `gorm:"primary_key" json:"id"`
		SiswaID             uint      `json:"siswa_id"`
		JenisLomba          string    `json:"jenis_lomba" gorm:"default:null"`
		TingkatLomba        string    `json:"tingkat_lomba" gorm:"default:null"`
		PeringkatLomba      string    `json:"peringkat_lomba" gorm:"default:null"`
		BidangPrestasi      string    `json:"bidang_prestasi" gorm:"default:null"`
		TingkatPrestasi     string    `json:"tingkat_prestasi" gorm:"default:null"`
		PeringkatPrestasi   string    `json:"peringkat_prestasi" gorm:"default:null"`
		TahunMeraihPrestasi uint      `json:"tahun_meraih_prestasi" gorm:"default:null"`
		CreatedAt           time.Time `json:"created_at"`
		UpdatedAt           time.Time `json:"updated_at"`
	}
)
