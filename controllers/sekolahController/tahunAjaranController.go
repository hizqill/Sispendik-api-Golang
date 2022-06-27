package sekolahController

import (
	"fmt"
	"net/http"
	sekolah "sispendik-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type InputTahunAjaran struct {
	SekolahID uint   `json:"sekolah_id" binding:"required,number"`
	Nama      string `json:"nama" binding:"required"`
	Semester  string `json:"semester" binding:"required"`
}

type UpdateInputTahunAjaran struct {
	Nama     string `json:"nama" binding:"required"`
	Semester string `json:"semester" binding:"required"`
}

// GetTahunAjaranbyIdSekolah godoc
// @Summary mendapatkan data tahun ajaran berdasarkan ID Sekolah.
// @Description Mendapatkan semua data tahun ajaran Berdasarkan ID Sekolah.
// @Tags Sekolah -> Tahun Ajaran
// @Produce json
// @Param id path string true "Sekolah id"
// @Success 200 {object} []sekolah.TahunAjaran
// @Router /sekolah/{id}/tahun-ajaran [get]
func GetTahunAjaranbyIdSekolah(c *gin.Context) { // Get model if exist
	var tahunAjaran []sekolah.TahunAjaran

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("sekolah_id  = ?", c.Param("id")).Find(&tahunAjaran).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tahunAjaran})
}

// GetTahunAjaranbyId godoc
// @Summary Mendapatkan data tahun ajaran.
// @Description Mendapatkan data tahun ajaran berdasarkan id.
// @Tags Sekolah -> Tahun Ajaran
// @Produce json
// @Param id path string true "Tahun ajaran id"
// @Success 200 {object} sekolah.TahunAjaran
// @Router /sekolah/tahun-ajaran/{id} [get]
func GetTahunAjaranbyId(c *gin.Context) { // Get model if exist
	var tahunAjaran sekolah.TahunAjaran

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&tahunAjaran).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tahunAjaran})
}

// CreateTahunAjaran godoc
// @Summary Membuat data tahun ajaran barun.
// @Description Membuat data tahun ajaran baru.
// @Tags Sekolah -> Tahun Ajaran
// @Param Body body InputTahunAjaran true "the body to create a new tahun ajaran"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} sekolah.TahunAjaran
// @Router /sekolah/tahun-ajaran [post]
func CreateTahunAjaran(c *gin.Context) {
	// Validate input
	var InputTahunAjaran InputTahunAjaran
	errInput := c.ShouldBindJSON(&InputTahunAjaran)

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

	// Create Rating
	tahunAjaran := sekolah.TahunAjaran{
		SekolahID: InputTahunAjaran.SekolahID,
		Semester:  InputTahunAjaran.Semester,
		Nama:      InputTahunAjaran.Nama,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&tahunAjaran)

	c.JSON(http.StatusOK, gin.H{"data": tahunAjaran})
}

// UpdateTahunAjaran godoc
// @Summary Memperbarui data tahun ajaran.
// @Description memperbaharui data tahun ajaran berdasarkan id.
// @Tags Sekolah -> Tahun Ajaran
// @Produce json
// @Param id path string true "Tahun ajaran id"
// @Param Body body UpdateInputTahunAjaran true "the body to update tahun ajaran"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} sekolah.TahunAjaran
// @Router /sekolah/tahun-ajaran/{id} [patch]
func UpdateTahunAjaran(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var tahunAjaran sekolah.TahunAjaran
	if err := db.Where("id = ?", c.Param("id")).First(&tahunAjaran).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var InputTahunAjaran UpdateInputTahunAjaran
	errInput := c.ShouldBindJSON(&InputTahunAjaran)

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

	var updatedInput sekolah.TahunAjaran
	updatedInput.Nama = InputTahunAjaran.Nama
	updatedInput.Semester = InputTahunAjaran.Semester
	updatedInput.UpdatedAt = time.Now()

	db.Model(&tahunAjaran).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": tahunAjaran})
}

// DeleteTahunAjaran godoc
// @Summary Menghapus data tahun ajaran.
// @Description Menghapus Sebuah data tahun ajaran  berdasarkan id.
// @Tags Sekolah -> Tahun Ajaran
// @Produce json
// @Param id path string true "tahun ajaran id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]string
// @Router /sekolah/tahun-ajaran/{id} [delete]
func DeleteTahunAjaran(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var tahunAjaran sekolah.TahunAjaran

	if err := db.Where("id = ?", c.Param("id")).First(&tahunAjaran).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&tahunAjaran)

	c.JSON(http.StatusOK, gin.H{"data": "Berhasil Menghapus Data"})
}
