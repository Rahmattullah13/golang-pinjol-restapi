package controller

import (
	"golang-pinjol/dto"
	"golang-pinjol/helper"
	"golang-pinjol/model"
	"golang-pinjol/services"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoanController interface {
	CreateLoanController(ctx *gin.Context)
	UpdateLoanController(ctx *gin.Context)
	SearchLoanByIdController(ctx *gin.Context)
	DeleteLoanController(ctx *gin.Context)
	UpdateStatusApprovalController(ctx *gin.Context)
}

type loanController struct {
	loanService services.LoanService
	jwtService  services.JwtService
}

func NewLoanController(ls services.LoanService, js services.JwtService) LoanController {
	return &loanController{
		loanService: ls,
		jwtService:  js,
	}
}

// @Summary Create Loan
// @Description Create a new loan
// @Tags Loans
// @Accept json
// @Produce json
// @Param loan body dto.CreateLoanDTO true "Loan object to create"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /app/loans/loan [post]
func (lc *loanController) CreateLoanController(ctx *gin.Context) {
	var createLoanDTO dto.CreateLoanDTO
	err := ctx.ShouldBind(&createLoanDTO)
	if err != nil {
		response := helper.ErrorResponse("failed to process request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		v, err := lc.loanService.CreateLoanService(createLoanDTO)
		if err != nil {
			response := helper.ErrorResponse("failed to process request create loan", err.Error(), helper.EmptyObject{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		} else {
			response := helper.ResponseOK(true, "OK!", v)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

// @Summary Update Loan
// @Description Update an existing loan
// @Tags Loans
// @Accept json
// @Produce json
// @Param id path string true "Loan ID"
// @Param loan body dto.UpdateLoanDTO true "Loan object to update"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /app/loans/{id} [put]
func (lc *loanController) UpdateLoanController(ctx *gin.Context) {
	var updateLoanDTO dto.UpdateLoanDTO
	err := ctx.ShouldBind(&updateLoanDTO)
	if err != nil {
		response := helper.ErrorResponse("failed to process request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
		if err != nil {
			response := helper.ErrorResponse("failed to procces parse id", err.Error(), helper.EmptyObject{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		updateLoanDTO.Id = id
		update, err := lc.loanService.UpdateLoanService(updateLoanDTO)
		if err != nil {
			response := helper.ErrorResponse("failed to process id not found", err.Error(), helper.EmptyObject{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			response := helper.ResponseOK(true, "OK!", update)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

// @Summary Search Loan by ID
// @Description Get loan details by ID
// @Tags Loans
// @Produce json
// @Param id path string true "Loan ID"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /app/loans/{id} [get]
func (lc *loanController) SearchLoanByIdController(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.ErrorResponse("failed to process parse id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	loan, err := lc.loanService.SearchLoanByIdService(id)
	if reflect.DeepEqual(*loan, model.Master_Loan{}) {
		response := helper.ErrorResponse("failed to process data id not found", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		response := helper.ResponseOK(true, "OK!", loan)
		ctx.JSON(http.StatusOK, response)
	}
}

// @Summary Delete Loan
// @Description Delete an existing loan
// @Tags Loans
// @Produce json
// @Param id path string true "Loan ID"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /app/loans/{id} [delete]
func (lc *loanController) DeleteLoanController(ctx *gin.Context) {
	var deleteLoan model.Master_Loan
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.ErrorResponse("failed to process request parse id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	deleteLoan.Id = id
	deletes := lc.loanService.DeleteLoanService(deleteLoan.Id)
	response := helper.ResponseOK(true, "OK!", deletes)
	ctx.JSON(http.StatusOK, response)
}

// @Summary Update Loan Approval Status
// @Description Update the approval status of a loan
// @Tags Loans
// @Produce json
// @Param id path string true "Loan ID"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /app/loans/verification/{id} [put]
func (lc *loanController) UpdateStatusApprovalController(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.ErrorResponse("failed to process request parse id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	loan, err := lc.loanService.UpdateLoanStatusService(id)
	if err != nil {
		response := helper.ErrorResponse("failed to update loan approval status", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.ResponseOK(true, "OK!", loan)
		ctx.JSON(http.StatusOK, response)
	}
}
