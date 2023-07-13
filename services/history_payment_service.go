package services

import (
	"golang-pinjol/model"
	"golang-pinjol/repository"
	"log"
)

type HistoryPaymentService interface {
	GetAllHistoriesPaymentService() ([]*model.Master_Payment_History, error)
	GetHistoryPaymentByIdService(id uint64) ([]*model.Master_Payment_History, error)
}

type historyPaymentService struct {
	historyRepository repository.HistoryPaymentRepository
}

func NewHistoryPaymentService(hps repository.HistoryPaymentRepository) HistoryPaymentService {
	return &historyPaymentService{
		historyRepository: hps,
	}
}

func (hp *historyPaymentService) GetAllHistoriesPaymentService() ([]*model.Master_Payment_History, error) {
	histories, err := hp.historyRepository.GetAllHistoriesPaymentRepository()
	if err != nil {
		return nil, err
	}
	return histories, nil
}

func (hp *historyPaymentService) GetHistoryPaymentByIdService(id uint64) ([]*model.Master_Payment_History, error) {
	history, err := hp.historyRepository.GetHistoryPaymentNasabahById(id)
	if err != nil {
		log.Printf("error history service %v", err)
	}
	return history, nil
}
