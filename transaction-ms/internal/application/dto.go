package application

import "time"

type CreateCreditCardTransactionInput struct {
	HolderName   string
	CardNumber   string
	CVV          string
	ExpireDate   time.Time
	Amount       uint
	Installments uint
	Description  string
}

type EditCreditCardTransactionInput struct {
	TransactionID string
	Description   string
}

type DeleteCreditCardTransactionInput struct {
	TransactionID string
}

type GetCreditCardTransactionInput struct {
	TransactionID string
}
