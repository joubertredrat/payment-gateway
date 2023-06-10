package application

import (
	"errors"
	"joubertredrat/transaction-ms/internal/domain"
)

type UsecaseCreateCreditCardTransction struct {
	logger                          Logger
	creditCardTransactionRepository domain.CreditCardTransactionRepository
	transactionStatusRepository     domain.TransactionStatusRepository
	authorizationService            domain.AuthorizationService
}

func NewUsecaseCreateCreditCardTransction(
	logger Logger,
	creditCardTransactionRepository domain.CreditCardTransactionRepository,
	transactionStatusRepository domain.TransactionStatusRepository,
	authorizationService domain.AuthorizationService,
) UsecaseCreateCreditCardTransction {
	return UsecaseCreateCreditCardTransction{
		logger:                          logger,
		creditCardTransactionRepository: creditCardTransactionRepository,
		transactionStatusRepository:     transactionStatusRepository,
		authorizationService:            authorizationService,
	}
}

func (u UsecaseCreateCreditCardTransction) Execute(input CreateCreditCardTransctionInput) (domain.CreditCardTransaction, error) {
	transaction, err := CreateCreditCardTransactionFromInput(input)
	if err != nil {
		return domain.CreditCardTransaction{}, errors.New("Use case error")
	}

	transactionCreated, err := u.creditCardTransactionRepository.Create(transaction)
	if err != nil {
		return domain.CreditCardTransaction{}, errors.New("Use case error")
	}

	statusCreated, err := u.appendStatus(transactionCreated, domain.TRANSACTION_STATUS_CREATED)
	if err != nil {
		return domain.CreditCardTransaction{}, errors.New("Use case error")
	}
	transactionCreated.TransactionStatus = append(transactionCreated.TransactionStatus, statusCreated)

	authorizationRequest, err := CreateAuthorizationRequestFromInput(input)
	if err != nil {
		return domain.CreditCardTransaction{}, errors.New("Use case error")
	}
	authorizationResponse, err := u.authorizationService.Handle(authorizationRequest)
	if err != nil {
		// status error
		return domain.CreditCardTransaction{}, errors.New("Use case error")
	}

	statusAuthorization, err := u.appendStatus(transactionCreated, authorizationResponse.GetTransactionStatus())
	if err != nil {
		return domain.CreditCardTransaction{}, errors.New("Use case error")
	}
	transactionCreated.TransactionStatus = append(transactionCreated.TransactionStatus, statusAuthorization)

	return transactionCreated, nil
}

func (u UsecaseCreateCreditCardTransction) appendStatus(t domain.CreditCardTransaction, status string) (domain.TransactionStatus, error) {
	statusCreated, err := domain.NewTransactionStatus(t.ID, status)
	if err != nil {
		return domain.TransactionStatus{}, err
	}

	return u.transactionStatusRepository.Create(statusCreated)
}
