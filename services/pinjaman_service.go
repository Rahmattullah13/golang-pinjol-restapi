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
	CreateLoanService(loan dto.CreatePinjamanDTO) (*model.Master_Loan, error)
	UpdateLoanService(loan dto.UpdatePinjamanDTO) (*model.Master_Loan, error)
	SearchLoanByIdService(id uint64) (*model.Master_Loan, error)
	DeleteLoanService(id uint64) error
	UpdateLoanStatusService(nasabahId uint64) (*model.Master_Loan, error)
}

type loanService struct {
	loanRepo    repository.LoansRepository
	nasabahRepo repository.NasabahRepository
}

func NewLoanService(loanRepo repository.LoansRepository, nasabah repository.NasabahRepository) LoanService {
	return &loanService{
		loanRepo:    loanRepo,
		nasabahRepo: nasabah,
	}
}

func (ls *loanService) CreateLoanService(loan dto.CreatePinjamanDTO) (*model.Master_Loan, error) {
	loans := &model.Master_Loan{}
	err := smapping.FillStruct(&loans, smapping.MapFields(&loan))
	if err != nil {
		log.Printf("Error %v", err)
	}

	nasabahId := strconv.Itoa(int(loan.Nasabah_Id))

	serviceNasabah := NewNasabahService(ls.nasabahRepo)
	nasabah := serviceNasabah.ProfileNasabah(nasabahId)
	log.Printf("status verified %v", nasabah.StatusVerified)
	if !nasabah.StatusVerified {
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

func (ls *loanService) UpdateLoanService(loan dto.UpdatePinjamanDTO) (*model.Master_Loan, error) {
	loans := &model.Master_Loan{}
	err := smapping.FillStruct(&loans, smapping.MapFields(&loan))
	if err != nil {
		log.Printf("Error map %v", err)
	}

	nasabahId := strconv.Itoa(int(loan.Nasabah_Id))

	serviceNasabah := NewNasabahService(ls.nasabahRepo)
	nasabah := serviceNasabah.ProfileNasabah(nasabahId)
	log.Printf("status verified %v", nasabah.StatusVerified)
	if !nasabah.StatusVerified {
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

func (ls *loanService) UpdateLoanStatusService(nasabahId uint64) (*model.Master_Loan, error) {
	loan, err := ls.loanRepo.UpdateLoanStatus(nasabahId)
	if err != nil {
		return nil, err
	}
	return loan, err
}
