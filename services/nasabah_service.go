package services

import (
	"golang-pinjol/dto"
	"golang-pinjol/model"
	"golang-pinjol/repository"
	"log"

	"github.com/mashingan/smapping"
)

type NasabahServices interface {
	UpdateNasabah(nasabah dto.UpdateNasabahDTO) *model.Master_Nasabah
	ProfileNasabah(nasabahId string) model.Master_Nasabah
}

type nasabahService struct {
	nasabahRepository repository.NasabahRepository
}

func NewNasabahService(nasabahRepo repository.NasabahRepository) NasabahServices {
	return &nasabahService{
		nasabahRepository: nasabahRepo,
	}
}

func (s *nasabahService) UpdateNasabah(nasabah dto.UpdateNasabahDTO) *model.Master_Nasabah {
	var NewNasabah model.Master_Nasabah
	err := smapping.FillStruct(&NewNasabah, smapping.MapFields(&nasabah))
	if err != nil {
		log.Printf("Error map %v", err)
	}
	update, _ := s.nasabahRepository.UpdateNasabah(NewNasabah.Id, &NewNasabah)
	return update
}

func (s *nasabahService) ProfileNasabah(nasabahId string) model.Master_Nasabah {
	return s.nasabahRepository.ProfileNasabah(nasabahId)
}
