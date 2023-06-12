package application

import (
	"joubertredrat/transaction-ms/internal/domain"
)

type UsecaseCreateCreditCardTransaction struct {
	logger                          Logger
	creditCardTransactionRepository domain.CreditCardTransactionRepository
	transactionStatusRepository     domain.TransactionStatusRepository
	authorizationService            domain.AuthorizationService
	dispatcher                      Dispatcher
}

func NewUsecaseCreateCreditCardTransaction(
	logger Logger,
	creditCardTransactionRepository domain.CreditCardTransactionRepository,
	transactionStatusRepository domain.TransactionStatusRepository,
	authorizationService domain.AuthorizationService,
	dispatcher Dispatcher,
) UsecaseCreateCreditCardTransaction {
	return UsecaseCreateCreditCardTransaction{
		logger:                          logger,
		creditCardTransactionRepository: creditCardTransactionRepository,
		transactionStatusRepository:     transactionStatusRepository,
		authorizationService:            authorizationService,
		dispatcher:                      dispatcher,
	}
}

func (u UsecaseCreateCreditCardTransaction) Execute(input CreateCreditCardTransactionInput) (domain.CreditCardTransaction, error) {
	transaction, err := CreateCreditCardTransactionFromInput(input)
	if err != nil {
		u.logger.Error(err)
		return domain.CreditCardTransaction{}, err
	}
	authorizationRequest, err := CreateAuthorizationRequestFromInput(input)
	if err != nil {
		u.logger.Error(err)
		return domain.CreditCardTransaction{}, err
	}

	transactionCreated, err := u.creditCardTransactionRepository.Create(transaction)
	if err != nil {
		u.logger.Error(err)
		return domain.CreditCardTransaction{}, ErrUsecaseCreateCreditCardTransactionHouston
	}

	statusCreated, err := u.appendStatus(transactionCreated, domain.TRANSACTION_STATUS_CREATED)
	if err != nil {
		u.logger.Error(err)
		return domain.CreditCardTransaction{}, ErrUsecaseCreateCreditCardTransactionHouston
	}
	transactionCreated.TransactionStatus = append(transactionCreated.TransactionStatus, statusCreated)

	authorizationResponse, err := u.authorizationService.Handle(authorizationRequest)
	if err != nil {
		u.logger.Error(err)
		u.appendErrorStatus(transactionCreated, domain.TRANSACTION_STATUS_ERROR_AUTHORIZATION)
		return domain.CreditCardTransaction{}, ErrUsecaseCreateCreditCardTransactionHouston
	}

	statusAuthorization, err := u.appendStatus(transactionCreated, authorizationResponse.GetTransactionStatus())
	if err != nil {
		u.logger.Error(err)
		return domain.CreditCardTransaction{}, ErrUsecaseCreateCreditCardTransactionHouston
	}
	transactionCreated.TransactionStatus = append(transactionCreated.TransactionStatus, statusAuthorization)

	if err := u.dispatcher.CreditCardTransactionCreated(transactionCreated); err != nil {
		u.logger.Error(err)
	}
	return transactionCreated, nil
}

func (u UsecaseCreateCreditCardTransaction) appendStatus(t domain.CreditCardTransaction, status string) (domain.TransactionStatus, error) {
	statusCreated, err := domain.NewTransactionStatus(t.ID, status)
	if err != nil {
		return domain.TransactionStatus{}, err
	}

	return u.transactionStatusRepository.Create(statusCreated)
}

func (u UsecaseCreateCreditCardTransaction) appendErrorStatus(t domain.CreditCardTransaction, status string) {
	_, err := u.appendStatus(t, status)
	if err != nil {
		u.logger.Error(err)
	}
}

type UsecaseEditCreditCardTransaction struct {
	logger                          Logger
	creditCardTransactionRepository domain.CreditCardTransactionRepository
	transactionStatusRepository     domain.TransactionStatusRepository
	dispatcher                      Dispatcher
}

func NewUsecaseEditCreditCardTransaction(
	logger Logger,
	creditCardTransactionRepository domain.CreditCardTransactionRepository,
	transactionStatusRepository domain.TransactionStatusRepository,
	dispatcher Dispatcher,
) UsecaseEditCreditCardTransaction {
	return UsecaseEditCreditCardTransaction{
		logger:                          logger,
		creditCardTransactionRepository: creditCardTransactionRepository,
		transactionStatusRepository:     transactionStatusRepository,
		dispatcher:                      dispatcher,
	}
}

func (u UsecaseEditCreditCardTransaction) Execute(input EditCreditCardTransactionInput) (domain.CreditCardTransaction, error) {
	transactionFound, err := u.creditCardTransactionRepository.GetByTransactionID(input.TransactionID)
	if err != nil {
		u.logger.Error(err)
		if _, ok := err.(domain.ErrCreditCardTransactionNotFound); ok {
			return domain.CreditCardTransaction{}, err
		}

		return domain.CreditCardTransaction{}, ErrUsecaseEditCreditCardTransactionHouston
	}

	transactionFound.Description = input.Description
	transactionEdited, err := u.creditCardTransactionRepository.Update(transactionFound)
	if err != nil {
		u.logger.Error(err)
		return domain.CreditCardTransaction{}, ErrUsecaseEditCreditCardTransactionHouston
	}

	status, err := u.transactionStatusRepository.GetByCreditCardTransactionID(transactionEdited.TransactionID)
	if err != nil {
		u.logger.Error(err)
		return domain.CreditCardTransaction{}, ErrUsecaseEditCreditCardTransactionHouston
	}
	transactionEdited.TransactionStatus = status

	if err := u.dispatcher.CreditCardTransactionEdited(transactionEdited); err != nil {
		u.logger.Error(err)
	}
	return transactionEdited, nil
}
