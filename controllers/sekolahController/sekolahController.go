package sekolahController

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

type UpdateInputKodeVerifikasiSiswa struct {
	KodeVerifikasiPendaftaranSiswa string `json:"kode_verifikasi_pendaftaran_siswa"`
}
type UpdateInputKodeVerifikasiPendidik struct {
	KodeVerifikasiPendaftaranPendidik string `json:"kode_verifikasi_pendaftaran_pendidik"`
}

// GetSekolahByIdUserid godoc
// @Summary Mendapatkan data sekolah berdasarkan user ID.
// @Description Mendapatkan data sekolah berdasarkan user ID.
// @Tags Sekolah
// @Produce json
// @Param id path string true "User id"
// @Success 200 {object} s.Sekolah
// @Router /sekolah/user/{id} [get]
func GetSekolahByIdUserid(c *gin.Context) {
	// Get model if exista
	var sekolah s.Sekolah

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("user_id = ?", c.Param("id")).First(&sekolah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sekolah})
}

// GetSekolahbyId godoc
// @Summary Mendapatkan data sekolah.
// @Description Mendapatkan data sekolah berdasarkan id.
// @Tags Sekolah
// @Produce json
// @Param id path string true "Sekolah id"
// @Success 200 {object} s.Sekolah
// @Router /sekolah/{id} [get]
func GetSekolahbyId(c *gin.Context) { // Get model if exist
	var sekolah s.Sekolah

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&sekolah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sekolah})
}

// GetPendaftaranSiswaAktif godoc
// @Summary mendapatkan data sekolah yang membuka jalur pendaftaran siswa.
// @Description Mendapatkan semua ata sekolah yang membuka jalur pendaftaran siswa.
// @Tags Sekolah
// @Produce json
// @Success 200 {object} []s.Sekolah
// @Router /sekolah/pendaftaran-siswa-aktif [get]
func GetPendaftaranSiswaAktif(c *gin.Context) {
	// Get model if exist
	var sekolah []s.Sekolah

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("pendaftaran_siswa  = 'aktif'").Find(&sekolah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sekolah})
}

// GetPendaftaranPendidikAktif godoc
// @Summary mendapatkan data sekolah yang membuka jalur pendaftaran pendidik.
// @Description Mendapatkan semua ata sekolah yang membuka jalur pendaftaran pendidik.
// @Tags Sekolah
// @Produce json
// @Success 200 {object} []s.Sekolah
// @Router /sekolah/pendaftaran-pendidik-aktif [get]
func GetPendaftaranPendidikAktif(c *gin.Context) {
	// Get model if exist
	var sekolah []s.Sekolah

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("pendaftaran_pendidik  = 'aktif'").Find(&sekolah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sekolah})
}

// BukaPendaftaranSiswaBySekolahID godoc
// @Summary Buka pendafatran siswa berdasarkan sekolah id.
// @Description Buka pendafatran siswa berdasarkan sekolah id.
// @Tags Sekolah
// @Produce json
// @Param id path string true "Sekolah id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} s.Sekolah
// @Router /sekolah/{id}/buka-pendaftaran-siswa [patch]
func BukaPendaftaranSiswaBySekolahID(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var sekolah s.Sekolah

	if err := db.Where("id = ?", c.Param("id")).First(&sekolah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updatedInput s.Sekolah
	updatedInput.PendaftaranSiswa = "aktif"

	db.Model(&sekolah).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{
		"Response_message": "Pendaftaran siswa berhasil dibuka",
		"data":             sekolah,
	})
}

// BukaPendaftaranPendidikBySekolahID godoc
// @Summary Buka pendafatran pendidik berdasarkan sekolah id.
// @Description Buka pendafatran pendidik berdasarkan sekolah id.
// @Tags Sekolah
// @Produce json
// @Param id path string true "Sekolah id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} s.Sekolah
// @Router /sekolah/{id}/buka-pendaftaran-pendidik [patch]
func BukaPendaftaranPendidikBySekolahID(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var sekolah s.Sekolah

	if err := db.Where("id = ?", c.Param("id")).First(&sekolah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updatedInput s.Sekolah
	updatedInput.PendaftaranPendidik = "aktif"

	db.Model(&sekolah).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{
		"Response_message": "Pendaftaran pendidik berhasil dibuka",
		"data":             sekolah,
	})
}

// TutupPendaftaranSiswaBySekolahID godoc
// @Summary Tutup pendafatran siswa berdasarkan sekolah id.
// @Description Tutup pendafatran siswa berdasarkan sekolah id.
// @Tags Sekolah
// @Produce json
// @Param id path string true "Sekolah id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} s.Sekolah
// @Router /sekolah/{id}/tutup-pendaftaran-siswa [patch]
func TutupPendaftaranSiswaBySekolahID(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var sekolah s.Sekolah

	if err := db.Where("id = ?", c.Param("id")).First(&sekolah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updatedInput s.Sekolah
	updatedInput.PendaftaranSiswa = "nonaktif"

	db.Model(&sekolah).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{
		"Response_message": "Pendaftaran siswa berhasil ditutup",
		"data":             sekolah,
	})
}

// TutupPendaftaranPendidikBySekolahID godoc
// @Summary Tutup pendafatran pendidik berdasarkan sekolah id.
// @Description Tutup pendafatran pendidik berdasarkan sekolah id.
// @Tags Sekolah
// @Produce json
// @Param id path string true "Sekolah id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} s.Sekolah
// @Router /sekolah/{id}/tutup-pendaftaran-pendidik [patch]
func TutupPendaftaranPendidikBySekolahID(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var sekolah s.Sekolah

	if err := db.Where("id = ?", c.Param("id")).First(&sekolah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updatedInput s.Sekolah
	updatedInput.PendaftaranPendidik = "nonaktif"

	db.Model(&sekolah).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{
		"Response_message": "Pendaftaran pendidik berhasil ditutup",
		"data":             sekolah,
	})
}

// AktivasiMenuSiswaBySekolahID godoc
// @Summary Aktivasi Menu pengisian/update biodata siswa Siswa berdasarkan sekolah id.
// @Description Aktivasi Menu pengisian/update biodata siswa Siswa berdasarkan sekolah id.
// @Tags Sekolah
// @Produce json
// @Param id path string true "Sekolah id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} s.Sekolah
// @Router /sekolah/{id}/aktivasi-menu-siswa [patch]
func AktivasiMenuSiswaBySekolahID(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var sekolah s.Sekolah

	if err := db.Where("id = ?", c.Param("id")).First(&sekolah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updatedInput s.Sekolah
	updatedInput.StatusMenuSiswa = "aktif"

	db.Model(&sekolah).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{
		"Response_message": "Menu pengisian/update biodata siswa berhasil dibuka",
		"data":             sekolah,
	})
}

// AktivasiMenuSiswaBySekolahID godoc
// @Summary Aktivasi Menu pengisian/update biodata siswa Siswa berdasarkan sekolah id.
// @Description Aktivasi Menu pengisian/update biodata siswa Siswa berdasarkan sekolah id.
// @Tags Sekolah
// @Produce json
// @Param id path string true "Sekolah id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} s.Sekolah
// @Router /sekolah/{id}/aktivasi-menu-pendidik [patch]
func AktivasiMenuPendidikBySekolahID(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var sekolah s.Sekolah

	if err := db.Where("id = ?", c.Param("id")).First(&sekolah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updatedInput s.Sekolah
	updatedInput.StatusMenuPendidik = "aktif"

	db.Model(&sekolah).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{
		"Response_message": "Menu pengisian/update biodata siswa berhasil dibuka",
		"data":             sekolah,
	})
}

// NonAktifMenuSiswaBySekolahID godoc
// @Summary Non aktif Menu pengisian/update biodata siswa Siswa berdasarkan sekolah id.
// @Description Non aktif Menu pengisian/update biodata siswa Siswa berdasarkan sekolah id.
// @Tags Sekolah
// @Produce json
// @Param id path string true "Sekolah id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} s.Sekolah
// @Router /sekolah/{id}/non-aktif-menu-siswa [patch]
func NonAktifMenuSiswaBySekolahID(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var sekolah s.Sekolah

	if err := db.Where("id = ?", c.Param("id")).First(&sekolah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updatedInput s.Sekolah
	updatedInput.StatusMenuSiswa = "nonaktif"

	db.Model(&sekolah).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{
		"Response_message": "Menu pengisian/update biodata siswa berhasil dinonaktifkan",
		"data":             sekolah,
	})
}

// NonAktifMenuPendidikBySekolahID godoc
// @Summary non aktif Menu pengisian/update biodata siswa Siswa berdasarkan sekolah id.
// @Description non aktif Menu pengisian/update biodata siswa Siswa berdasarkan sekolah id.
// @Tags Sekolah
// @Produce json
// @Param id path string true "Sekolah id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} s.Sekolah
// @Router /sekolah/{id}/non-aktif-menu-pendidik [patch]
func NonAktifMenuPendidikBySekolahID(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var sekolah s.Sekolah

	if err := db.Where("id = ?", c.Param("id")).First(&sekolah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updatedInput s.Sekolah
	updatedInput.StatusMenuPendidik = "nonaktif"

	db.Model(&sekolah).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{
		"Response_message": "Menu pengisian/update biodata siswa berhasil dinonaktifkan",
		"data":             sekolah,
	})
}

// UpdateKodeVerifikasiPendaftaranSiswa godoc
// @Summary Memperbarui kode verifikasi pendaftaran siswa.
// @Description memperbaharui kode verifikasi pendaftaran siswa berdasarkan id sekolah.
// @Tags Sekolah
// @Produce json
// @Param id path string true "Tahun ajaran id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param Body body UpdateInputKodeVerifikasiSiswa true "the body to update kode verifikasi pendaftaran siswa"
// @Success 200 {object} s.Sekolah
// @Router /sekolah/{id}/kode-verifikasi-pendaftaran-siswa [patch]
func UpdateKodeVerifikasiPendaftaranSiswa(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var sekolah s.Sekolah
	if err := db.Where("id = ?", c.Param("id")).First(&sekolah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var InputKodeVerifikasiSiswa UpdateInputKodeVerifikasiSiswa
	errInput := c.ShouldBindJSON(&InputKodeVerifikasiSiswa)

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

	var updatedInput s.Sekolah
	updatedInput.KodeVerifikasiPendaftaranSiswa = InputKodeVerifikasiSiswa.KodeVerifikasiPendaftaranSiswa
	updatedInput.UpdatedAt = time.Now()

	db.Model(&sekolah).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{
		"Response_Messange":                 "Berhasil memperbarui kode verifikasi pendaftran siswa",
		"kode_verifikasi_pendaftaran_siswa": InputKodeVerifikasiSiswa.KodeVerifikasiPendaftaranSiswa,
		"data":                              sekolah,
	})
}

// UpdateKodeVerifikasiPendaftaranPendidik godoc
// @Summary Memperbarui kode verifikasi pendaftaran pendidik.
// @Description memperbaharui kode verifikasi pendaftaran pendidik berdasarkan id sekolah.
// @Tags Sekolah
// @Produce json
// @Param id path string true "Tahun ajaran id"
// @Param Body body UpdateInputKodeVerifikasiPendidik true "the body to update kode verifikasi pendaftaran pendidik"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} s.Sekolah
// @Router /sekolah/{id}/kode-verifikasi-pendaftaran-pendidik [patch]
func UpdateKodeVerifikasiPendaftaranPendidik(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var sekolah s.Sekolah
	if err := db.Where("id = ?", c.Param("id")).First(&sekolah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var InputKodeVerifikasiPendidik UpdateInputKodeVerifikasiPendidik
	errInput := c.ShouldBindJSON(&InputKodeVerifikasiPendidik)

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

	var updatedInput s.Sekolah
	updatedInput.KodeVerifikasiPendaftaranPendidik = InputKodeVerifikasiPendidik.KodeVerifikasiPendaftaranPendidik
	updatedInput.UpdatedAt = time.Now()

	db.Model(&sekolah).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{
		"Response_Messange":                 "Berhasil memperbarui kode verifikasi pendaftran siswa",
		"kode_verifikasi_pendaftaran_siswa": InputKodeVerifikasiPendidik.KodeVerifikasiPendaftaranPendidik,
		"data":                              sekolah,
	})
}

// DeleteSekolah godoc
// @Summary Menghapus data sekolah.
// @Description Menghapus semua data sekolah berdasarkan id. mulai dari akun, .
// @Tags Siswa
// @Produce json
// @Param id path string true "Sekolah id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]string
// @Router /sekolah/{id} [delete]
func DeleteSekolah(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var sekolah s.Sekolah

	if err := db.Where("id = ?", c.Param("id")).First(&sekolah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var user models.Users
	if err := db.Where("id = ?", sekolah.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record user not found!"})
		return
	}

	db.Delete(&user)
	db.Delete(&sekolah)

	c.JSON(http.StatusOK, gin.H{"data": "Berhasil Menghapus Data"})
}
