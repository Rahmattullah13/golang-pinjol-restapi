package model

type Master_Customer struct {
	Id uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
	Email string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"-"`
	PhoneNumber string `gorm:"type:varchar(255)" json:"phone_number"`
	Address string `gorm:"type:varchar(255)" json:"address"`
	NoKtp string `gorm:"uniqueIndex;type:varchar(255)" json:"no_ktp"`
	StatusVerified bool `gorm:"type:boolean; not null;default:false" json:"status_verified"`
	Token  string  `gorm:"-" json:"token,omitempty"`
	Jobs []Master_Jobs_customers `gorm:"foreignKey:Customer_Id" json:"jobs,omitempty"`
	Master_Document_Customers []Master_Document_Customer `gorm:"foreignKey:Customer_Id" json:"document,omitempty"`

}
