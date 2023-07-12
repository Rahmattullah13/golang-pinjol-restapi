package services

import (
	"fmt"
	"golang-pinjol/dto"
	"golang-pinjol/model"
	"golang-pinjol/repository"
	"log"
	"strconv"
	"time"

	"github.com/mashingan/smapping"
)

type PaymentService interface {
	PaymentLoanService(payment *dto.CreatePaymentDTO) (*model.Transactions_Payment_Loan, error)
	UpdatePaymentService(updatePayment dto.UpdatePaymentDTO) (*model.Transactions_Payment_Loan, error)
	ListPaymentByStatusService(status string) ([]*model.Transactions_Payment_Loan, error)
	GetPaymentPerMonthService(loanID int) (int, error)
	GetTotalPaymentService(loan_id int) (int, error)
	DeletePaymentService(id int) error
}

type paymentService struct {
	paymentRepository repository.PaymentRepository
}

func NewPaymentService(paymentRepo repository.PaymentRepository) PaymentService {
	return &paymentService{
		paymentRepository: paymentRepo,
	}
}

func (ps *paymentService) PaymentLoanService(payment *dto.CreatePaymentDTO) (*model.Transactions_Payment_Loan, error) {
	var txLoans model.Transactions_Payment_Loan
	err := smapping.FillStruct(&txLoans, smapping.MapFields(payment))
	if err != nil {
		log.Printf("Error map %v : ", err)
	}

	currentMonth := time.Now().Month()
	if currentMonth == txLoans.Payment_Date.Month() {
		txLoans.Payment_Status = true
	}

	txLoans.Payment_Date = time.Now()

	payments, err := ps.paymentRepository.CreatePaymentRepository(&txLoans)
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (ps *paymentService) UpdatePaymentService(updatePayment dto.UpdatePaymentDTO) (*model.Transactions_Payment_Loan, error) {
	var txLoan model.Transactions_Payment_Loan
	err := smapping.FillStruct(&txLoan, smapping.MapFields(&updatePayment))
	if err != nil {
		return nil, fmt.Errorf("Error mapping input : %v", err)
	}

	updatedLoan, err := ps.paymentRepository.UpdatePaymentRepository(updatePayment.Id, &txLoan)
	if err != nil {
		return nil, fmt.Errorf("Error updating payment : %v", err)
	}

	return updatedLoan, nil
}

func (ps *paymentService) ListPaymentByStatusService(status string) ([]*model.Transactions_Payment_Loan, error) {
	payments, err := ps.paymentRepository.ListPaymentRepository()
	if err != nil {
		return nil, err
	}

	var filteredPayments []*model.Transactions_Payment_Loan
	for _, p := range payments {
		if strconv.FormatBool(p.Payment_Status) == status {
			filteredPayments = append(filteredPayments, p)
		}
	}
	return filteredPayments, nil
}

func (ps *paymentService) GetPaymentPerMonthService(loanID int) (int, error) {
	return ps.paymentRepository.GetPaymentPerMonthRepository(loanID)
}

func (ps *paymentService) GetTotalPaymentService(loan_id int) (int, error) {
	totalPayments, err := ps.paymentRepository.GetTotalPaymentsRepository(loan_id)
	if err != nil {
		return 0, err
	}
	return totalPayments, nil
}

func (ps *paymentService) DeletePaymentService(id int) error {
	return ps.paymentRepository.DeletePaymentRepository(id)
}
