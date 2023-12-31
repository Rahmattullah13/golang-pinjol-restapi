package controller

import (
	"golang-pinjol/model"
	"golang-pinjol/services"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UploadFileController interface {
	UploadFile(ctx *gin.Context)
}

type uploadFileController struct {
	jwtService services.JwtService
	db         *gorm.DB
}

// @Summary Upload File
// @Schemes {http, https}
// @Description Upload a file for a specific customer
// @Tags Document
// @Accept multipart/form-data
// @Param customer_id path string true "Customer ID"
// @Param document formData file true "Document file to upload"
// @Produce json
// @Success 200 {object} interface{}
// @Failure 400 {object} interface{}
// @Failure 401 {object} interface{}
// @Failure 500 {object} interface{}
// @Router /app/document/upload/{customer_id} [put]
func (uc *uploadFileController) UploadFile(ctx *gin.Context) {
	// Ensure the method is PUT
	if ctx.Request.Method != http.MethodPut {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	// Retrieve the nasabah_id from the URL
	customerID := ctx.Param("customer_id")

	file, header, err := ctx.Request.FormFile("document")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}
	defer file.Close()

	folder := "assets"
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.Mkdir(folder, os.ModePerm)
	}

	path := folder + "/" + header.Filename
	out, err := os.Create(path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Find the user with the given nasabah_id
	var customer model.Master_Customer
	result := uc.db.First(&customer, customerID)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find the user"})
		return
	}

	// Update status_verified to true if the request is successful
	if ctx.Writer.Status() == http.StatusOK {
		customer.StatusVerified = true
		result = uc.db.Save(&customer)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update status_verified"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "successfully upload document, status_verified: true", "path": path})
}

func NewUploadFileController(jwtService services.JwtService, db *gorm.DB) UploadFileController {
	return &uploadFileController{
		jwtService: jwtService,
		db:         db,
	}
}
