package model

import "time"

type Transactions_Payment_Loan struct {
	ID int`gorm:"primary_key;auto_increment" json:"id"`
	Loan_id int `gorm:"not null" json:"loan_id, omitempty"`
	Loan Master_Loan `gorm:"foreignKey:Loan_id;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"loan_id,omitmepty"`
	Monthly_Payments int `gorm:"type:integer" json:"monthly_payment"`
	Payment_Status bool `gorm:"type:boolean" json:"payment_status"`
	Payment_Date time.Time `gorm:"type:time" json:"payment_date"`
	}