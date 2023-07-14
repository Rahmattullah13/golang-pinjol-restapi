package services

import (
	"errors"
	"golang-pinjol/dto"
	"golang-pinjol/model"
	"golang-pinjol/repository"
	"log"
	"strconv"

	"github.com/mashingan/smapping"
)

type LoanService interface {
	CreateLoanService(loan dto.CreateLoanDTO) (*model.Master_Loan, error)
	UpdateLoanService(loan dto.UpdateLoanDTO) (*model.Master_Loan, error)
	SearchLoanByIdService(id uint64) (*model.Master_Loan, error)
	DeleteLoanService(id uint64) error
	UpdateLoanStatusService(customerId uint64) (*model.Master_Loan, error)
}

type loanService struct {
	loanRepo     repository.LoansRepository
	customerRepo repository.CustomerRepository
}

func NewLoanService(loanRepo repository.LoansRepository, customer repository.CustomerRepository) LoanService {
	return &loanService{
		loanRepo:     loanRepo,
		customerRepo: customer,
	}
}

func (ls *loanService) CreateLoanService(loan dto.CreateLoanDTO) (*model.Master_Loan, error) {
	loans := &model.Master_Loan{}
	err := smapping.FillStruct(&loans, smapping.MapFields(&loan))
	if err != nil {
		log.Printf("Error %v", err)
	}

	customerId := strconv.Itoa(int(loan.Customer_Id))

	serviceCustomer := NewCustomerService(ls.customerRepo)
	customer := serviceCustomer.ProfileCustomer(customerId)
	log.Printf("status verified %v", customer.StatusVerified)
	if !customer.StatusVerified {
		if loan.Amount > 500000 {
			return nil, errors.New("batas pinjaman untuk user yang belum terverifikasi adalah sebesar 500000")
		}
	} else {
		if loan.Amount > 10000000 {
			return nil, errors.New("batas pinjaman untuk user yang sudah terverifikasi adalah sebesar 10000000")
		}
	}

	err = ls.loanRepo.CreateLoanRepository(loans)
	if err != nil {
		return nil, err
	}
	return loans, nil
}

func (ls *loanService) UpdateLoanService(loan dto.UpdateLoanDTO) (*model.Master_Loan, error) {
	loans := &model.Master_Loan{}
	err := smapping.FillStruct(&loans, smapping.MapFields(&loan))
	if err != nil {
		log.Printf("Error map %v", err)
	}

	customerId := strconv.Itoa(int(loan.Customer_Id))

	serviceCustomer := NewCustomerService(ls.customerRepo)
	customer := serviceCustomer.ProfileCustomer(customerId)
	log.Printf("status verified %v", customer.StatusVerified)
	if !customer.StatusVerified {
		if loan.Amount > 500000 {
			return nil, errors.New("batas pinjaman untuk user yang belum terverifikasi adalah sebesar 500000")
		}
	} else {
		if loan.Amount > 10000000 {
			return nil, errors.New("batas pinjaman untuk user yang sudah terverifikasi adalah sebesar 10000000")
		}
	}

	err = ls.loanRepo.UpdateLoanRepository(loan.Id, loans)
	if err != nil {
		return nil, err
	}
	return loans, nil
}

func (ls *loanService) SearchLoanByIdService(id uint64) (*model.Master_Loan, error) {
	return ls.loanRepo.SearchLoanByIdRepository(id)
}

func (ls *loanService) DeleteLoanService(id uint64) error {
	return ls.loanRepo.DeleteLoanRepository(id)
}

func (ls *loanService) UpdateLoanStatusService(customerId uint64) (*model.Master_Loan, error) {
	loan, err := ls.loanRepo.UpdateLoanStatus(customerId)
	if err != nil {
		return nil, err
	}
	return loan, err
}
