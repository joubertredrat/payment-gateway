package infra

import "joubertredrat/transaction-ms/internal/domain"

type QueueDispatcher struct {
}

func NewQueueDispatcher() QueueDispatcher {
	return QueueDispatcher{}
}

func (d QueueDispatcher) CreditCardTransactionCreated(domain.CreditCardTransaction) error {
	return nil
}

func (d QueueDispatcher) CreditCardTransactionEdited(domain.CreditCardTransaction) error {
	return nil
}

func (d QueueDispatcher) CreditCardTransactionDeleted(TransactionID string) error {
	return nil
}

func (d QueueDispatcher) CreditCardTransactionGot(domain.CreditCardTransaction) error {
	return nil
}

func (d QueueDispatcher) CreditCardTransactionListed(domain.PaginationCriteria) error {
	return nil
}
