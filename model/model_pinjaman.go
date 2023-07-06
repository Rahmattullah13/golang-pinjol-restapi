package model

type Master_Loan struct {
	Id                  uint64         `gorm:"primary_key;auto_increment;column:id" json:"id"`
	Nasabah_Id          uint64         `gorm:"not null" json:"nasabah_id,omitempty"`
	Nasabah             Master_Nasabah `gorm:"association_foreignkey:Nasabah_Id;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	Amount              int            `gorm:"type:integer" json:"amount,omitempty"`
	Loan_Interest_Rates int            `gorm:"type:integer" json:"loan_interest_rates"`
	Loan_Duration       int            `gorm:"type:integer" json:"loan_duration"`
	StatusApproved      bool           `gorm:"type:boolean" json:"status_approved"`
}
