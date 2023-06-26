package main

import (
	"golang-pinjol/config"
	"golang-pinjol/controller"
	"golang-pinjol/repository"
	"golang-pinjol/services"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()

	// Repository
	nasabahRepository repository.NasabahRepository = repository.NewNasabahRepository(db)

	// Service
	jwtService  services.JwtService            = services.NewJwtService()
	authService services.AuthenticationService = services.NewAuthenticationService(nasabahRepository)

	// Controller
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
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

	r.Run("localhost:3000")
}
