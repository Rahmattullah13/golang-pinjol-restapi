package controller

import (
	"golang-pinjol/helper"
	"golang-pinjol/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HistoryPaymentController interface {
	GetAllHistoryPaymentController(ctx *gin.Context)
	GetHistoryPaymentByIdController(ctx *gin.Context)
}

type historyPaymentController struct {
	hps        services.HistoryPaymentService
	jwtService services.JwtService
}

func NewHistoryPaymentController(historyPaymentService services.HistoryPaymentService, jwtService services.JwtService) HistoryPaymentController {
	return &historyPaymentController{
		hps:        historyPaymentService,
		jwtService: jwtService,
	}
}

func (c *historyPaymentController) GetAllHistoryPaymentController(ctx *gin.Context) {
	history, err := c.hps.GetAllHistoriesPaymentService()
	if err != nil {
		log.Printf("Error history controller %v", err)
		response := helper.ErrorResponse("Failed to process request get all data", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ResponseOK(true, "OK!", history)
	ctx.JSON(http.StatusOK, response)
}

func (c *historyPaymentController) GetHistoryPaymentByIdController(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.ErrorResponse("Failed to process parse id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	history, err := c.hps.GetHistoryPaymentByIdService(id)
	if err != nil {
		response := helper.ErrorResponse("Failed to process get by id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseOK(true, "OK!", history)
	ctx.JSON(http.StatusOK, response)
}
