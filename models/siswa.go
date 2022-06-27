package models

import (
	"time"
)

type (
	// siswa
	Siswa struct {
		ID                  uint      `gorm:"primary_key" json:"id"`
		Nama                string    `json:"nama"`
		UserID              uint      `json:"user_id"`
		SekolahID           uint      `json:"sekolah_id"`
		Nisn                string    `json:"nisn"`
		NisLokal            string    `json:"nis_lokal" gorm:"default:null"`
		NikSiswa            string    `json:"nik_siswa" gorm:"default:null"`
		TempatLahir         string    `json:"tempat_lahir" gorm:"default:null"`
		TanggalLahir        time.Time `json:"tanggal_lahir" gorm:"default:null"`
		JenisKelamin        string    `json:"jenis_kelamin" gorm:"default:null"`
		Agama               string    `json:"agama" gorm:"default:null"`
		Email               string    `json:"email" gorm:"default:null"`
		NoHpSiswa           string    `json:"no_hp_siswa" gorm:"default:null"`
		Hobi                string    `json:"hobi" gorm:"default:null"`
		CitaCita            string    `json:"cita_cita" gorm:"default:null"`
		HafalanJuz          string    `json:"hafalan_juz" gorm:"default:null"`
		HafalanSurat        string    `json:"hafalan_surat" gorm:"default:null"`
		JumlahSaudara       string    `json:"jumlah_saudara" gorm:"default:null"`
		TempatTinggal       string    `json:"tempat_tinggal" gorm:"default:null"`
		Alamat              string    `json:"alamat" gorm:"default:null"`
		Provinsi            string    `json:"provinsi" gorm:"default:null"`
		Kabupaten           string    `json:"kabupaten" gorm:"default:null"`
		Kecamatan           string    `json:"kecamatan" gorm:"default:null"`
		Desa                string    `json:"desa" gorm:"default:null"`
		KodePos             string    `json:"kode_pos" gorm:"default:null"`
		JarakRumah          string    `json:"jarak_rumah" gorm:"default:null"`
		Transportasi        string    `json:"transportasi" gorm:"default:null"`
		NoKk                string    `json:"no_kk" gorm:"default:null"`
		NamaAyah            string    `json:"nama_ayah" gorm:"default:null"`
		NikAyah             string    `json:"nik_ayah" gorm:"default:null"`
		NoHpAyah            string    `json:"no_hp_ayah" gorm:"default:null"`
		PendidikanAyah      string    `json:"pendidikan_ayah" gorm:"default:null"`
		PekerjaanAyah       string    `json:"pekerjaan_ayah" gorm:"default:null"`
		NamaIbu             string    `json:"nama_ibu" gorm:"default:null"`
		NikIbu              string    `json:"nik_ibu" gorm:"default:null"`
		NoHpIbu             string    `json:"no_hp_ibu" gorm:"default:null"`
		PendidikanIbu       string    `json:"pendidikan_ibu" gorm:"default:null"`
		PekerjaanIbu        string    `json:"pekerjaan_ibu" gorm:"default:null"`
		PenghasilanOrangtua string    `json:"penghasilan_orangtua" gorm:"default:null"`
		NamaWali            string    `json:"nama_wali" gorm:"default:null"`
		NikWali             string    `json:"nik_wali" gorm:"default:null"`
		NoHpWali            string    `json:"no_hp_wali" gorm:"default:null"`
		PendidikanWali      string    `json:"pendidikan_wali" gorm:"default:null"`
		PekerjaanWali       string    `json:"pekerjaan_wali" gorm:"default:null"`
		PenghasilanWali     string    `json:"penghasilan_wali" gorm:"default:null"`
		CreatedAt           time.Time `json:"created_at"`
		UpdatedAt           time.Time `json:"updated_at"`
	}
)
