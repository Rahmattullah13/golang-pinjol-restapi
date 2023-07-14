package controller

import (
	"fmt"
	"golang-pinjol/dto"
	"golang-pinjol/helper"
	"golang-pinjol/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	UpdateCustomerController(ctx *gin.Context)
	ProfileCustomerController(ctx *gin.Context)
}

type customerController struct {
	customerService services.CustomerServices
	jwtService      services.JwtService
}

func NewCustomerController(customerservice services.CustomerServices, jwtService services.JwtService) CustomerController {
	return &customerController{
		customerService: customerservice,
		jwtService:      jwtService,
	}
}

// @Summary Update Customer
// @Tags Customer
// @Description Update customer information
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param updateCustomer body dto.UpdateCustomerDTO true "Customer data to be updated"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /app/customer/update [put]
func (c *customerController) UpdateCustomerController(context *gin.Context) {
	var customerUpdateDTO dto.UpdateCustomerDTO
	err := context.ShouldBind(&customerUpdateDTO)
	if err != nil {
		response := helper.ErrorResponse("Fail to process request", err.Error(), helper.EmptyObject{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateTokenService(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["customer_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	customerUpdateDTO.Id = id
	customer := c.customerService.UpdateCustomer(customerUpdateDTO)
	response := helper.ResponseOK(true, "OK!", customer)
	context.JSON(http.StatusOK, response)
}

// @Summary Get Customer Profile
// @Tags Customer
// @Description Get customer profile
// @Security BearerAuth
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /app/customer/profile [get]
func (nc *customerController) ProfileCustomerController(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := nc.jwtService.ValidateTokenService(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["customer_id"])
	customer := nc.customerService.ProfileCustomer(id)
	response := helper.ResponseOK(true, "OK!", customer)
	ctx.JSON(http.StatusOK, response)
}
