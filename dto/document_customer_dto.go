package dto

type UploadedFileDTO struct {
	FileName    string `json:"fileName" binding:"required"`
	ContentType string `json:"contentType" binding:"required"`
	FileData    string `json:"fileData" binding:"required"`
}
