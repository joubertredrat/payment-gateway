package domain

type PaginationCriteria struct {
	Page         uint
	ItemsPerPage uint
}

type CreditCardTransactionRepository interface {
	Create(c CreditCardTransaction) (CreditCardTransaction, error)
	Update(c CreditCardTransaction) (CreditCardTransaction, error)
	GetByTransactionID(TransactionID string) (CreditCardTransaction, error)
	GetList(p PaginationCriteria) ([]CreditCardTransaction, error)
}

type TransactionStatusRepository interface {
	Create(t TransactionStatus) (TransactionStatus, error)
	GetByID(ID string) (TransactionStatus, error)
	GetByCreditCardTransctionID(TransactionID string) ([]TransactionStatus, error)
}
