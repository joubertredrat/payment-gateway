package application

import "joubertredrat/transaction-ms/internal/domain"

type Dispatcher interface {
	CreditCardTransactionCreated(c domain.CreditCardTransaction) error
	CreditCardTransactionEdited(c domain.CreditCardTransaction) error
	CreditCardTransactionDeleted(TransactionID string) error
	CreditCardTransactionGot(c domain.CreditCardTransaction) error
	CreditCardTransactionListed(p domain.PaginationCriteria) error
}
