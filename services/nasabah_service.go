package services

import (
	"golang-pinjol/dto"
	"golang-pinjol/model"
	"golang-pinjol/repository"
	"log"

	"github.com/mashingan/smapping"
)

type NasabahServices interface {
	UpdateNasabah(nasabah dto.UpdateNasabahDTO) *model.Master_Customer
	ProfileNasabah(nasabahId string) model.Master_Customer
}

type nasabahService struct {
	nasabahRepository repository.NasabahRepository
}

func NewNasabahService(nasabahRepo repository.NasabahRepository) NasabahServices {
	return &nasabahService{
		nasabahRepository: nasabahRepo,
	}
}

func (s *nasabahService) UpdateNasabah(nasabah dto.UpdateNasabahDTO) *model.Master_Customer {
	var NewNasabah model.Master_Customer
	err := smapping.FillStruct(&NewNasabah, smapping.MapFields(&nasabah))
	if err != nil {
		log.Printf("Error map %v", err)
	}

	update, _ := s.nasabahRepository.UpdateNasabah(NewNasabah.Id, &NewNasabah)
	return update
}

func (s *nasabahService) ProfileNasabah(nasabahId string) model.Master_Customer {
	return s.nasabahRepository.ProfileNasabah(nasabahId)
}
