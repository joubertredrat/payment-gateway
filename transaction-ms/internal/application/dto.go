package application

import "time"

type CreateCreditCardTransctionInput struct {
	HolderName   string
	CardNumber   string
	CVV          string
	ExpireDate   time.Time
	Amount       uint
	Installments uint
	Description  string
}

type EditCreditCardTransctionInput struct {
	TransactionID string
	Description   string
}
