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

type PekerjaanNasabahController interface {
	AddNasabahJobsController(ctx *gin.Context)
	NasabahUpdateJobsController(ctx *gin.Context)
	SearchNasabahJobsByIdController(ctx *gin.Context)
	DeleteNasabahJobsController(ctx *gin.Context)
}

type pekerjaanNasabahController struct {
	nasabahJobsService services.PekerjaanNasabahService
	jwtService         services.JwtService
}

func NewPekerjaanNasabahController(nasabahJobs services.PekerjaanNasabahService, jwtService services.JwtService) PekerjaanNasabahController {
	return &pekerjaanNasabahController{
		nasabahJobsService: nasabahJobs,
		jwtService:         jwtService,
	}
}

func (pnc *pekerjaanNasabahController) AddNasabahJobsController(ctx *gin.Context) {
	var addJobsNasabahDTO dto.CreatePekerjaanNasabahDTO
	err := ctx.ShouldBind(&addJobsNasabahDTO)
	if err != nil {
		response := helper.ErrorResponse("failed to process request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		v, err := pnc.nasabahJobsService.AddNasabahJobsService(addJobsNasabahDTO)
		if err != nil {
			response := helper.ErrorResponse("failed to process request add jobs nasabah", err.Error(), helper.EmptyObject{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			response := helper.ResponseOK(true, "OK!", v)
			ctx.JSON(http.StatusOK, response)
		}
	}

}

func (pnc *pekerjaanNasabahController) NasabahUpdateJobsController(ctx *gin.Context) {
	var updateJobsNasabahDTO dto.UpdatePekerjaanNasabahDTO
	err := ctx.ShouldBind(&updateJobsNasabahDTO)
	if err != nil {
		response := helper.ErrorResponse("failed to process request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.ErrorResponse("failed to process request update nasabah jobs", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	updateJobsNasabahDTO.Id = int(id)
	update, _ := pnc.nasabahJobsService.NasabahJobsUpdateService(updateJobsNasabahDTO)
	response := helper.ResponseOK(true, "OK!", update)
	ctx.JSON(http.StatusOK, response)
}

func (pnc *pekerjaanNasabahController) SearchNasabahJobsByIdController(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.ErrorResponse("failed to process request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	jobs, err := pnc.nasabahJobsService.SearchNasabahJobsByIdService(int(id))
	if (jobs == &model.Master_Jobs_Nasabah{}) {
		response := helper.ErrorResponse("failed to process data id not found", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		response := helper.ResponseOK(true, "OK!", jobs)
		ctx.JSON(http.StatusOK, response)
	}
}

func (pnc *pekerjaanNasabahController) DeleteNasabahJobsController(ctx *gin.Context) {
	var deleteJobsNasabah model.Master_Jobs_Nasabah
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.ErrorResponse("failed to process request parse id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	deleteJobsNasabah.ID = int(id)
	deletes := pnc.nasabahJobsService.DeleteNasabahJobsService(deleteJobsNasabah.ID)
	response := helper.ResponseOK(true, "OK!", deletes)
	ctx.JSON(http.StatusOK, response)
}
