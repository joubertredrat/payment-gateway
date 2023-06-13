package infra

import (
	"database/sql"
	"fmt"
	"joubertredrat/transaction-ms/internal/application"
	"joubertredrat/transaction-ms/internal/domain"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func GetMysqlDSN(host, port, name, user, password string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, name)
}

func GetDatabaseConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}

func GetLogrusStdout() *logrus.Logger {
	log := logrus.New()
	log.Out = os.Stdout

	return log
}

func GetLoggerStdout(logger *logrus.Logger) application.Logger {
	return NewLoggerStdout(logger)
}

func GetCreditCardTransactionRepository(logger application.Logger, db *sql.DB) domain.CreditCardTransactionRepository {
	return NewCreditCardTransactionRepositoryMysql(logger, db)
}

func GetTransactionStatusRepository(logger application.Logger, db *sql.DB) domain.TransactionStatusRepository {
	return NewTransactionStatusRepositoryMysql(logger, db)
}

func GetQueueDispatcher() application.Dispatcher {
	return NewQueueDispatcher()
}

func GetAuthorizationServiceMicroservice() domain.AuthorizationService {
	return NewAuthorizationServiceMicroservice()
}

func GetUsecaseCreateCreditCardTransaction(
	l application.Logger,
	cr domain.CreditCardTransactionRepository,
	tr domain.TransactionStatusRepository,
	as domain.AuthorizationService,
	d application.Dispatcher,
) application.UsecaseCreateCreditCardTransaction {
	return application.NewUsecaseCreateCreditCardTransaction(l, cr, tr, as, d)
}

func GetUsecaseEditCreditCardTransaction(
	l application.Logger,
	cr domain.CreditCardTransactionRepository,
	tr domain.TransactionStatusRepository,
	d application.Dispatcher,
) application.UsecaseEditCreditCardTransaction {
	return application.NewUsecaseEditCreditCardTransaction(l, cr, tr, d)
}

func GetUsecaseDeleteCreditCardTransaction(
	l application.Logger,
	cr domain.CreditCardTransactionRepository,
	d application.Dispatcher,
) application.UsecaseDeleteCreditCardTransaction {
	return application.NewUsecaseDeleteCreditCardTransaction(l, cr, d)
}

func GetUsecaseGetCreditCardTransaction(
	l application.Logger,
	cr domain.CreditCardTransactionRepository,
	tr domain.TransactionStatusRepository,
	d application.Dispatcher,
) application.UsecaseGetCreditCardTransaction {
	return application.NewUsecaseGetCreditCardTransaction(l, cr, tr, d)
}

func GetUsecaseListCreditCardTransaction(
	l application.Logger,
	cr domain.CreditCardTransactionRepository,
	d application.Dispatcher,
) application.UsecaseListCreditCardTransaction {
	return application.NewUsecaseListCreditCardTransaction(l, cr, d)
}
