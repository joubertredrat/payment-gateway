package infra

import (
	"database/sql"
	"joubertredrat/transaction-ms/internal/application"
	"joubertredrat/transaction-ms/internal/domain"
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
	return domain.CreditCardTransaction{}, nil
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
	return []domain.CreditCardTransaction{}, nil
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
	return domain.TransactionStatus{}, nil
}

func (r TransactionStatusRepositoryMysql) GetByCreditCardTransactionID(TransactionID string) ([]domain.TransactionStatus, error) {
	return []domain.TransactionStatus{}, nil
}
