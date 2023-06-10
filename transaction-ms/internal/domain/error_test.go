package domain_test

import (
	"fmt"
	"joubertredrat/transaction-ms/internal/domain"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrInvalidCreditCardNumber(t *testing.T) {
	invalidNumber := "123456"
	errExpected := fmt.Sprintf("Invalid credit card number [ %s ]", invalidNumber)
	errGot := domain.NewErrInvalidCreditCardNumber(invalidNumber)

	assert.Equal(t, errExpected, errGot.Error())
}

func TestErrCreditCardTransctionInstallments(t *testing.T) {
	invalidInstallments := uint(16)
	errExpected := fmt.Sprintf(
		"Invalid installments, expected between [ %d ] and [ %d ], got [ %d ]",
		domain.INSTALLMENTS_MIN,
		domain.INSTALLMENTS_MAX,
		invalidInstallments,
	)
	errGot := domain.NewErrCreditCardTransctionInstallments(invalidInstallments)

	assert.Equal(t, errExpected, errGot.Error())
}

func TestErrTransctionStatusInvalid(t *testing.T) {
	invalidStatus := "done"
	errExpected := fmt.Sprintf(
		"Invalid transction status, expected one of [ %s ], got [ %s ]",
		strings.Join(domain.GetTransactionStatusAvailable(), ", "),
		invalidStatus,
	)
	errGot := domain.NewErrTransctionStatusInvalid(invalidStatus)

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
