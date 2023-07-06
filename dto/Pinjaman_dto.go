package dto

type CreatePinjamanDTO struct {
	Nasabah_Id          uint64 `json:"nasabah_id" form:"nasabah_id" binding:"required"`
	Amount              int    `json:"amount" form:"amount" binding:"required"`
	Loan_Interest_Rates int    `json:"loan_interest_rates" form:"loan_interest_rates"binding:"required"`
	Loan_Duration       int    `json:"loan_duration" form:"loan_duration" binding:"required"`
}

type UpdatePinjamanDTO struct {
	Id                  uint64 `json:"id" form:"id"`
	Nasabah_Id          uint64 `json:"nasabah_id" form:"nasabah_id" binding:"required"`
	Amount              int    `json:"amount" form:"amount" binding:"required"`
	Loan_Interest_Rates int    `json:"loan_interest_rates" form:"loan_interest_rates"binding:"required"`
	Loan_Duration       int    `json:"loan_duration" form:"loan_duration" binding:"required"`
}
