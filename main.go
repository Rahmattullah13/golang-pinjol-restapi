package main

import (
	"golang-pinjol/config"
	"golang-pinjol/controller"
	"golang-pinjol/middleware"
	"golang-pinjol/repository"
	"golang-pinjol/services"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()

	// Repository
	nasabahRepository   repository.NasabahRepository     = repository.NewNasabahRepository(db)
	pekerjaanRepository repository.RepositoryNasabahJobs = repository.NewRepositoryNasabahJobs(db)

	// Service
	jwtService              services.JwtService              = services.NewJwtService()
	authService             services.AuthenticationService   = services.NewAuthenticationService(nasabahRepository)
	nasabahService          services.NasabahServices         = services.NewNasabahService(nasabahRepository)
	pekerjaanNasabahService services.PekerjaanNasabahService = services.NewPekerjaanNasabahService(pekerjaanRepository)

	// Controller
	authController             controller.AuthController             = controller.NewAuthController(authService, jwtService)
	nasabahController          controller.NasabahController          = controller.NewNasabahController(nasabahService, jwtService)
	uploadFileController       controller.UploadFileController       = controller.NewUploadFileController(jwtService, db)
	pekerjaanNasabahController controller.PekerjaanNasabahController = controller.NewPekerjaanNasabahController(pekerjaanNasabahService, jwtService)
)

func main() {
	defer config.CloseDB(db)
	log.Println(db)

	r := gin.Default()

	auth := r.Group("app/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	nasabah := r.Group("app/nasabah", middleware.Authorize(jwtService))
	{
		nasabah.PUT("/update", nasabahController.UpdateNasabahController)
		nasabah.GET("/profile", nasabahController.ProfileNasabahController)
	}

	documentNasabah := r.Group("app/document", middleware.Authorize(jwtService))
	{
		documentNasabah.PUT("/upload/:nasabah_id", uploadFileController.UploadFile)
	}

	pekerjaanNasabah := r.Group("app/jobs", middleware.Authorize(jwtService))
	{
		pekerjaanNasabah.POST("/addJobs", pekerjaanNasabahController.AddNasabahJobsController)
		pekerjaanNasabah.PUT("/:id", pekerjaanNasabahController.NasabahUpdateJobsController)
		pekerjaanNasabah.GET("/:id", pekerjaanNasabahController.SearchNasabahJobsByIdController)
		pekerjaanNasabah.DELETE("/:id", pekerjaanNasabahController.DeleteNasabahJobsController)
	}

	r.Run("localhost:3000")
}
