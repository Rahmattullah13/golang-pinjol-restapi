package dto


type CreatePembayaranDTO struct {
	Loan_id int `form:"loan_id" json:"loan_id" binding:"required"`
	Monthly_Payments int `form:"monthly_payment" json:"monthly_payment"`
}

type UpdatePembayaranDTO struct {
	Id int `form:"id" json:"id"`
	Loan_id int `form:"loan_id" json:"loan_id" binding:"required"`
	Monthly_Payments int `form:"monthly_payment" json:"monthly_payment" binding:"required"`
}

