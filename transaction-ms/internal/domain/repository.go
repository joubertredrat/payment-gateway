package domain

const (
	ITEMS_PER_PAGE_MIN = 10
	ITEMS_PER_PAGE_MAX = 100
)

type PaginationCriteria struct {
	Page         uint
	ItemsPerPage uint
}

func NewPaginationCriteria(page, itemsPerPage uint) (PaginationCriteria, error) {
	if page < 1 {
		return PaginationCriteria{}, NewErrPaginationCriteriaPage(page)
	}
	if !IsValidItemsPerPage(itemsPerPage) {
		return PaginationCriteria{}, NewErrPaginationCriteriaItemsPerPage(itemsPerPage)
	}

	return PaginationCriteria{
		Page:         page,
		ItemsPerPage: itemsPerPage,
	}, nil
}

func IsValidItemsPerPage(i uint) bool {
	return i >= ITEMS_PER_PAGE_MIN && i <= ITEMS_PER_PAGE_MAX
}

type CreditCardTransactionRepository interface {
	Create(c CreditCardTransaction) (CreditCardTransaction, error)
	Update(c CreditCardTransaction) (CreditCardTransaction, error)
	DeleteByTransactionID(TransactionID string) error
	GetByTransactionID(TransactionID string) (CreditCardTransaction, error)
	GetList(p PaginationCriteria) ([]CreditCardTransaction, error)
}

type TransactionStatusRepository interface {
	Create(t TransactionStatus) (TransactionStatus, error)
	GetByCreditCardTransactionID(TransactionID string) ([]TransactionStatus, error)
}
