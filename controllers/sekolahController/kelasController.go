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

type InputKelas struct {
	SekolahID     uint   `json:"sekolah_id" binding:"required,number"`
	TahunAjaranID uint   `json:"tahun_ajaran_id" binding:"required,number"`
	Nama          string `json:"nama" binding:"required"`
}

type UpdateInputKelas struct {
	Nama string `json:"nama" binding:"required"`
}

// GetKelasbyIdTahunAjaran godoc
// @Summary mendapatkan data kelas berdasarkan ID Tahun Ajaran.
// @Description Mendapatkan semua data kelas Berdasarkan ID Tahun Ajaran.
// @Tags Sekolah -> Kelas
// @Produce json
// @Param id path string true "Tahun ajaran id  id"
// @Success 200 {object} []sekolah.Kelas
// @Router /sekolah/tahun-ajaran/{id}/kelas [get]
func GetKelasbyIdTahunAjaran(c *gin.Context) {
	// Get model if exist
	var kelas []sekolah.Kelas

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("tahun_ajaran_id   = ?", c.Param("id")).Find(&kelas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": kelas})
}

// GeKelasbyId godoc
// @Summary Mendapatkan data kelas.
// @Description Mendapatkan data kelas berdasarkan id.
// @Tags Sekolah -> Kelas
// @Produce json
// @Param id path string true "kelas id"
// @Success 200 {object} sekolah.Kelas
// @Router /sekolah/kelas/{id} [get]
func GetKelasbyId(c *gin.Context) { // Get model if exist
	var kelas sekolah.Kelas

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&kelas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": kelas})
}

// CreateKelas godoc
// @Summary Membuat data Kelas barun.
// @Description Membuat data Kelas baru.
// @Tags Sekolah -> Kelas
// @Param Body body InputKelas true "the body to create a new Kelas"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} sekolah.Kelas
// @Router /sekolah/kelas [post]
func CreateKelas(c *gin.Context) {
	// Validate input
	var InputKelas InputKelas
	errInput := c.ShouldBindJSON(&InputKelas)

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

	// Create Kelas
	kelas := sekolah.Kelas{
		SekolahID:     InputKelas.SekolahID,
		TahunAjaranID: InputKelas.TahunAjaranID,
		Nama:          InputKelas.Nama,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&kelas)

	c.JSON(http.StatusOK, gin.H{"data": kelas})
}

// UpdateKelas godoc
// @Summary Memperbarui data Kelas.
// @Description memperbaharui data Kelas berdasarkan id.
// @Tags Sekolah -> Kelas
// @Produce json
// @Param id path string true "Kelas id"
// @Param Body body UpdateInputKelas true "the body to update Kelas"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} sekolah.Kelas
// @Router /sekolah/kelas/{id} [patch]
func UpdateKelas(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var kelas sekolah.Kelas
	if err := db.Where("id = ?", c.Param("id")).First(&kelas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var inputKelas UpdateInputKelas
	errInput := c.ShouldBindJSON(&inputKelas)

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

	var updatedInput sekolah.Kelas
	updatedInput.Nama = inputKelas.Nama
	updatedInput.UpdatedAt = time.Now()

	db.Model(&kelas).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": kelas})
}

// DeleteKelas godoc
// @Summary Menghapus data Kelas.
// @Description Menghapus Sebuah data Kelas  berdasarkan id.
// @Tags Sekolah -> Kelas
// @Produce json
// @Param id path string true "Kelas id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]string
// @Router /sekolah/kelas/{id} [delete]
func DeleteKelas(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var kelas sekolah.Kelas

	if err := db.Where("id = ?", c.Param("id")).First(&kelas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&kelas)

	c.JSON(http.StatusOK, gin.H{"data": "Berhasil Menghapus Data"})
}
