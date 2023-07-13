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
	nasabahRepository   repository.NasabahRepository        = repository.NewNasabahRepository(db)
	pekerjaanRepository repository.RepositoryNasabahJobs    = repository.NewRepositoryNasabahJobs(db)
	loanRepository      repository.LoansRepository          = repository.NewLoansRepository(db)
	paymentRepository   repository.PaymentRepository        = repository.NewPaymentRepository(db)
	historyRepository   repository.HistoryPaymentRepository = repository.NewHistoryPaymentRepository(db)

	// Service
	jwtService              services.JwtService              = services.NewJwtService()
	authService             services.AuthenticationService   = services.NewAuthenticationService(nasabahRepository)
	nasabahService          services.NasabahServices         = services.NewNasabahService(nasabahRepository)
	pekerjaanNasabahService services.PekerjaanNasabahService = services.NewPekerjaanNasabahService(pekerjaanRepository)
	loanService             services.LoanService             = services.NewLoanService(loanRepository, nasabahRepository)
	paymentService          services.PaymentService          = services.NewPaymentService(paymentRepository)
	historyService          services.HistoryPaymentService   = services.NewHistoryPaymentService(historyRepository)

	// Controller
	authController             controller.AuthController             = controller.NewAuthController(authService, jwtService)
	nasabahController          controller.NasabahController          = controller.NewNasabahController(nasabahService, jwtService)
	uploadFileController       controller.UploadFileController       = controller.NewUploadFileController(jwtService, db)
	pekerjaanNasabahController controller.PekerjaanNasabahController = controller.NewPekerjaanNasabahController(pekerjaanNasabahService, jwtService)
	loanController             controller.LoanController             = controller.NewLoanController(loanService, jwtService)
	paymentController          controller.PaymentController          = controller.NewPaymentController(paymentService, jwtService)
	historyPaymentController   controller.HistoryPaymentController   = controller.NewHistoryPaymentController(historyService, jwtService)
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

	loanNasabah := r.Group("app/loans", middleware.Authorize(jwtService))
	{
		loanNasabah.POST("/loan", loanController.CreateLoanController)
		loanNasabah.PUT("/:id", loanController.UpdateLoanController)
		loanNasabah.GET("/:id", loanController.SearchLoanByIdController)
		loanNasabah.PUT("/verification/:id", loanController.UpdateStatusApprovalController)
		loanNasabah.DELETE("/:id", loanController.DeleteLoanController)
	}

	paymentsNasabah := r.Group("app/payments", middleware.Authorize(jwtService))
	{
		paymentsNasabah.POST("/payment", paymentController.PaymentLoanController)
		paymentsNasabah.GET("/status/:status", paymentController.ListPaymentByStatusController)
		paymentsNasabah.PUT("/:id", paymentController.UpdatePaymentController)
		paymentsNasabah.GET("/:id", paymentController.GetPaymentPerBulanController)
		paymentsNasabah.GET("/total-payments/:id", paymentController.GetTotalPaymentController)
		paymentsNasabah.DELETE("/:id", paymentController.DeletePaymentController)
	}

	historyPaymentNasabah := r.Group("app/history/payment", middleware.Authorize(jwtService))
	{
		historyPaymentNasabah.GET("/", historyPaymentController.GetAllHistoryPaymentController)
		historyPaymentNasabah.GET("/:id", historyPaymentController.GetHistoryPaymentByIdController)
	}

	r.Run("localhost:3000")
}
