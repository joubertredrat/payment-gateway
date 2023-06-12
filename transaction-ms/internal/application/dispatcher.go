package application

import "joubertredrat/transaction-ms/internal/domain"

type Dispatcher interface {
	CreditCardTransctionCreated(domain.CreditCardTransaction) error
}
