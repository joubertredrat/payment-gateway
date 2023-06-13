package infra

import (
	"database/sql"
	"joubertredrat/transaction-ms/internal/application"
	"joubertredrat/transaction-ms/internal/domain"
	"joubertredrat/transaction-ms/pkg"
)

type CreditCardTransactionRepositoryMysql struct {
	logger application.Logger
	db     *sql.DB
}

func NewCreditCardTransactionRepositoryMysql(logger application.Logger, db *sql.DB) CreditCardTransactionRepositoryMysql {
	return CreditCardTransactionRepositoryMysql{
		logger: logger,
		db:     db,
	}
}

func (r CreditCardTransactionRepositoryMysql) Create(c domain.CreditCardTransaction) (domain.CreditCardTransaction, error) {
	return domain.CreditCardTransaction{
		ID:            1,
		TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
		CardNumber:    "513073XXXXXX7489",
		Amount: domain.Amount{
			Value: 1250,
		},
		Installments:      2,
		Description:       "usb cable",
		TransactionStatus: []domain.TransactionStatus{},
		CreatedAt:         pkg.TimeFromCanonical("2023-06-10 17:00:00"),
	}, nil
}

func (r CreditCardTransactionRepositoryMysql) Update(c domain.CreditCardTransaction) (domain.CreditCardTransaction, error) {
	return domain.CreditCardTransaction{}, nil
}

func (r CreditCardTransactionRepositoryMysql) DeleteByTransactionID(TransactionID string) error {
	return nil
}

func (r CreditCardTransactionRepositoryMysql) GetByTransactionID(TransactionID string) (domain.CreditCardTransaction, error) {
	return domain.CreditCardTransaction{}, nil
}

func (r CreditCardTransactionRepositoryMysql) GetList(p domain.PaginationCriteria) ([]domain.CreditCardTransaction, error) {
	return []domain.CreditCardTransaction{
		{
			ID:            1,
			TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
			CardNumber:    "513073XXXXXX7489",
			Amount: domain.Amount{
				Value: 1250,
			},
			Installments:      2,
			Description:       "usb cable",
			TransactionStatus: []domain.TransactionStatus{},
			CreatedAt:         pkg.TimeFromCanonical("2023-06-10 17:00:00"),
		},
		{
			ID:            2,
			TransactionID: "01H2V8NJQK039S5TPB2FBWW58C",
			CardNumber:    "513073XXXXXX9915",
			Amount: domain.Amount{
				Value: 1725,
			},
			Installments:      1,
			Description:       "good things aaa",
			TransactionStatus: []domain.TransactionStatus{},
			CreatedAt:         pkg.TimeFromCanonical("2023-06-10 17:00:00"),
			UpdatedAt:         pkg.TimeFromCanonical("2023-06-11 17:16:00"),
		},
	}, nil
}

type TransactionStatusRepositoryMysql struct {
	logger application.Logger
	db     *sql.DB
}

func NewTransactionStatusRepositoryMysql(logger application.Logger, db *sql.DB) TransactionStatusRepositoryMysql {
	return TransactionStatusRepositoryMysql{
		logger: logger,
		db:     db,
	}
}

func (r TransactionStatusRepositoryMysql) Create(t domain.TransactionStatus) (domain.TransactionStatus, error) {
	return domain.TransactionStatus{
		ID:                      10,
		CreditCardTransactionID: 1,
		Status:                  domain.TRANSACTION_STATUS_CREATED,
		CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:01:00"),
	}, nil
}

func (r TransactionStatusRepositoryMysql) GetByCreditCardTransactionID(TransactionID string) ([]domain.TransactionStatus, error) {
	return []domain.TransactionStatus{}, nil
}
