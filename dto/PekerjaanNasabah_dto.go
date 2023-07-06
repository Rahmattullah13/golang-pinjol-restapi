package dto


type CreatePekerjaanNasabahDTO struct{
	Nasabah_Id int `json:"nasabah_id" form:"nasabah_id" binding:"required"`
	Company_Address string `json:"company_address" form:"company_address" binding:"required"`
	Payday_Date string `json:"payDay_date" form:"payDay_date" binding:"required"`
	Job_Position string `json:"job_position" form:"job_position" binding:"required"`
}

type UpdatePekerjaanNasabahDTO struct{
	Id int `json:"id" form:"id"`
	Nasabah_Id int `json:"nasabah_id" form:"nasabah_id" binding:"required"`
	Company_Address string `json:"company_address" form:"company_address" binding:"required"`
	Payday_Date string `json:"payDay_date" form:"payDay_date" binding:"required"`
	Job_Position string `json:"job_position" form:"job_position" binding:"required"`
}