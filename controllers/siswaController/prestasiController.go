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

type UpdateInputPrestasiSiswa struct {
	JenisLomba          string `json:"jenis_lomba" `
	TingkatLomba        string `json:"tingkat_lomba" `
	PeringkatLomba      string `json:"peringkat_lomba" `
	BidangPrestasi      string `json:"bidang_prestasi" `
	TingkatPrestasi     string `json:"tingkat_prestasi" `
	PeringkatPrestasi   string `json:"peringkat_prestasi" `
	TahunMeraihPrestasi uint   `json:"tahun_meraih_prestasi" `
}

// GetPrestasiById godoc
// @Summary Mendapatkan data prestasi siswa.
// @Description Mendapatkan data prestasi siswa berdasarkan id.
// @Tags Siswa -> Prestasi
// @Produce json
// @Param id path string true "Prestasi_siswa ID"
// @Success 200 {object} s.PrestasiSiswas
// @Router /siswa/prestasi/{id} [get]
func GetPrestasiById(c *gin.Context) {
	// Get model if exista
	var siswa s.PrestasiSiswas

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&siswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": siswa})
}

// GetPrestasiByIdSiswa godoc
// @Summary Mendapatkan data prestasi siswa berdasarkan id siswa.
// @Description Mendapatkan data prestasi siswa berdasarkan id siswa.
// @Tags Siswa -> Prestasi
// @Produce json
// @Param id path string true "Siswa ID"
// @Success 200 {object} s.PrestasiSiswas
// @Router /siswa/{id}/prestasi [get]
func GetPrestasiByIdSiswa(c *gin.Context) {
	// Get model if exista
	var siswa s.PrestasiSiswas

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("siswa_id = ?", c.Param("id")).First(&siswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": siswa})
}

// UpdatePrestasi godoc
// @Summary Memperbarui data prestasi siswa.
// @Description memperbaharui data prestasi siswa berdasarkan id.
// @Tags Siswa -> Prestasi
// @Produce json
// @Param id path string true "Prestasi_siswa id"
// @Param Body body UpdateInputPrestasiSiswa true "the body to update prestasi Siswa"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} s.PrestasiSiswas
// @Router /siswa/prestasi/{id} [patch]
func UpdatePrestasi(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var siswa s.PrestasiSiswas

	if err := db.Where("id = ?", c.Param("id")).First(&siswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var inputbeasiswa UpdateInputPrestasiSiswa

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

	var updatedInput s.PrestasiSiswas
	updatedInput.JenisLomba = inputbeasiswa.JenisLomba
	updatedInput.TingkatLomba = inputbeasiswa.TingkatLomba
	updatedInput.PeringkatLomba = inputbeasiswa.PeringkatLomba
	updatedInput.BidangPrestasi = inputbeasiswa.BidangPrestasi
	updatedInput.TingkatPrestasi = inputbeasiswa.TingkatPrestasi
	updatedInput.PeringkatPrestasi = inputbeasiswa.PeringkatPrestasi
	updatedInput.TahunMeraihPrestasi = inputbeasiswa.TahunMeraihPrestasi
	updatedInput.UpdatedAt = time.Now()

	db.Model(&siswa).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": siswa})
}
