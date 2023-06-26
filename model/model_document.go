package model

type Master_Document_Customer struct {
	Id uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Customer_Id uint64 ` gorm:"not null" json:"customer_id, omitempty"`
	Customer Master_Customer `gorm:"association_foreignkey:Customer_Id;association_autoupdate:false" json:"customer_id, omitempty"`
	DocumentType string `gorm:"type:varchar(255)" json:"document_type"`
	FilePath string `gorm:"type:varchar(255)" json:"file_path"`
}
