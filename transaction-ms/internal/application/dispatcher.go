package application

import "joubertredrat/transaction-ms/internal/domain"

type Dispatcher interface {
	CreditCardTransactionCreated(domain.CreditCardTransaction) error
	CreditCardTransactionEdited(domain.CreditCardTransaction) error
}
