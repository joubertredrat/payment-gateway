package infra

import (
	"database/sql"
	"fmt"
	"joubertredrat/transaction-ms/internal/application"
	"joubertredrat/transaction-ms/internal/domain"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func GetMysqlDSN(host, port, name, user, password string) string {
	return fmt.Sprintf("%s:%s@%s:%s/%s", user, password, host, port, name)
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
