package model

// import "time"

// type Master_Document_Nasabah struct {
// 	ID          uint64           `gorm:"primaryKey;autoIncrement" json:"id"`
// 	Nasabah_Id  uint64         `gorm:"not null" json:"nasabah_id,omitempty"`
// 	FileName    string         `gorm:"not null" json:"filename"`
// 	ContentType string         `gorm:"not null" json:"content_type"`
// 	FileData    string         `gorm:"type:text;not null" json:"-"`
// 	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
// 	Nasabah     Master_Nasabah `gorm:"association_foreignkey:Nasabah_Id;association_autoupdate:false" json:"nasabah,omitempty"`
// }

// type Master_Document_Nasabah struct {
// 	Id           uint64         `gorm:"primary_key;auto_increment" json:"id"`
// 	Nasabah_Id   uint64         `gorm:"not null" json:"nasabah_id,omitempty"`
// 	Nasabah      Master_Nasabah `gorm:"association_foreignkey:Nasabah_Id;association_autoupdate:false" json:"nasabah,omitempty"`
// 	DocumentType string         `gorm:"type:varchar(255)" json:"document_type"`
// }
