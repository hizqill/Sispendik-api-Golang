package siswaController

import (
	"fmt"
	"net/http"
	s "sispendik-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdateInputAsalSekolahSiswa struct {
	JenisSekolah     string `json:"jenis_sekolah"`
	StatusSekolah    string `json:"status_sekolah"`
	LokasiSekolah    string `json:"lokasi_sekolah"`
	NoPesertaUN      string `json:"no_peserta_un"`
	NpsnSekolah      string `json:"npsn_sekolah"`
	NoBlankoSKHUN    string `json:"no_blanko_skhun"`
	NoSeriIjazah     string `json:"no_seri_ijazah"`
	TotalNilaiUN     string `json:"total_nilai_un"`
	TanggalKelulusan string `json:"tanggal_kelulusan"`
}

// GetAsalSekolahById godoc
// @Summary Mendapatkan data asal sekolah siswa.
// @Description Mendapatkan data asal sekolah siswa berdasarkan id.
// @Tags Siswa -> Asal Sekolah
// @Produce json
// @Param id path string true "asalSekolah_siswa ID"
// @Success 200 {object} s.AsalsekolahSiswas
// @Router /siswa/asal-sekolah/{id} [get]
func GetAsalSekolahById(c *gin.Context) {
	// Get model if exista
	var siswa s.AsalsekolahSiswas

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&siswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": siswa})
}

// GetAsalSekolahByIdSiswa godoc
// @Summary Mendapatkan data asal sekolah siswa berdasarkan id siswa.
// @Description Mendapatkan data asal sekolah siswa berdasarkan id siswa.
// @Tags Siswa -> Asal Sekolah
// @Produce json
// @Param id path string true "Siswa ID"
// @Success 200 {object} s.AsalsekolahSiswas
// @Router /siswa/{id}/asal-sekolah [get]
func GetAsalSekolahByIdSiswa(c *gin.Context) {
	// Get model if exista
	var siswa s.AsalsekolahSiswas

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("siswa_id = ?", c.Param("id")).First(&siswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": siswa})
}

// UpdateAsalSekolah godoc
// @Summary Memperbarui data asal sekolah siswa.
// @Description memperbaharui data asal sekolah siswa berdasarkan id.
// @Tags Siswa -> Asal Sekolah
// @Produce json
// @Param id path string true "asal_Sekolah id"
// @Param Body body UpdateInputAsalSekolahSiswa true "the body to update beasiswa Siswa"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} s.AsalsekolahSiswas
// @Router /siswa/asal-sekolah/{id} [patch]
func UpdateAsalSekolah(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var siswa s.AsalsekolahSiswas

	if err := db.Where("id = ?", c.Param("id")).First(&siswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var inputAsalSekolah UpdateInputAsalSekolahSiswa

	errInput := c.ShouldBindJSON(&inputAsalSekolah)

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

	TanggalKelulusan, errTanggalLahir := time.Parse(layoutDateTimeSiswa, inputAsalSekolah.TanggalKelulusan)

	if errTanggalLahir != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errTanggalLahir.Error()})
		return
	}

	var updatedInput s.AsalsekolahSiswas
	updatedInput.JenisSekolah = inputAsalSekolah.JenisSekolah
	updatedInput.StatusSekolah = inputAsalSekolah.StatusSekolah
	updatedInput.LokasiSekolah = inputAsalSekolah.LokasiSekolah
	updatedInput.NoPesertaUN = inputAsalSekolah.NoPesertaUN
	updatedInput.NpsnSekolah = inputAsalSekolah.NpsnSekolah
	updatedInput.NoBlankoSKHUN = inputAsalSekolah.NoBlankoSKHUN
	updatedInput.NoSeriIjazah = inputAsalSekolah.NoSeriIjazah
	updatedInput.TotalNilaiUN = inputAsalSekolah.TotalNilaiUN
	updatedInput.TanggalKelulusan = TanggalKelulusan

	db.Model(&siswa).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": siswa})
}
