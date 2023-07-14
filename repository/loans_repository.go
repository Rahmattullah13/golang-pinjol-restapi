package repository

import (
	"errors"
	"golang-pinjol/model"

	"gorm.io/gorm"
)

type LoansRepository interface {
	CreateLoanRepository(loan *model.Master_Loan) error
	UpdateLoanRepository(id uint64, loan *model.Master_Loan) error
	SearchLoanByIdRepository(id uint64) (*model.Master_Loan, error)
	DeleteLoanRepository(id uint64) error
	UpdateLoanStatus(customerID uint64) (*model.Master_Loan, error)
}

type loansRepository struct {
	db *gorm.DB
}

func NewLoansRepository(db *gorm.DB) LoansRepository {
	return &loansRepository{
		db: db,
	}
}

func (db *loansRepository) CreateLoanRepository(loan *model.Master_Loan) error {
	if err := db.db.Create(loan).Error; err != nil {
		return nil
	}
	return nil
}

func (db *loansRepository) UpdateLoanRepository(id uint64, loan *model.Master_Loan) error {
	if err := db.db.Model(model.Master_Loan{Id: id}).Updates(loan).Error; err != nil {
		return err
	}
	return nil
}

func (db *loansRepository) SearchLoanByIdRepository(id uint64) (*model.Master_Loan, error) {
	var loan model.Master_Loan
	if err := db.db.First(&loan, id).Error; err != nil {
		return nil, err
	}
	return &loan, nil
}

func (db *loansRepository) DeleteLoanRepository(id uint64) error {
	if err := db.db.Where("id = $1", id).Delete(&model.Master_Loan{}).Error; err != nil {
		return nil
	}
	return nil
}

func (db *loansRepository) UpdateLoanStatus(customerID uint64) (*model.Master_Loan, error) {
	var customer model.Master_Customer
	if err := db.db.Where("id = $1", customerID).First(&customer).Error; err != nil {
		return nil, err
	}

	if customer.StatusVerified {
		var loan model.Master_Loan
		if err := db.db.Where("customer_id = $1", customerID).First(&loan).Error; err != nil {
			return nil, err
		}

		loan.StatusApproved = true
		if err := db.db.Save(&loan).Error; err != nil {
			return nil, err
		}
		return &loan, nil
	}
	return nil, errors.New("customer status is not verified")
}
