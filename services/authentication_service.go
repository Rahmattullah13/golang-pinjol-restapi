package services

import (
	"fmt"
	"golang-pinjol/dto"
	"golang-pinjol/model"
	"golang-pinjol/repository"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationService interface {
	CreateNasabah(nasabah dto.RegisterNasabahDTO) *model.Master_Nasabah
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) bool
	IsDuplicateNIK(noKtp string) bool
	FindByNIK(nik string) model.Master_Nasabah
}

type authenticationService struct {
	nasabahRepository repository.NasabahRepository
}

func NewAuthenticationService(nasabah repository.NasabahRepository) AuthenticationService {
	return &authenticationService{
		nasabahRepository: nasabah,
	}
}

func (s *authenticationService) CreateNasabah(nasabah dto.RegisterNasabahDTO) *model.Master_Nasabah {
	NewNasabah := model.Master_Nasabah{}
	err := smapping.FillStruct(&NewNasabah, smapping.MapFields(nasabah))
	if err != nil {
		fmt.Errorf("Error map %v", err)
	}

	response, _ := s.nasabahRepository.CreateNasabah(&NewNasabah)
	return response
}

func (s *authenticationService) VerifyCredential(email string, password string) interface{} {
	response := s.nasabahRepository.VerifyCredential(email, password)
	if v, ok := response.(model.Master_Nasabah); ok {
		comparePassword := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(password))
		if v.Email == email && comparePassword == nil {
			return response
		}
		return false
	}
	return nil
}

func (s *authenticationService) IsDuplicateEmail(email string) bool {
	response := s.nasabahRepository.IsDuplicateEmail(email)
	return !(response.Error == nil)
}

func (s *authenticationService) IsDuplicateNIK(noKtp string) bool {
	response := s.nasabahRepository.IsDuplicateNIK(noKtp)
	return !(response.Error == nil)
}

func (s *authenticationService) FindByNIK(nik string) model.Master_Nasabah {
	return s.nasabahRepository.FindByNIK(nik)
}
