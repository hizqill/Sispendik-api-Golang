package siswaController

import (
	"fmt"
	"net/http"
	"sispendik-api/models"
	s "sispendik-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

const (
	layoutDateTimeSiswa = "2006-01-02 15:04:05"
)

type UpdateInputSiswa struct {
	Nama                string `json:"nama" binding:"required"`
	Nisn                string `json:"nisn" binding:"required"`
	NisLokal            string `json:"nis_lokal"`
	NikSiswa            string `json:"nik_siswa"`
	TempatLahir         string `json:"tempat_lahir"`
	TanggalLahir        string `json:"tanggal_lahir"`
	JenisKelamin        string `json:"jenis_kelamin"`
	Agama               string `json:"agama"`
	Email               string `json:"email" binding:"email"`
	NoHpSiswa           string `json:"no_hp_siswa"`
	Hobi                string `json:"hobi"`
	CitaCita            string `json:"cita_cita"`
	HafalanJuz          string `json:"hafalan_juz"`
	HafalanSurat        string `json:"hafalan_surat"`
	JumlahSaudara       string `json:"jumlah_saudara"`
	TempatTinggal       string `json:"tempat_tinggal"`
	Alamat              string `json:"alamat"`
	Provinsi            string `json:"provinsi"`
	Kabupaten           string `json:"kabupaten"`
	Kecamatan           string `json:"kecamatan"`
	Desa                string `json:"desa"`
	KodePos             string `json:"kode_pos"`
	JarakRumah          string `json:"jarak_rumah"`
	Transportasi        string `json:"transportasi"`
	NoKk                string `json:"no_kk"`
	NamaAyah            string `json:"nama_ayah"`
	NikAyah             string `json:"nik_ayah"`
	NoHpAyah            string `json:"no_hp_ayah"`
	PendidikanAyah      string `json:"pendidikan_ayah"`
	PekerjaanAyah       string `json:"pekerjaan_ayah"`
	NamaIbu             string `json:"nama_ibu"`
	NikIbu              string `json:"nik_ibu"`
	NoHpIbu             string `json:"no_hp_ibu"`
	PendidikanIbu       string `json:"pendidikan_ibu"`
	PekerjaanIbu        string `json:"pekerjaan_ibu"`
	PenghasilanOrangtua string `json:"penghasilan_orangtua"`
	NamaWali            string `json:"nama_wali"`
	NikWali             string `json:"nik_wali"`
	NoHpWali            string `json:"no_hp_wali"`
	PendidikanWali      string `json:"pendidikan_wali"`
	PekerjaanWali       string `json:"pekerjaan_wali"`
	PenghasilanWali     string `json:"penghasilan_wali"`
}

// GetSiswaByIdUserid godoc
// @Summary Mendapatkan data siswa berdasarkan user ID.
// @Description Mendapatkan data siswa berdasarkan user ID.
// @Tags Siswa
// @Produce json
// @Param id path string true "User id"
// @Success 200 {object} s.Siswa
// @Router /siswa/user/{id} [get]
func GetSiswaByIdUserid(c *gin.Context) {
	// Get model if exista
	var siswa s.Siswa

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("user_id = ?", c.Param("id")).First(&siswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": siswa})
}

// GetSiswaById godoc
// @Summary Mendapatkan data siswa.
// @Description Mendapatkan data siswa berdasarkan id.
// @Tags Siswa
// @Produce json
// @Param id path string true "Siswa id"
// @Success 200 {object} s.Siswa
// @Router /siswa/{id} [get]
func GetSiswaById(c *gin.Context) {
	// Get model if exista
	var siswa s.Siswa

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&siswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": siswa})
}

// UpdateSiswa godoc
// @Summary Memperbarui data Siswa.
// @Description memperbaharui data siswa berdasarkan id.
// @Tags Siswa
// @Produce json
// @Param id path string true "Siswa id"
// @Param Body body UpdateInputSiswa true "the body to update Siswa"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} s.Siswa
// @Router /siswa/{id} [patch]
func UpdateSiswa(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var siswa s.Siswa

	if err := db.Where("id = ?", c.Param("id")).First(&siswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var inputSiswa UpdateInputSiswa
	errInput := c.ShouldBindJSON(&inputSiswa)

	if errInput != nil {
		errorMessanges := []string{}
		for _, e := range errInput.(validator.ValidationErrors) {
			errorMessange := fmt.Sprintf("Error on field %s, candition: %s", e.Field(), e.ActualTag())
			errorMessanges = append(errorMessanges, errorMessange)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessanges,
		})
		return
	}

	TanggalLahir, errTanggalLahir := time.Parse(layoutDateTimeSiswa, inputSiswa.TanggalLahir)

	if errTanggalLahir != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errTanggalLahir.Error()})
		return
	}

	var updatedInput s.Siswa
	updatedInput.Nama = inputSiswa.Nama
	updatedInput.Nisn = inputSiswa.Nisn
	updatedInput.NisLokal = inputSiswa.NisLokal
	updatedInput.NikSiswa = inputSiswa.NikSiswa
	updatedInput.TempatLahir = inputSiswa.TempatLahir
	updatedInput.TanggalLahir = TanggalLahir
	updatedInput.JenisKelamin = inputSiswa.JenisKelamin
	updatedInput.Agama = inputSiswa.Agama
	updatedInput.Email = inputSiswa.Email
	updatedInput.NoHpSiswa = inputSiswa.NoHpSiswa
	updatedInput.Hobi = inputSiswa.Hobi
	updatedInput.CitaCita = inputSiswa.CitaCita
	updatedInput.HafalanJuz = inputSiswa.HafalanJuz
	updatedInput.HafalanSurat = inputSiswa.HafalanSurat
	updatedInput.JumlahSaudara = inputSiswa.JumlahSaudara
	updatedInput.TempatTinggal = inputSiswa.TempatTinggal
	updatedInput.Alamat = inputSiswa.Alamat
	updatedInput.Provinsi = inputSiswa.Provinsi
	updatedInput.Kabupaten = inputSiswa.Kabupaten
	updatedInput.Kecamatan = inputSiswa.Kecamatan
	updatedInput.Desa = inputSiswa.Desa
	updatedInput.KodePos = inputSiswa.KodePos
	updatedInput.JarakRumah = inputSiswa.JarakRumah
	updatedInput.Transportasi = inputSiswa.Transportasi
	updatedInput.NoKk = inputSiswa.NoKk
	updatedInput.NamaAyah = inputSiswa.NamaAyah
	updatedInput.NikAyah = inputSiswa.NikAyah
	updatedInput.NoHpAyah = inputSiswa.NoHpAyah
	updatedInput.PendidikanAyah = inputSiswa.PendidikanAyah
	updatedInput.PekerjaanAyah = inputSiswa.PekerjaanAyah
	updatedInput.NamaIbu = inputSiswa.NamaIbu
	updatedInput.NikIbu = inputSiswa.NikIbu
	updatedInput.NoHpIbu = inputSiswa.NoHpIbu
	updatedInput.PendidikanIbu = inputSiswa.PendidikanIbu
	updatedInput.PekerjaanIbu = inputSiswa.PekerjaanIbu
	updatedInput.PenghasilanOrangtua = inputSiswa.PenghasilanOrangtua
	updatedInput.NamaWali = inputSiswa.NamaWali
	updatedInput.NikWali = inputSiswa.NikWali
	updatedInput.NoHpWali = inputSiswa.NoHpWali
	updatedInput.PendidikanWali = inputSiswa.PendidikanWali
	updatedInput.PekerjaanWali = inputSiswa.PekerjaanWali
	updatedInput.PenghasilanWali = inputSiswa.PenghasilanWali
	updatedInput.UpdatedAt = time.Now()

	db.Model(&siswa).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": siswa})
}

// DeleteSiswa godoc
// @Summary Menghapus data siswa.
// @Description Menghapus semua data siswa berdasarkan id. mulai dari akun, biodata, asal sekolah, prestasi, beasiswa.
// @Tags Sekolah
// @Produce json
// @Param id path string true "Siswa id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]string
// @Router /siswa/{id} [delete]
func DeleteSiswa(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var siswa s.Siswa

	if err := db.Where("id = ?", c.Param("id")).First(&siswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var user models.Users
	if err := db.Where("id = ?", siswa.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record user not found!"})
		return
	}

	db.Delete(&user)
	db.Delete(&siswa)

	c.JSON(http.StatusOK, gin.H{"data": "Berhasil Menghapus Data"})
}
