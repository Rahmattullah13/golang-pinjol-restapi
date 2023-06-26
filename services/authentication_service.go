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
	CreateNasabah(nasabah dto.RegisterNasabahDTO) *model.Master_Customer
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) bool
	IsDuplicateNIK(noKtp string) bool
	FindByNIK(nik string) model.Master_Customer
}

type authenticationService struct {
	nasabahRepository repository.NasabahRepository
}

func NewAuthenticationService(nasabah repository.NasabahRepository) AuthenticationService {
	return &authenticationService{
		nasabahRepository: nasabah,
	}
}

func (s *authenticationService) CreateNasabah(nasabah dto.RegisterNasabahDTO) *model.Master_Customer {
	NewNasabah := model.Master_Customer{}
	err := smapping.FillStruct(&NewNasabah, smapping.MapFields(nasabah))
	if err != nil {
		fmt.Errorf("error map %v", err)
	}

	response, _ := s.nasabahRepository.CreateNasabah(&NewNasabah)
	return response
}

func (s *authenticationService) VerifyCredential(email string, password string) interface{} {
	response := s.nasabahRepository.VerifyCredential(email, password)
	if v, ok := response.(model.Master_Customer); ok {
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

func (s *authenticationService) FindByNIK(nik string) model.Master_Customer {
	return s.nasabahRepository.FindByNIK(nik)
}