package infra

import (
	"context"
	"database/sql"
	"joubertredrat/transaction-ms/internal/application"
	"joubertredrat/transaction-ms/internal/domain"
	"joubertredrat/transaction-ms/pkg"
	"time"
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
	createdAt := time.Now()
	c.CreatedAt = &createdAt

	insertResult, err := r.db.ExecContext(
		context.Background(),
		`INSERT INTO creditcard_transactions (
			transaction_id,
			card_number,
			amount,
			installments,
			description,
			created_at
		) VALUES (?, ?, ?, ?, ?, ?)`,
		c.TransactionID,
		c.CardNumber,
		c.Amount.Value,
		c.Installments,
		c.Description,
		DatetimeCanonical(c.CreatedAt),
	)
	if err != nil {
		r.logger.Error(err)
		return domain.CreditCardTransaction{}, domain.ErrCreditCardTransactionRepositoryHouston
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		r.logger.Error(err)
		return domain.CreditCardTransaction{}, domain.ErrCreditCardTransactionRepositoryHouston
	}
	c.ID = uint(id)

	return c, nil
}

func (r CreditCardTransactionRepositoryMysql) Update(c domain.CreditCardTransaction) (domain.CreditCardTransaction, error) {
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

func (r CreditCardTransactionRepositoryMysql) DeleteByTransactionID(TransactionID string) error {
	return nil
}

func (r CreditCardTransactionRepositoryMysql) GetByTransactionID(TransactionID string) (domain.CreditCardTransaction, error) {
	stmt, err := r.db.PrepareContext(
		context.Background(),
		`SELECT
			id,
			transaction_id,
			card_number,
			amount,
			installments,
			description,
			created_at,
			updated_at
		FROM creditcard_transactions
		WHERE transaction_id = ? AND deleted_at IS NULL`,
	)
	if err != nil {
		r.logger.Error(err)
		return domain.CreditCardTransaction{}, domain.ErrCreditCardTransactionRepositoryHouston
	}
	defer stmt.Close()

	var c domain.CreditCardTransaction

	row := stmt.QueryRowContext(context.Background(), TransactionID)
	errs := row.Scan(
		&c.ID,
		&c.TransactionID,
		&c.CardNumber,
		&c.Amount.Value,
		&c.Installments,
		&c.Description,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	if errs != nil {
		r.logger.Error(errs)
		switch {
		case errs == sql.ErrNoRows:
			return domain.CreditCardTransaction{}, domain.NewErrCreditCardTransactionNotFound("TransactionID", TransactionID)
		}
		return domain.CreditCardTransaction{}, domain.ErrCreditCardTransactionRepositoryHouston
	}

	return c, nil
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
	createdAt := time.Now()
	t.CreatedAt = &createdAt

	insertResult, err := r.db.ExecContext(
		context.Background(),
		`INSERT INTO creditcard_transactions_status (
			creditcard_transaction_id,
			status,
			created_at
		) VALUES (?, ?, ?)`,
		t.CreditCardTransactionID,
		t.Status,
		DatetimeCanonical(t.CreatedAt),
	)
	if err != nil {
		r.logger.Error(err)
		return domain.TransactionStatus{}, domain.ErrTransactionStatusRepositoryHouston
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		r.logger.Error(err)
		return domain.TransactionStatus{}, domain.ErrTransactionStatusRepositoryHouston
	}
	t.ID = uint(id)

	return t, nil
}

func (r TransactionStatusRepositoryMysql) GetByCreditCardTransactionID(TransactionID string) ([]domain.TransactionStatus, error) {
	stmt, err := r.db.PrepareContext(
		context.Background(),
		`SELECT
			t.id,
			t.creditcard_transaction_id,
			t.status,
			t.created_at
		FROM creditcard_transactions_status t
		JOIN creditcard_transactions c ON t.creditcard_transaction_id = c.id
		WHERE c.transaction_id = ?
		ORDER BY t.created_at DESC`,
	)
	if err != nil {
		r.logger.Error(err)
		return []domain.TransactionStatus{}, domain.ErrTransactionStatusRepositoryHouston
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(context.Background(), TransactionID)
	if err != nil {
		r.logger.Error(err)
		return []domain.TransactionStatus{}, domain.ErrTransactionStatusRepositoryHouston
	}
	defer rows.Close()

	status := []domain.TransactionStatus{}

	for rows.Next() {
		var s domain.TransactionStatus
		if err := rows.Scan(&s.ID, &s.CreditCardTransactionID, &s.Status, &s.CreatedAt); err != nil {
			r.logger.Error(err)
			return []domain.TransactionStatus{}, domain.ErrTransactionStatusRepositoryHouston
		}
		status = append(status, s)
	}
	if err := rows.Err(); err != nil {
		r.logger.Error(err)
		return []domain.TransactionStatus{}, domain.ErrTransactionStatusRepositoryHouston
	}

	return status, nil
}
