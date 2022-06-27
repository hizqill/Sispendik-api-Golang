package models

import (
	"time"
)

type (
	// Sekolah
	BeasiswaSiswas struct {
		ID                       uint      `gorm:"primary_key" json:"id"`
		SiswaId                  uint      `json:"siswa_id"`
		No_Kks                   string    `json:"no_kks" gorm:"default:null"`
		No_Pkh                   string    `json:"no_pkh" gorm:"default:null"`
		No_Kip                   string    `json:"no_kip" gorm:"default:null"`
		StatusPip                string    `json:"status_pip" gorm:"default:null"`
		AlasanMenerimaPip        string    `json:"alasan_menerima_pip" gorm:"default:null"`
		TahunMulaiMenerimaPip    string    `json:"tahun_mulai_menerima_pip" gorm:"default:null"`
		PeriodeMenerimaPip       string    `json:"periode_menerima_pip" gorm:"default:null"`
		StatusBeasiswaNonPip     string    `json:"status_beasiswa_non_pip" gorm:"default:null"`
		SumberBeasiswaNonPip     string    `json:"sumber_beasiswa_non_pip" gorm:"default:null"`
		JenisBeasiswaNonPip      string    `json:"jenis_beasiswa_non_pip" gorm:"default:null"`
		TahunMulaiMenerimaNonPip string    `json:"tahun_mulai_menerima_non_pip" gorm:"default:null"`
		JangkaWaktuNonPip        string    `json:"jangka_waktu_non_pip" gorm:"default:null"`
		BesarUangDiterima        string    `json:"besar_uang_diterima" gorm:"default:null"`
		CreatedAt                time.Time `json:"created_at"`
		UpdatedAt                time.Time `json:"updated_at"`
	}
)
