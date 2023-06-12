package application

import (
	"joubertredrat/transaction-ms/internal/domain"
)

type UsecaseCreateCreditCardTransction struct {
	logger                          Logger
	creditCardTransactionRepository domain.CreditCardTransactionRepository
	transactionStatusRepository     domain.TransactionStatusRepository
	authorizationService            domain.AuthorizationService
	dispatcher                      Dispatcher
}

func NewUsecaseCreateCreditCardTransction(
	logger Logger,
	creditCardTransactionRepository domain.CreditCardTransactionRepository,
	transactionStatusRepository domain.TransactionStatusRepository,
	authorizationService domain.AuthorizationService,
	dispatcher Dispatcher,
) UsecaseCreateCreditCardTransction {
	return UsecaseCreateCreditCardTransction{
		logger:                          logger,
		creditCardTransactionRepository: creditCardTransactionRepository,
		transactionStatusRepository:     transactionStatusRepository,
		authorizationService:            authorizationService,
		dispatcher:                      dispatcher,
	}
}

func (u UsecaseCreateCreditCardTransction) Execute(input CreateCreditCardTransctionInput) (domain.CreditCardTransaction, error) {
	transaction, err := CreateCreditCardTransactionFromInput(input)
	if err != nil {
		u.logger.Error(err)
		return domain.CreditCardTransaction{}, err
	}
	authorizationRequest, err := CreateAuthorizationRequestFromInput(input)
	if err != nil {
		u.logger.Error(err)
		return domain.CreditCardTransaction{}, ErrUsecaseCreateCreditCardTransctionHouston
	}

	transactionCreated, err := u.creditCardTransactionRepository.Create(transaction)
	if err != nil {
		u.logger.Error(err)
		return domain.CreditCardTransaction{}, ErrUsecaseCreateCreditCardTransctionHouston
	}

	statusCreated, err := u.appendStatus(transactionCreated, domain.TRANSACTION_STATUS_CREATED)
	if err != nil {
		u.logger.Error(err)
		return domain.CreditCardTransaction{}, ErrUsecaseCreateCreditCardTransctionHouston
	}
	transactionCreated.TransactionStatus = append(transactionCreated.TransactionStatus, statusCreated)

	authorizationResponse, err := u.authorizationService.Handle(authorizationRequest)
	if err != nil {
		// status error
		u.logger.Error(err)
		return domain.CreditCardTransaction{}, ErrUsecaseCreateCreditCardTransctionHouston
	}

	statusAuthorization, err := u.appendStatus(transactionCreated, authorizationResponse.GetTransactionStatus())
	if err != nil {
		u.logger.Error(err)
		return domain.CreditCardTransaction{}, ErrUsecaseCreateCreditCardTransctionHouston
	}
	transactionCreated.TransactionStatus = append(transactionCreated.TransactionStatus, statusAuthorization)

	if err := u.dispatcher.CreditCardTransctionCreated(transactionCreated); err != nil {
		u.logger.Error(err)
	}
	return transactionCreated, nil
}

func (u UsecaseCreateCreditCardTransction) appendStatus(t domain.CreditCardTransaction, status string) (domain.TransactionStatus, error) {
	statusCreated, err := domain.NewTransactionStatus(t.ID, status)
	if err != nil {
		return domain.TransactionStatus{}, err
	}

	return u.transactionStatusRepository.Create(statusCreated)
}
