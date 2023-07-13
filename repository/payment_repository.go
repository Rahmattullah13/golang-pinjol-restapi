package repository

import (
	"fmt"
	"golang-pinjol/model"
	"log"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	CreatePaymentRepository(payment *model.Transactions_Payment_Loan) (*model.Transactions_Payment_Loan, error)
	UpdatePaymentRepository(id int, payment *model.Transactions_Payment_Loan) (*model.Transactions_Payment_Loan, error)
	FindsPaymentByIdRepository(id int) (*model.Transactions_Payment_Loan, error)
	ListPaymentRepository() ([]*model.Transactions_Payment_Loan, error)
	GetPaymentPerMonthRepository(loansID int) (int, error)
	GetTotalPaymentsRepository(loansID int) (int, error)
	DeletePaymentRepository(id int) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}

func (db *paymentRepository) CreatePaymentRepository(payment *model.Transactions_Payment_Loan) (*model.Transactions_Payment_Loan, error) {
	if err := db.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(payment).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.Master_Loan{}).Where("id", payment.Loan_id).UpdateColumn("amount", gorm.Expr("amount - ?", payment.Monthly_Payments)).Error; err != nil {
			return nil
		}

		go func() {
			history := &model.Master_Payment_History{
				Loan_id:     payment.Loan_id,
				Payment_id:  payment.ID,
				Date:        payment.Payment_Date,
				Loan:        payment.Loan,
				Transaction: *payment,
			}
			histories := NewHistoryPaymentRepository(db.db)
			if err := histories.CreateHistoryPaymentRepository(history); err != nil {
				log.Println(err)
			}
		}()
		return nil
	}); err != nil {
		return nil, err
	}
	return payment, nil
}

func (db *paymentRepository) UpdatePaymentRepository(id int, payment *model.Transactions_Payment_Loan) (*model.Transactions_Payment_Loan, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID is not valid : %v", id)
	}

	var currentPayment model.Transactions_Payment_Loan
	err := db.db.First(&currentPayment, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("record not found with id %d", id)
		}
		return nil, err
	}

	err = db.db.Model(&currentPayment).Updates(payment).Error
	if err != nil {
		return nil, err
	}
	return &currentPayment, nil
}

func (db *paymentRepository) FindsPaymentByIdRepository(id int) (*model.Transactions_Payment_Loan, error) {
	var payment model.Transactions_Payment_Loan
	if err := db.db.First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (db *paymentRepository) ListPaymentRepository() ([]*model.Transactions_Payment_Loan, error) {
	var payments []*model.Transactions_Payment_Loan
	if err := db.db.Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (db *paymentRepository) GetPaymentPerMonthRepository(loanID int) (int, error) {
	var loan model.Master_Loan
	if err := db.db.First(&loan, loanID).Error; err != nil {
		return 0, err
	}

	loansAmount := loan.Amount
	interestRate := loan.Loan_Interest_Rates
	loanDuration := loan.Loan_Duration
	paymentPerMonth := (loansAmount*interestRate)/(12*100) + (loansAmount / loanDuration)

	return paymentPerMonth, nil
}

func (db *paymentRepository) GetTotalPaymentsRepository(loanID int) (int, error) {
	var loan model.Master_Loan
	if err := db.db.First(&loan, loanID).Error; err != nil {
		return 0, err
	}

	loansAmount := loan.Amount
	interestRate := loan.Loan_Interest_Rates
	loanDuration := loan.Loan_Duration
	paymentPerMonth := (loansAmount*interestRate)/(12*100) + (loansAmount / loanDuration)
	totalPayments := paymentPerMonth * loanDuration

	return totalPayments, nil
}

func (db *paymentRepository) DeletePaymentRepository(id int) error {
	if err := db.db.Where("id = $1", id).Delete(&model.Transactions_Payment_Loan{}).Error; err != nil {
		return err
	}
	return nil
}
