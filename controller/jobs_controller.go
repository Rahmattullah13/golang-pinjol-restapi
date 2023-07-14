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

type CustomerJobsController interface {
	AddCustomerJobsController(ctx *gin.Context)
	UpdateCustomerJobsController(ctx *gin.Context)
	SearchCustomerJobsByIdController(ctx *gin.Context)
	DeleteCustomerJobsController(ctx *gin.Context)
}

type customerJobsController struct {
	customerJobsService services.JobsCustomerService
	jwtService          services.JwtService
}

func NewJobsCustomerController(customerJobs services.JobsCustomerService, jwtService services.JwtService) CustomerJobsController {
	return &customerJobsController{
		customerJobsService: customerJobs,
		jwtService:          jwtService,
	}
}

// @Summary Add Customer Jobs
// @Description Add new customer jobs
// @Tags Customer Jobs
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param addJobsCustomerDTO body dto.CreateJobsCustomerDTO true "Add Customer Jobs DTO"
// @Success 200 {object} helper.Response
// @Router /app/jobs/addJobs [post]
func (c *customerJobsController) AddCustomerJobsController(ctx *gin.Context) {
	var addJobsCustomerDTO dto.CreateJobsCustomerDTO
	err := ctx.ShouldBind(&addJobsCustomerDTO)
	if err != nil {
		response := helper.ErrorResponse("failed to process request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		v, err := c.customerJobsService.AddCustomerJobsService(addJobsCustomerDTO)
		if err != nil {
			response := helper.ErrorResponse("failed to process request add jobs customer", err.Error(), helper.EmptyObject{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		} else {
			response := helper.ResponseOK(true, "OK!", v)
			ctx.JSON(http.StatusOK, response)
		}
	}

}

// @Summary Update Customer Jobs
// @Description Update existing customer jobs
// @Tags Customer Jobs
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Customer Job ID"
// @Param updateCustomerJobsDTO body dto.UpdateJobsCustomerDTO true "Update Customer Jobs DTO"
// @Success 200 {object} helper.Response
// @Router /app/jobs/{id} [put]
func (c *customerJobsController) UpdateCustomerJobsController(ctx *gin.Context) {
	var updateCustomerJobsDTO dto.UpdateJobsCustomerDTO
	err := ctx.ShouldBind(&updateCustomerJobsDTO)
	if err != nil {
		response := helper.ErrorResponse("failed to process request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.ErrorResponse("failed to process request update customer jobs", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	updateCustomerJobsDTO.Id = int(id)
	update, _ := c.customerJobsService.UpdateCustomerJobsService(updateCustomerJobsDTO)
	response := helper.ResponseOK(true, "OK!", update)
	ctx.JSON(http.StatusOK, response)
}

// @Summary Search Customer Jobs By ID
// @Description Get customer jobs by ID
// @Tags Customer Jobs
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Customer Job ID"
// @Success 200 {object} helper.Response
// @Router /app/jobs/{id} [get]
func (c *customerJobsController) SearchCustomerJobsByIdController(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.ErrorResponse("failed to process request", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	jobs, err := c.customerJobsService.SearchCustomerJobsByIdService(int(id))
	if (jobs == &model.Master_Jobs_Customer{}) {
		response := helper.ErrorResponse("failed to process data id not found", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		response := helper.ResponseOK(true, "OK!", jobs)
		ctx.JSON(http.StatusOK, response)
	}
}

// @Summary Delete Customer Jobs
// @Description Delete customer jobs
// @Tags Customer Jobs
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Customer Job ID"
// @Success 200 {object} helper.Response
// @Router /app/jobs/{id} [delete]
func (c *customerJobsController) DeleteCustomerJobsController(ctx *gin.Context) {
	var deleteJobsCustomer model.Master_Jobs_Customer
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.ErrorResponse("failed to process request parse id", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	deleteJobsCustomer.ID = int(id)
	deletes := c.customerJobsService.DeleteCustomerJobsService(deleteJobsCustomer.ID)
	response := helper.ResponseOK(true, "OK!", deletes)
	ctx.JSON(http.StatusOK, response)
}
