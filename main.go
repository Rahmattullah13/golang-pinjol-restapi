package main

import (
	"golang-pinjol/config"
	"golang-pinjol/controller"
	"golang-pinjol/middleware"
	"golang-pinjol/repository"
	"golang-pinjol/services"
	"log"

	_ "golang-pinjol/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()

	// Repository
	customerRepository repository.CustomerRepository       = repository.NewCustomerRepository(db)
	jobsRepository     repository.CustomerJobsRepository   = repository.NewCustomerJobsRepository(db)
	loanRepository     repository.LoansRepository          = repository.NewLoansRepository(db)
	paymentRepository  repository.PaymentRepository        = repository.NewPaymentRepository(db)
	historyRepository  repository.HistoryPaymentRepository = repository.NewHistoryPaymentRepository(db)

	// Service
	jwtService          services.JwtService            = services.NewJwtService()
	authService         services.AuthenticationService = services.NewAuthenticationService(customerRepository)
	customerService     services.CustomerServices      = services.NewCustomerService(customerRepository)
	customerJobsService services.JobsCustomerService   = services.NewJobsCustomerService(jobsRepository)
	loanService         services.LoanService           = services.NewLoanService(loanRepository, customerRepository)
	paymentService      services.PaymentService        = services.NewPaymentService(paymentRepository)
	historyService      services.HistoryPaymentService = services.NewHistoryPaymentService(historyRepository)

	// Controller
	authController           controller.AuthController           = controller.NewAuthController(authService, jwtService)
	customerController       controller.CustomerController       = controller.NewCustomerController(customerService, jwtService)
	uploadFileController     controller.UploadFileController     = controller.NewUploadFileController(jwtService, db)
	customerJobsController   controller.CustomerJobsController   = controller.NewJobsCustomerController(customerJobsService, jwtService)
	loanController           controller.LoanController           = controller.NewLoanController(loanService, jwtService)
	paymentController        controller.PaymentController        = controller.NewPaymentController(paymentService, jwtService)
	historyPaymentController controller.HistoryPaymentController = controller.NewHistoryPaymentController(historyService, jwtService)
)


// @title Golang Pinjol
// @description Dokumentasi REST API
// @version 0.1
// @host localhost:3000

func main() {
	defer config.CloseDB(db)
	log.Println(db)

	r := gin.Default()

	auth := r.Group("app/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	customer := r.Group("app/customer", middleware.Authorize(jwtService))
	{
		customer.PUT("/update", customerController.UpdateCustomerController)
		customer.GET("/profile", customerController.ProfileCustomerController)
	}

	documentCustomer := r.Group("app/document", middleware.Authorize(jwtService))
	{
		documentCustomer.PUT("/upload/:customer_id", uploadFileController.UploadFile)
	}

	customerJobs := r.Group("app/jobs", middleware.Authorize(jwtService))
	{
		customerJobs.POST("/addJobs", customerJobsController.AddCustomerJobsController)
		customerJobs.PUT("/:id", customerJobsController.UpdateCustomerJobsController)
		customerJobs.GET("/:id", customerJobsController.SearchCustomerJobsByIdController)
		customerJobs.DELETE("/:id", customerJobsController.DeleteCustomerJobsController)
	}

	customerLoans := r.Group("app/loans", middleware.Authorize(jwtService))
	{
		customerLoans.POST("/loan", loanController.CreateLoanController)
		customerLoans.PUT("/:id", loanController.UpdateLoanController)
		customerLoans.GET("/:id", loanController.SearchLoanByIdController)
		customerLoans.PUT("/verification/:id", loanController.UpdateStatusApprovalController)
		customerLoans.DELETE("/:id", loanController.DeleteLoanController)
	}

	customerPayments := r.Group("app/payments", middleware.Authorize(jwtService))
	{
		customerPayments.POST("/payment", paymentController.PaymentLoanController)
		customerPayments.GET("/status/:status", paymentController.ListPaymentByStatusController)
		customerPayments.PUT("/:id", paymentController.UpdatePaymentController)
		customerPayments.GET("/:id", paymentController.GetPaymentPerBulanController)
		customerPayments.GET("/total-payments/:id", paymentController.GetTotalPaymentController)
		customerPayments.DELETE("/:id", paymentController.DeletePaymentController)
	}

	historyPaymentCustomer := r.Group("app/history/payment", middleware.Authorize(jwtService))
	{
		historyPaymentCustomer.GET("/", historyPaymentController.GetAllHistoryPaymentController)
		historyPaymentCustomer.GET("/:id", historyPaymentController.GetHistoryPaymentByIdController)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run("localhost:3000")
}
