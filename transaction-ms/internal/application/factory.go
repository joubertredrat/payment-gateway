package application

import (
	"joubertredrat/transaction-ms/internal/domain"
)

func CreateCreditCardTransactionFromInput(input CreateCreditCardTransactionInput) (domain.CreditCardTransaction, error) {
	return domain.NewCreditCardTransaction(
		domain.TransactionID(),
		input.CardNumber,
		input.Amount,
		input.Installments,
		input.Description,
		[]domain.TransactionStatus{},
	)
}

func CreateAuthorizationRequestFromInput(input CreateCreditCardTransactionInput) (domain.AuthorizationRequest, error) {
	return domain.NewAuthorizationRequest(
		input.HolderName,
		input.CardNumber,
		input.CVV,
		input.ExpireDate,
		input.Amount,
		input.Installments,
	)
}
