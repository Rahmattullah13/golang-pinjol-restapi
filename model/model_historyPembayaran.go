package model

import "time"

type Master_Payment_History struct {
	Id uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Loan_id int `gorm:"not null" json:"loan_id, omitempty"`
	Payment_id int `gorm:"not null"json:"payment_id, omitempty"`
	Date time.Time `gorm:"not null"`
	Loan Master_Loan `gorm:"foreignKey:Loan_id;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"loan, omitempty"`
	Transaction Transactions_Payment_Loan `gorm:"foreignKey:Payment_id;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"payment, omitempty"`
}