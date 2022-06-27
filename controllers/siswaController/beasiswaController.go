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

type UpdateInputBeasiswaSiswa struct {
	No_Kks                   string `json:"no_kks"`
	No_Pkh                   string `json:"no_pkh"`
	No_Kip                   string `json:"no_kip"`
	StatusPip                string `json:"status_pip"`
	AlasanMenerimaPip        string `json:"alasan_menerima_pip"`
	TahunMulaiMenerimaPip    string `json:"tahun_mulai_menerima_pip"`
	PeriodeMenerimaPip       string `json:"periode_menerima_pip"`
	StatusBeasiswaNonPip     string `json:"status_beasiswa_non_pip"`
	SumberBeasiswaNonPip     string `json:"sumber_beasiswa_non_pip"`
	JenisBeasiswaNonPip      string `json:"jenis_beasiswa_non_pip"`
	TahunMulaiMenerimaNonPip string `json:"tahun_mulai_menerima_non_pip"`
	JangkaWaktuNonPip        string `json:"jangka_waktu_non_pip"`
	BesarUangDiterima        string `json:"besar_uang_diterima"`
}

// GetBeasiswaById godoc
// @Summary Mendapatkan data beasiswa siswa.
// @Description Mendapatkan data beasiswa siswa berdasarkan id.
// @Tags Siswa -> Beasiswa
// @Produce json
// @Param id path string true "Beasiswa_siswa ID"
// @Success 200 {object} s.BeasiswaSiswas
// @Router /siswa/beasiswa/{id} [get]
func GetBeasiswaById(c *gin.Context) {
	// Get model if exista
	var siswa s.BeasiswaSiswas

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&siswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": siswa})
}

// GetBeasiswaByIdSiswa godoc
// @Summary Mendapatkan data beasiswa siswa berdasarkan id siswa.
// @Description Mendapatkan data beasiswa siswa berdasarkan id siswa.
// @Tags Siswa -> Beasiswa
// @Produce json
// @Param id path string true "Siswa ID"
// @Success 200 {object} s.BeasiswaSiswas
// @Router /siswa/{id}/beasiswa [get]
func GetBeasiswaByIdSiswa(c *gin.Context) {
	// Get model if exista
	var siswa s.BeasiswaSiswas

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("siswa_id = ?", c.Param("id")).First(&siswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": siswa})
}

// UpdateBeasiswa godoc
// @Summary Memperbarui data beasiswa siswa.
// @Description memperbaharui data beasiswa siswa berdasarkan id.
// @Tags Siswa -> Beasiswa
// @Produce json
// @Param id path string true "Beasiswa_siswa id"
// @Param Body body UpdateInputBeasiswaSiswa true "the body to update beasiswa Siswa"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} s.BeasiswaSiswas
// @Router /siswa/beasiswa/{id} [patch]
func UpdateBeasiswa(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var siswa s.BeasiswaSiswas

	if err := db.Where("id = ?", c.Param("id")).First(&siswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var inputbeasiswa UpdateInputBeasiswaSiswa

	errInput := c.ShouldBindJSON(&inputbeasiswa)

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

	var updatedInput s.BeasiswaSiswas
	updatedInput.No_Kks = inputbeasiswa.No_Kks
	updatedInput.No_Pkh = inputbeasiswa.No_Pkh
	updatedInput.No_Kip = inputbeasiswa.No_Kip
	updatedInput.StatusPip = inputbeasiswa.StatusPip
	updatedInput.AlasanMenerimaPip = inputbeasiswa.AlasanMenerimaPip
	updatedInput.TahunMulaiMenerimaPip = inputbeasiswa.TahunMulaiMenerimaPip
	updatedInput.PeriodeMenerimaPip = inputbeasiswa.PeriodeMenerimaPip
	updatedInput.StatusBeasiswaNonPip = inputbeasiswa.StatusBeasiswaNonPip
	updatedInput.SumberBeasiswaNonPip = inputbeasiswa.SumberBeasiswaNonPip
	updatedInput.JenisBeasiswaNonPip = inputbeasiswa.JenisBeasiswaNonPip
	updatedInput.TahunMulaiMenerimaNonPip = inputbeasiswa.TahunMulaiMenerimaNonPip
	updatedInput.JangkaWaktuNonPip = inputbeasiswa.JangkaWaktuNonPip
	updatedInput.BesarUangDiterima = inputbeasiswa.BesarUangDiterima
	updatedInput.UpdatedAt = time.Now()

	db.Model(&siswa).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": siswa})
}
