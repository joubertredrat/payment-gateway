package application

import "joubertredrat/transaction-ms/internal/domain"

type Dispatcher interface {
	CreditCardTransactionCreated(domain.CreditCardTransaction) error
	CreditCardTransactionEdited(domain.CreditCardTransaction) error
	CreditCardTransactionDeleted(TransactionID string) error
	CreditCardTransactionGot(domain.CreditCardTransaction) error
	CreditCardTransactionListed(domain.PaginationCriteria) error
}
