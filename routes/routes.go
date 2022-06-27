package routes

import (
	"sispendik-api/controllers"
	"sispendik-api/controllers/sekolahController"
	"sispendik-api/controllers/siswaController"
	"sispendik-api/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/login", controllers.Login)
	r.POST("/daftar-sekolah", controllers.RegisterSekolah)
	r.POST("/daftar-siswa", controllers.RegisterSiswa)
	r.PATCH("/ganti-password/:username", controllers.GantiPassword)

	r.GET("/roles", controllers.RoleUserLogin)
	r.GET("/users", controllers.GetALLUsers)

	r.GET("/sekolah/:id", sekolahController.GetSekolahbyId)
	r.GET("/sekolah/user/:id", sekolahController.GetSekolahByIdUserid)
	r.GET("/sekolah/pendaftaran-siswa-aktif", sekolahController.GetPendaftaranSiswaAktif)
	r.GET("/sekolah/pendaftaran-pendidik-aktif", sekolahController.GetPendaftaranPendidikAktif)

	r.GET("/sekolah/:id/angkatan", sekolahController.GetAngkatanbyIdSekolah)
	r.GET("/sekolah/angkatan/:id", sekolahController.GetAngkatanbyId)

	r.GET("/sekolah/:id/tahun-ajaran", sekolahController.GetTahunAjaranbyIdSekolah)
	r.GET("/sekolah/tahun-ajaran/:id", sekolahController.GetTahunAjaranbyId)

	r.GET("/siswa/:id", siswaController.GetSiswaById)
	r.GET("/siswa/user/:id", siswaController.GetSiswaByIdUserid)

	r.GET("/sekolah/tahun-ajaran/:id/kelas", sekolahController.GetKelasbyIdTahunAjaran)
	r.GET("/sekolah/kelas/:id", sekolahController.GetKelasbyId)

	r.GET("/siswa/prestasi/:id", siswaController.GetPrestasiById)
	r.GET("/siswa/:id/prestasi", siswaController.GetPrestasiByIdSiswa)

	r.GET("/siswa/asal-sekolah/:id", siswaController.GetAsalSekolahById)
	r.GET("/siswa/:id/asal-sekolah", siswaController.GetAsalSekolahByIdSiswa)

	r.GET("/siswa/beasiswa/:id", siswaController.GetBeasiswaById)
	r.GET("/siswa/:id/beasiswa", siswaController.GetBeasiswaByIdSiswa)

	groupSekolahRoute := r.Group("/sekolah")
	// sekolah
	groupSekolahRoute.Use(middlewares.JwtAuthMiddleware())
	groupSekolahRoute.PATCH("/:id/buka-pendaftaran-siswa", sekolahController.BukaPendaftaranSiswaBySekolahID)
	groupSekolahRoute.PATCH("/:id/buka-pendaftaran-pendidik", sekolahController.BukaPendaftaranPendidikBySekolahID)
	groupSekolahRoute.PATCH("/:id/tutup-pendaftaran-siswa", sekolahController.TutupPendaftaranSiswaBySekolahID)
	groupSekolahRoute.PATCH("/:id/tutup-pendaftaran-pendidik", sekolahController.TutupPendaftaranPendidikBySekolahID)
	groupSekolahRoute.PATCH("/:id/kode-verifikasi-pendaftaran-siswa", sekolahController.UpdateKodeVerifikasiPendaftaranSiswa)
	groupSekolahRoute.PATCH("/:id/kode-verifikasi-pendaftaran-pendidik", sekolahController.UpdateKodeVerifikasiPendaftaranPendidik)
	groupSekolahRoute.PATCH("/:id/aktivasi-menu-siswa", sekolahController.AktivasiMenuSiswaBySekolahID)
	groupSekolahRoute.PATCH("/:id/aktivasi-menu-pendidik", sekolahController.AktivasiMenuPendidikBySekolahID)
	groupSekolahRoute.PATCH("/:id/non-aktif-menu-siswa", sekolahController.NonAktifMenuSiswaBySekolahID)
	groupSekolahRoute.PATCH("/:id/non-aktif-menu-pendidik", sekolahController.NonAktifMenuPendidikBySekolahID)
	groupSekolahRoute.DELETE("/:id", sekolahController.DeleteSekolah)
	groupSekolahRoute.POST("/angkatan", sekolahController.CreateAngkatan)
	groupSekolahRoute.PATCH("/angkatan/:id", sekolahController.UpdateAngkatan)
	groupSekolahRoute.DELETE("/angkatan/:id", sekolahController.DeleteAngkatan)
	groupSekolahRoute.POST("/tahun-ajaran", sekolahController.CreateTahunAjaran)
	groupSekolahRoute.PATCH("/tahun-ajaran/:id", sekolahController.UpdateTahunAjaran)
	groupSekolahRoute.DELETE("/tahun-ajaran/:id", sekolahController.DeleteTahunAjaran)
	groupSekolahRoute.POST("/kelas", sekolahController.CreateKelas)
	groupSekolahRoute.PATCH("/kelas/:id", sekolahController.UpdateKelas)
	groupSekolahRoute.DELETE("/kelas/:id", sekolahController.DeleteKelas)

	groupSiswaRoute := r.Group("/siswa")
	groupSiswaRoute.Use(middlewares.JwtAuthMiddleware())
	groupSiswaRoute.PATCH("/:id", siswaController.UpdateSiswa)
	groupSiswaRoute.DELETE("/:id", siswaController.DeleteSiswa)
	groupSiswaRoute.PATCH("/beasiswa/:id", siswaController.UpdateBeasiswa)
	groupSiswaRoute.PATCH("/asal-sekolah/:id", siswaController.UpdateAsalSekolah)
	groupSiswaRoute.PATCH("/prestasi/:id", siswaController.UpdatePrestasi)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
