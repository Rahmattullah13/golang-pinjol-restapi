package controller

import (
	"golang-pinjol/dto"
	"golang-pinjol/helper"
	"golang-pinjol/model"
	"golang-pinjol/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentController interface {
	PaymentLoanController(ctx *gin.Context)
	UpdatePaymentController(ctx *gin.Context)
	ListPaymentByStatusController(ctx *gin.Context)
	GetPaymentPerBulanController(ctx *gin.Context)
	GetTotalPaymentController(ctx *gin.Context)
	DeletePaymentController(ctx *gin.Context)
}

type paymentController struct {
	paymentService services.PaymentService
	jwtService     services.JwtService
}

func NewPaymentController(ps services.PaymentService, js services.JwtService) PaymentController {
	return &paymentController{
		paymentService: ps,
		jwtService:     js,
	}
}

// @Summary Create Payment Loan
// @Tags Payment
// @Description Create a new payment for a loan
// @Accept json
// @Produce json
// @Param payment body dto.CreatePaymentDTO true "Payment details"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /app/payments/payment [post]
func (pc *paymentController) PaymentLoanController(ctx *gin.Context) {
	var payments dto.CreatePaymentDTO
	err := ctx.ShouldBind(&payments)
	if err != nil {
		response := helper.ErrorResponse("Failed to process request payment", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	payment, err := pc.paymentService.PaymentLoanService(&payments)
	if err != nil {
		response := helper.ErrorResponse("Failed to create payment", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	response := helper.ResponseOK(true, "OK!", payment)
	ctx.JSON(http.StatusOK, response)
}

// @Summary Update Payment
// @Tags Payment
// @Description Update an existing payment
// @Accept json
// @Produce json
// @Param id path int true "Payment ID"
// @Param updates body dto.UpdatePaymentDTO true "Updated payment details"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /app/payments/{id} [put]
func (pc *paymentController) UpdatePaymentController(ctx *gin.Context) {
	var updates dto.UpdatePaymentDTO
	err := ctx.ShouldBindJSON(&updates)
	if err != nil {
		response := helper.ErrorResponse("Failed to process request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response := helper.ErrorResponse("Failed to parse id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	updates.Id = int(id)
	updatePayment, err := pc.paymentService.UpdatePaymentService(updates)
	if err != nil {
		response := helper.ErrorResponse("Failed to update payment", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseOK(true, "OK!", updatePayment)
	ctx.JSON(http.StatusOK, response)
}

// @Summary List Payments by Status
// @Tags Payment
// @Description Get a list of payments by status
// @Produce json
// @Param status path string true "Payment status"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /app/payments/status/{status} [get]
func (pc *paymentController) ListPaymentByStatusController(ctx *gin.Context) {
	status := ctx.Param("status")
	payments, err := pc.paymentService.ListPaymentByStatusService(status)
	if err != nil {
		response := helper.ErrorResponse("Error fetching payments", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ResponseOK(true, "OK!", payments)
	ctx.JSON(http.StatusOK, response)
}

// @Summary Get Payment per Month
// @Tags Payment
// @Description Get payment details per month
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /app/payments/{id} [get]
func (pc *paymentController) GetPaymentPerBulanController(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.ErrorResponse("Failed to process request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	paymentPerMonth, err := pc.paymentService.GetPaymentPerMonthService(int(id))
	if err != nil {
		response := helper.ErrorResponse("Failed to process request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseOK(true, "OK!", paymentPerMonth)
	ctx.JSON(http.StatusOK, response)
}

// @Summary Get Total Payment
// @Tags Payment
// @Description Get total payment amount
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /app/payments/total-payments/{id} [get]
func (pc *paymentController) GetTotalPaymentController(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.ErrorResponse("Failed to parse id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	totalPayments, err := pc.paymentService.GetTotalPaymentService(int(id))
	if err != nil {
		response := helper.ErrorResponse("Failed to process id is not found", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseOK(true, "OK!", totalPayments)
	ctx.JSON(http.StatusOK, response)
}

// @Summary Delete Payment
// @Tags Payment
// @Description Delete an existing payment
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /app/payments/{id} [delete]
func (pc *paymentController) DeletePaymentController(ctx *gin.Context) {
	var txLoan model.Transactions_Payment_Loan
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.ErrorResponse("Failed to process parse id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	txLoan.ID = int(id)
	err = pc.paymentService.DeletePaymentService(txLoan.ID)
	if err != nil {
		response := helper.ErrorResponse("Failed to process request delete transaction", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseOK(true, "OK!", helper.EmptyObject{})
	ctx.JSON(http.StatusOK, response)
}
