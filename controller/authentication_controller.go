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

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService services.AuthenticationService
	jwtService  services.JwtService
}

func NewAuthController(authService services.AuthenticationService, jwtService services.JwtService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

// @Summary Register Customer
// @Description Register a new customer
// @Tags Authentication
// @Accept json
// @Produce json
// @Param registerDTO body dto.RegisterCustomerDTO true "Registration details"
// @Success 201 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Router /app/auth/register [post]
func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterCustomerDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.ErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) || !c.authService.IsDuplicateNIK(registerDTO.NoKtp) {
		response := helper.ErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObject{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdCustomer := c.authService.CreateCustomer(registerDTO)
		token := c.jwtService.GenerateTokenService(strconv.FormatUint(createdCustomer.Id, 10))
		createdCustomer.Token = token
		response := helper.ResponseOK(true, "OK!", createdCustomer)
		ctx.JSON(http.StatusCreated, response)
	}
}

// @Summary Customer Login
// @Description Login a customer
// @Tags Authentication
// @Accept json
// @Produce json
// @Param loginDTO body dto.LoginCustomerDTO true "Login details"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Router /app/auth/login [post]
func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginCustomerDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.ErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if authResult == nil {
		response := helper.ErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	v := authResult.(model.Master_Customer)
	generatedToken := c.jwtService.GenerateTokenService(strconv.FormatUint(v.Id, 10))
	v.Token = generatedToken
	response := helper.ResponseOK(true, "OK!", v)
	ctx.JSON(http.StatusOK, response)
}
