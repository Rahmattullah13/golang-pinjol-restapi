package dto

type CreatePinjamanDTO struct {
	Customer_Id uint64 `json:"customer_id" form:"customer_id" binding:"required"`
	Amount int `json:"amount" form:"amount" binding:"required"`
	Loan_Interest_Rates int `json:"loan_interest_rates" form:"loan_interest_rates"binding:"required"`
	Loan_Duration int `json:"loan_duration" form:"loan_duration" binding:"required"`
}

type UpdatePinjamanDTO struct {
	Id uint64 `json:"id" form:"id"`
	Customer_Id uint64 `json:"customer_id" form:"customer_id" binding:"required"`
	Amount int `json:"amount" form:"amount" binding:"required"`
	Loan_Interest_Rates int `json:"loan_interest_rates" form:"loan_interest_rates"binding:"required"`
	Loan_Duration int `json:"loan_duration" form:"loan_duration" binding:"required"`
}
