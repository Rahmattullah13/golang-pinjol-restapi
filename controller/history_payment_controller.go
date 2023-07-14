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

// @Summary Get all history payment
// @Description Get a list of all available history payments
// @Tags History Payment
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /app/history/payment/ [get]
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

// @Summary Get history payment by ID
// @Description Get a history payment by its ID
// @Tags History Payment
// @Param id path int true "ID of the history payment"
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /app/history/payment/{id} [get]
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
