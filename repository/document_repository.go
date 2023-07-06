package repository

// import (
// 	"golang-pinjol/model"

// 	"gorm.io/gorm"
// )

// type UploadFileRepostory interface {
// 	SaveFile(document *model.Master_Document_Nasabah) (*model.Master_Document_Nasabah, error)
// }

// type uploadedFileRepository struct {
// 	db *gorm.DB
// }

// func (u *uploadedFileRepository) SaveFile(document *model.Master_Document_Nasabah) (*model.Master_Document_Nasabah, error) {
// 	tx := u.db.Begin()

// 	if err := tx.Create(document).Error; err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}

// 	if err := tx.Model(&model.Master_Nasabah{Id: document.Nasabah_Id}).UpdateColumn("status_verified", true).Error; err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}

// 	tx.Commit()

// 	return document, nil

// }

// func NewUploadFileRepository(db *gorm.DB) UploadFileRepostory {
// 	return &uploadedFileRepository{
// 		db: db,
// 	}
// }
