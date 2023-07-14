package services

import (
	"golang-pinjol/dto"
	"golang-pinjol/model"
	"golang-pinjol/repository"
	"log"

	"github.com/mashingan/smapping"
)

type CustomerServices interface {
	UpdateCustomer(customer dto.UpdateCustomerDTO) *model.Master_Customer
	ProfileCustomer(customerId string) model.Master_Customer
}

type customerService struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) CustomerServices {
	return &customerService{
		customerRepository: customerRepo,
	}
}

func (s *customerService) UpdateCustomer(customer dto.UpdateCustomerDTO) *model.Master_Customer {
	var NewCustomer model.Master_Customer
	err := smapping.FillStruct(&NewCustomer, smapping.MapFields(&customer))
	if err != nil {
		log.Printf("Error map %v", err)
	}
	update, _ := s.customerRepository.UpdateCustomer(NewCustomer.Id, &NewCustomer)
	return update
}

func (s *customerService) ProfileCustomer(customerId string) model.Master_Customer {
	return s.customerRepository.ProfileCustomer(customerId)
}
