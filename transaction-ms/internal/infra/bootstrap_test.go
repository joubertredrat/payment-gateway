package infra_test

import (
	"joubertredrat/transaction-ms/internal/infra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMysqlDSN(t *testing.T) {
	dsnExpected := "dbuser:ae23f678-5a34-469f-90a3-d13985a27ea6@tcp(127.0.0.1:3306)/transactions_ms?charset=utf8"
	dsnGot := infra.GetMysqlDSN("127.0.0.1", "3306", "transactions_ms", "dbuser", "ae23f678-5a34-469f-90a3-d13985a27ea6")

	assert.Equal(t, dsnExpected, dsnGot)
}
