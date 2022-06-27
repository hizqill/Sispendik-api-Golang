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

type InputAngakatan struct {
	SekolahID uint   `json:"sekolah_id" binding:"required,number"`
	Nama      string `json:"nama" binding:"required"`
}

type UpdateInputAngakatan struct {
	Nama string `json:"nama" binding:"required"`
}

// GetAngkatanbyIdSekolah godoc
// @Summary mendapatkan data angkatan berdasarkan ID Sekolah.
// @Description Mendapatkan semua data angkatan Berdasarkan ID Sekolah.
// @Tags Sekolah -> Angkatan
// @Produce json
// @Param id path string true "Sekolah id"
// @Success 200 {object} []sekolah.Angkatan
// @Router /sekolah/{id}/angkatan [get]
func GetAngkatanbyIdSekolah(c *gin.Context) { // Get model if exist
	var angkatan []sekolah.Angkatan

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("sekolah_id  = ?", c.Param("id")).Find(&angkatan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": angkatan})
}

// GetAngkatanbyId godoc
// @Summary Mendapatkan data angkatan.
// @Description Mendapatkan data angkatan berdasarkan id.
// @Tags Sekolah -> Angkatan
// @Produce json
// @Param id path string true "Angkatan id"
// @Success 200 {object} sekolah.Angkatan
// @Router /sekolah/angkatan/{id} [get]
func GetAngkatanbyId(c *gin.Context) { // Get model if exist
	var angkatan sekolah.Angkatan

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&angkatan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": angkatan})
}

// CreateAngkatan godoc
// @Summary Membuat data Angkatan barun.
// @Description Membuat data Angkatan baru.
// @Tags Sekolah -> Angkatan
// @Param Body body InputAngakatan true "the body to create a new Angkatan"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} sekolah.Angkatan
// @Router /sekolah/angkatan [post]
func CreateAngkatan(c *gin.Context) {
	// Validate input
	var InputAngakatan InputAngakatan
	errInput := c.ShouldBindJSON(&InputAngakatan)

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
	angkatan := sekolah.Angkatan{
		SekolahID: InputAngakatan.SekolahID,
		Nama:      InputAngakatan.Nama,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&angkatan)

	c.JSON(http.StatusOK, gin.H{"data": angkatan})
}

// UpdateAngkatan godoc
// @Summary Memperbarui data Angkatan.
// @Description memperbaharui data angkatan berdasarkan id.
// @Tags Sekolah -> Angkatan
// @Produce json
// @Param id path string true "Angkatan id"
// @Param Body body UpdateInputAngakatan true "the body to update Angkatan"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} sekolah.Angkatan
// @Router /sekolah/angkatan/{id} [patch]
func UpdateAngkatan(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var angakatan sekolah.Angkatan
	if err := db.Where("id = ?", c.Param("id")).First(&angakatan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var inputAngakatan UpdateInputAngakatan
	errInput := c.ShouldBindJSON(&inputAngakatan)

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

	var updatedInput sekolah.Angkatan
	updatedInput.Nama = inputAngakatan.Nama
	updatedInput.UpdatedAt = time.Now()

	db.Model(&angakatan).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": angakatan})
}

// DeleteAngkatan godoc
// @Summary Menghapus data angkatan.
// @Description Menghapus Sebuah data angkatan  berdasarkan id.
// @Tags Sekolah -> Angkatan
// @Produce json
// @Param id path string true "Angkatan id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]string
// @Router /sekolah/angkatan/{id} [delete]
func DeleteAngkatan(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var angakatan sekolah.Angkatan

	if err := db.Where("id = ?", c.Param("id")).First(&angakatan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&angakatan)

	c.JSON(http.StatusOK, gin.H{"data": "Berhasil Menghapus Data"})
}
