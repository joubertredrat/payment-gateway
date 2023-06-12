package domain_test

import (
	"fmt"
	"joubertredrat/transaction-ms/internal/domain"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrCreditCardTransactionNotFound(t *testing.T) {
	invalidCriteria := "TransactionID"
	invalidValue := "01H2RFARETAMR3G3HTZCDPFZ16"
	errExpected := fmt.Sprintf("Credit card transaction not found by criteria [ %s ] and value [ %s ]", invalidCriteria, invalidValue)
	errGot := domain.NewErrCreditCardTransactionNotFound(invalidCriteria, invalidValue)

	assert.Equal(t, errExpected, errGot.Error())
}

func TestErrInvalidCreditCardNumber(t *testing.T) {
	invalidNumber := "123456"
	errExpected := fmt.Sprintf("Invalid credit card number [ %s ]", invalidNumber)
	errGot := domain.NewErrInvalidCreditCardNumber(invalidNumber)

	assert.Equal(t, errExpected, errGot.Error())
}

func TestErrCreditCardTransactionInstallments(t *testing.T) {
	invalidInstallments := uint(16)
	errExpected := fmt.Sprintf(
		"Invalid installments, expected between [ %d ] and [ %d ], got [ %d ]",
		domain.INSTALLMENTS_MIN,
		domain.INSTALLMENTS_MAX,
		invalidInstallments,
	)
	errGot := domain.NewErrCreditCardTransactionInstallments(invalidInstallments)

	assert.Equal(t, errExpected, errGot.Error())
}

func TestErrTransactionStatusInvalid(t *testing.T) {
	invalidStatus := "done"
	errExpected := fmt.Sprintf(
		"Invalid transaction status, expected one of [ %s ], got [ %s ]",
		strings.Join(domain.GetTransactionStatusAvailable(), ", "),
		invalidStatus,
	)
	errGot := domain.NewErrTransactionStatusInvalid(invalidStatus)

	assert.Equal(t, errExpected, errGot.Error())
}

func TestErrAuthorizationResponseStatusInvalid(t *testing.T) {
	invalidStatus := "done"
	errExpected := fmt.Sprintf(
		"Invalid authorization response status, expected one of [ %s ], got [ %s ]",
		strings.Join(domain.GetAuthorizationStatusAvailable(), ", "),
		invalidStatus,
	)
	errGot := domain.NewErrAuthorizationResponseStatusInvalid(invalidStatus)

	assert.Equal(t, errExpected, errGot.Error())
}
