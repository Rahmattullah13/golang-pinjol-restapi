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

type NasabahController interface {
	UpdateNasabahController(ctx *gin.Context)
	ProfileNasabahController(ctx *gin.Context)
}

type nasabahController struct {
	nasabahService services.NasabahServices
	jwtService     services.JwtService
}

func NewNasabahController(nasabahservice services.NasabahServices, jwtService services.JwtService) NasabahController {
	return &nasabahController{
		nasabahService: nasabahservice,
		jwtService:     jwtService,
	}
}

func (nc *nasabahController) UpdateNasabahController(ctx *gin.Context) {
	var nasabahUpdateDTO dto.UpdateNasabahDTO
	err := ctx.ShouldBind(&nasabahUpdateDTO)
	if err != nil {
		response := helper.ErrorResponse("Fail to process request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := nc.jwtService.ValidateTokenService(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["customer_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	nasabahUpdateDTO.Id = id
	nasabah := nc.nasabahService.UpdateNasabah(nasabahUpdateDTO)
	response := helper.ResponseOK(true, "OK", nasabah)
	ctx.JSON(http.StatusOK, response)
}

func (nc *nasabahController) ProfileNasabahController(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := nc.jwtService.ValidateTokenService(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["customer_id"])
	nasabah := nc.nasabahService.ProfileNasabah(id)
	response := helper.ResponseOK(true, "OK!", nasabah)
	ctx.JSON(http.StatusOK, response)
}
