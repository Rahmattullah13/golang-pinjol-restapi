package model

type Master_Jobs_Nasabah struct {
	ID              int            `gorm:"primary_key;auto_increment" json:"id"`
	Nasabah_Id      int            `gorm:"not null" json:"nasabah_id,omitempty"`
	Company_Address string         `gorm:"type:varchar(255)" json:"company_address"`
	Payday_Date     string         `gorm:"type:varchar(50)" json:"payDay_date"`
	Job_Position    string         `gorm:"type:varchar(255)" json:"job_position"`
	Nasabah        Master_Nasabah `gorm:"association_foreignkey:Nasabah_Id" json:"nasabah,omitempty"`
}
