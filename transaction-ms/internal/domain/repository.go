package domain

type PaginationCriteria struct {
	Page         uint
	ItemsPerPage uint
}

type CreditCardTransctionRepository interface {
	Create(c CreditCardTransction) (CreditCardTransction, error)
	GetByID(ID string) (CreditCardTransction, error)
	GetList(p PaginationCriteria) ([]CreditCardTransction, error)
}

type TransctionStatusRepository interface {
	Create(t TransctionStatus) (TransctionStatus, error)
	GetByID(ID string) (TransctionStatus, error)
	GetByCreditCardTransctionID(ID string) ([]TransctionStatus, error)
}
