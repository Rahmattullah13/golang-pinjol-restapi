package dto

type CreateDocumentNasabahDTO struct{
	Customer_Id uint64 `json:"customer_id" form:"customer_id" binding:"required"`
	DocumentType string `json:"document_type" form:"document_type" binding:"required"`
	FilePath string `json:"file_path" form:"file_path" binding:"required"`
}

type UpdateDocumentNasabahDTO struct{
	Id uint64 `json:"id" form:"id"`
	Customer_Id uint64 `json:"customer_id" form:"customer_id" binding:"required"`
	DocumentType string `json:"document_type" form:"document_type" binding:"required"`
	FilePath string `json:"file_path" form:"file_path" binding:"required"`
}