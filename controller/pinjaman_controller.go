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

func (lc *loanController) CreateLoanController(ctx *gin.Context) {
	var createLoanDTO dto.CreatePinjamanDTO
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

func (lc *loanController) UpdateLoanController(ctx *gin.Context) {
	var updateLoanDTO dto.UpdatePinjamanDTO
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
