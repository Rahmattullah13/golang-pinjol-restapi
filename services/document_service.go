package services

// import (
// 	"golang-pinjol/dto"
// 	"golang-pinjol/model"
// 	"golang-pinjol/repository"
// 	"log"

// 	"github.com/mashingan/smapping"
// )

// type UploadFileService interface {
// 	UploadFile(document *dto.UploadedFileDTO) (*model.Master_Document_Nasabah, error)
// }

// type uploadFileService struct {
// 	uploadFileRepository repository.UploadFileRepostory
// }

// func NewUploadFileService(uploadFileRepository repository.UploadFileRepostory) UploadFileService {
// 	return &uploadFileService{
// 		uploadFileRepository: uploadFileRepository,
// 	}
// }

// func (u *uploadFileService) UploadFile(document *dto.UploadedFileDTO) (*model.Master_Document_Nasabah, error) {
// 	var uploadFile model.Master_Document_Nasabah
// 	err := smapping.FillStruct(&uploadFile, smapping.MapFields(&document))
// 	if err != nil {
// 		log.Printf("Error mapping fields : %v", err)
// 		return nil, err
// 	}

// 	upload, err := u.uploadFileRepository.SaveFile(&uploadFile)
// 	if err != nil {
// 		log.Printf("Error Saving document: %v", err)
// 		return nil, err
// 	}

// 	return upload, nil
// }
