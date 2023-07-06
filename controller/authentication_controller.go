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

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterNasabahDTO
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
		createdNasabah := c.authService.CreateNasabah(registerDTO)
		token := c.jwtService.GenerateTokenService(strconv.FormatUint(createdNasabah.Id, 10))
		createdNasabah.Token = token
		response := helper.ResponseOK(true, "OK!", createdNasabah)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginNasabahDTO
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
	v := authResult.(model.Master_Nasabah)
	generatedToken := c.jwtService.GenerateTokenService(strconv.FormatUint(v.Id, 10))
	v.Token = generatedToken
	response := helper.ResponseOK(true, "OK!", v)
	ctx.JSON(http.StatusOK, response)
}
