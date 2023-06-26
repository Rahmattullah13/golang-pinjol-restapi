package dto


type CreatePekerjaanNasabahDTO struct{
	Customer_Id int `json:"customer_id" form:"customer_id" binding:"required"`
	Company_Address string `json:"company_address" form:"company_address" binding:"required"`
	Payday_Date string `json:"payDay_date" form:"payDay_date" binding:"required"`
	Job_Position string `json:"job_position" form:"job_position" binding:"required"`
}

type UpdatePekerjaanNasabahDTO struct{
	Id int `json:"id" form:"id"`
	Customer_Id int `json:"customer_id" form:"customer_id" binding:"required"`
	Company_Address string `json:"company_address" form:"company_address" binding:"required"`
	Payday_Date string `json:"payDay_date" form:"payDay_date" binding:"required"`
	Job_Position string `json:"job_position" form:"job_position" binding:"required"`
}