package dto

type CreatePaymentDTO struct {
	Loan_id          int `form:"loan_id" json:"loan_id" binding:"required"`
	Monthly_Payments int `form:"monthly_payment" json:"monthly_payment"`
}

type UpdatePaymentDTO struct {
	Id               int `form:"id" json:"id"`
	Loan_id          int `form:"loan_id" json:"loan_id" binding:"required"`
	Monthly_Payments int `form:"monthly_payment" json:"monthly_payment" binding:"required"`
}
