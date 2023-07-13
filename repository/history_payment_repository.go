package repository

import (
	"errors"
	"golang-pinjol/model"
	"log"

	"gorm.io/gorm"
)

type HistoryPaymentRepository interface {
	CreateHistoryPaymentRepository(history *model.Master_Payment_History) error
	GetAllHistoriesPaymentRepository() ([]*model.Master_Payment_History, error)
	GetHistoryPaymentByIdRepository(id uint64) (*model.Master_Payment_History, error)
	GetHistoryPaymentNasabahById(id uint64) ([]*model.Master_Payment_History, error)
}

type historyPaymentRepository struct {
	db *gorm.DB
}

func NewHistoryPaymentRepository(db *gorm.DB) HistoryPaymentRepository {
	return &historyPaymentRepository{
		db: db,
	}
}

func (db *historyPaymentRepository) CreateHistoryPaymentRepository(history *model.Master_Payment_History) error {
	if err := db.db.Create(history).Error; err != nil {
		return err
	}
	return nil
}

func (db *historyPaymentRepository) GetAllHistoriesPaymentRepository() ([]*model.Master_Payment_History, error) {
	var historyPayments []*model.Master_Payment_History
	err := db.db.Table("master_payment_histories").Select("master_payment_histories.*, master_loans.nasabah_id, transactions_payment_loans.payment_status").Joins("JOIN transactions_payment_loans ON transactions_payment_loans.id = master_payment_histories.payment_id").Joins("JOIN master_loans ON master_loans.id = master_payment_histories.loan_id").Scan(&historyPayments).Error

	if err != nil {
		log.Println(err)
		return nil, err
	}
	if len(historyPayments) == 0 {
		log.Println("Data NOT FOUND")
		return nil, errors.New("Data NOT FOUND")
	}
	return historyPayments, nil
}

func (db *historyPaymentRepository) GetHistoryPaymentByIdRepository(id uint64) (*model.Master_Payment_History, error) {
	var history model.Master_Payment_History
	if err := db.db.First(&history, id).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

func (db *historyPaymentRepository) GetHistoryPaymentNasabahById(id uint64) ([]*model.Master_Payment_History, error) {
	var historyPayment []*model.Master_Payment_History
	err := db.db.Table("master_payment_histories").Select("master_payment_histories.*, master_loans.nasabah_id, transactions_payment_loans.payment_status").Joins("JOIN transactions_payment_loans ON transactions_payment_loans.id = master_payment_histories.payment_id").Joins("JOIN master_loans ON master_loans.id = master_payment_histories.loan_id").Where("master_payment_histories.id = $1", id).Scan(&historyPayment).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if historyPayment[0].Id == 0 {
		log.Println("Data NOT FOUND")
		return nil, errors.New("Data NOT FOUND")
	}
	return historyPayment, nil
}
