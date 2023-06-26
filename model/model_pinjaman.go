package model

type Master_Loan struct {
	Id uint64 `gorm:"primary_key;auto_increment;column:id" json:"id"`
	Customer_Id uint64 `gorm:"not null" json:"customer_id,omitempty"`
	Customer Master_Customer `gorm:"association_foreignkey:Customer_Id;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"customer_id, omitempty"`
	Amount int `gorm:"type:integer" json:"amount,omitempty"`
	Loan_Interest_Rates int `gorm:"type:integer" json:"loan_interest_rates"`
	Loan_Duration int `gorm:"type:integer" json:"loan_duration"`
	StatusApproved bool `gorm:"type:boolean" json:"status_approved"`
}
