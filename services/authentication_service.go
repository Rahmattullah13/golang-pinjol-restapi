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
	CreateCustomer(customer dto.RegisterCustomerDTO) *model.Master_Customer
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) bool
	IsDuplicateNIK(noKtp string) bool
	FindByNIK(nik string) model.Master_Customer
}

type authenticationService struct {
	customerRepository repository.CustomerRepository
}

func NewAuthenticationService(customer repository.CustomerRepository) AuthenticationService {
	return &authenticationService{
		customerRepository: customer,
	}
}

func (s *authenticationService) CreateCustomer(customer dto.RegisterCustomerDTO) *model.Master_Customer {
	NewCustomer := model.Master_Customer{}
	err := smapping.FillStruct(&NewCustomer, smapping.MapFields(customer))
	if err != nil {
		fmt.Printf("Error map %v", err)
	}

	response, _ := s.customerRepository.CreateCustomer(&NewCustomer)
	return response
}

func (s *authenticationService) VerifyCredential(email string, password string) interface{} {
	response := s.customerRepository.VerifyCredential(email, password)
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
	response := s.customerRepository.IsDuplicateEmail(email)
	return !(response.Error == nil)
}

func (s *authenticationService) IsDuplicateNIK(noKtp string) bool {
	response := s.customerRepository.IsDuplicateNIK(noKtp)
	return !(response.Error == nil)
}

func (s *authenticationService) FindByNIK(nik string) model.Master_Customer {
	return s.customerRepository.FindByNIK(nik)
}
