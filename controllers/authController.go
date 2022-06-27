package controllers

import (
	"fmt"
	"net/http"
	"sispendik-api/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInputSekolah struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nama     string `json:"nama" binding:"required"`
}

type InputUpdatePassword struct {
	Password string `json:"password" binding:"required"`
}

type RegisterInputSiswa struct {
	Username                  string `json:"username" binding:"required"`
	Password                  string `json:"password" binding:"required"`
	Nama                      string `json:"nama" binding:"required"`
	Nisn                      string `json:"nisn" binding:"required"`
	SekolahID                 uint   `json:"sekolah_id" binding:"required"`
	KodeVerifikasiPendaftaran string `json:"kode_verifikasi_pendaftaran" binding:"required"`
}

func createBeasiswaSiswa(ch chan *gorm.DB, c *gin.Context, siswaId uint) {
	db := c.MustGet("db").(*gorm.DB)
	beasiswaSiswa := models.BeasiswaSiswas{}
	beasiswaSiswa.SiswaId = siswaId
	ch <- db.Create(&beasiswaSiswa)
}

func createAsalSekolahSiswa(ch chan *gorm.DB, c *gin.Context, siswaId uint) {
	db := c.MustGet("db").(*gorm.DB)
	asalSekolahSiswa := models.AsalsekolahSiswas{}
	asalSekolahSiswa.SiswaID = siswaId
	ch <- db.Create(&asalSekolahSiswa)
}

func createPrestasiSiswa(ch chan *gorm.DB, c *gin.Context, siswaId uint) {
	db := c.MustGet("db").(*gorm.DB)
	prestasiSiswa := models.PrestasiSiswas{}
	prestasiSiswa.SiswaID = siswaId
	ch <- db.Create(&prestasiSiswa)
}

// LoginUser godoc
// @Summary Login as as user.
// @Description Logging in to get jwt token to access admin or user api by roles.
// @Tags Auth
// @Param Body body LoginInput true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// errInput := c.ShouldBindJSON(&input)

	// if errInput != nil {
	// 	errorMessanges := []string{}
	// 	for _, e := range errInput.(validator.ValidationErrors) {
	// 		errorMessange := fmt.Sprintf("Error on field %s, candition: %s", e.Field(), e.ActualTag())
	// 		errorMessanges = append(errorMessanges, errorMessange)
	// 	}

	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": errorMessanges,
	// 	})
	// 	return
	// }

	u := models.Users{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheck(u.Username, u.Password, db)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	var user models.Users

	if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username not found!"})
		return
	}

	id := strconv.Itoa(int(user.ID))
	role := strconv.Itoa(int(user.RoleID))
	resUser := map[string]string{
		"id":       id,
		"role_id":  role,
		"nama":     user.Nama,
		"username": user.Username,
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "user": resUser, "token": token})

}

// RegisterSekolah godoc
// @Summary Daftar Akun Sekolah (akun, kontrol sekolah -> menu-menu, pendaftaran pendidik/siswa).
// @Description Daftar Akun Sekolah hanya oleh super admin.
// @Tags Auth
// @Param Body body RegisterInputSekolah true "body unutk daftar akun Sekolah"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /daftar-sekolah [post]
func RegisterSekolah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var inputSekolah RegisterInputSekolah

	errInput := c.ShouldBindJSON(&inputSekolah)

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

	// if err := c.ShouldBindJSON(&inputSekolah); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	u := models.Users{}
	u.Nama = inputSekolah.Nama
	u.RoleID = 2
	u.Username = inputSekolah.Username
	u.Password = inputSekolah.Password
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	_, err := u.SaveUser(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userid := uint(u.ID)

	s := models.Sekolah{}
	s.Nama = inputSekolah.Nama
	s.UserID = userid
	s.PendaftaranSiswa = "nonaktif"
	s.PendaftaranPendidik = "nonaktif"
	s.StatusMenuSiswa = "nonaktif"
	s.StatusMenuPendidik = "nonaktif"
	db.Create(&s)

	id := strconv.Itoa(int(u.ID))
	role := strconv.Itoa(int(u.RoleID))
	user := map[string]string{
		"id":       id,
		"role_id":  role,
		"nama":     u.Nama,
		"username": u.Username,
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "user": user})

}

// RegisterSiswa godoc
// @Summary Daftar Akun Siswa (akun, biodata, asal sekolah, prestasi, beasiswa).
// @Description Daftar Akun Siswa table yang terdaftar ada 4 yaitu table users sebagai akun, table siswa sebagai biodata, table prestasi sebagai manajemen prestasi siswa, table asal sekolah sebagai manajemen asal sekolah siswa.
// @Tags Auth
// @Param Body body RegisterInputSiswa true "body untuk daftar akun siswa"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /daftar-siswa [post]
func RegisterSiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var inputSiswa RegisterInputSiswa

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

	var sekolah models.Sekolah

	if err := db.Where("id = ?", inputSiswa.SekolahID).First(&sekolah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sekolah_id not found!"})
		return
	}

	if sekolah.KodeVerifikasiPendaftaranSiswa != inputSiswa.KodeVerifikasiPendaftaran {
		c.JSON(http.StatusBadRequest, gin.H{"error": "kode verifikasi pendaftaran tidak sesuai"})
		return
	}

	u := models.Users{}
	u.Nama = inputSiswa.Nama
	u.RoleID = 4
	u.Username = inputSiswa.Username
	u.Password = inputSiswa.Password
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	_, err := u.SaveUser(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sekolahid := uint(inputSiswa.SekolahID)

	s := models.Siswa{}
	s.UserID = u.ID
	s.SekolahID = sekolahid
	s.Nisn = inputSiswa.Nisn
	s.Nama = inputSiswa.Nama
	db.Create(&s)

	beasiswaSiswaChannel := make(chan *gorm.DB)
	go createBeasiswaSiswa(beasiswaSiswaChannel, c, s.ID)
	asalSekolahSiswaChannel := make(chan *gorm.DB)
	go createAsalSekolahSiswa(asalSekolahSiswaChannel, c, s.ID)
	prestasiSiswaChannel := make(chan *gorm.DB)
	go createPrestasiSiswa(prestasiSiswaChannel, c, s.ID)

	id := strconv.Itoa(int(u.ID))
	role := strconv.Itoa(int(u.RoleID))
	idSekolah := strconv.Itoa(int(s.SekolahID))
	user := map[string]string{
		"id":         id,
		"role_id":    role,
		"sekolah_id": idSekolah,
		"nisn":       inputSiswa.Nisn,
		"nama":       u.Nama,
		"username":   u.Username,
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "user": user})

}

// GantiPassword godoc
// @Summary Menggati Password akun.
// @Description  Menggati Password akun berdasarkan username.
// @Tags Auth
// @Produce json
// @Param username path string true "Username"
// @Param Body body InputUpdatePassword true "body untuk ganti password"
// @Success 200 {object} models.Users
// @Router /ganti-password/{username} [patch]
func GantiPassword(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var user models.Users
	var inputPassword InputUpdatePassword

	if err := db.Where("username = ?", c.Param("username")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input

	errInput := c.ShouldBindJSON(&inputPassword)
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

	var updatedInput models.Users
	updatedInput.Password = inputPassword.Password
	updatedInput.UpdatedAt = time.Now()

	db.Model(&user).Updates(updatedInput)

	id := strconv.Itoa(int(user.ID))
	role := strconv.Itoa(int(user.RoleID))
	res := map[string]string{
		"id":       id,
		"role_id":  role,
		"nama":     user.Nama,
		"username": user.Username,
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil Memperbaharui password", "user": res})

}
