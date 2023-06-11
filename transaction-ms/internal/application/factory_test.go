package application_test

import (
	"joubertredrat/transaction-ms/internal/application"
	"joubertredrat/transaction-ms/internal/domain"
	"joubertredrat/transaction-ms/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCreditCardTransactionFromInput(t *testing.T) {
	input := application.CreateCreditCardTransctionInput{
		HolderName:   "John Doe",
		CardNumber:   "5130731304267489",
		CVV:          "456",
		ExpireDate:   *pkg.TimeFromCanonical("2025-05-01 00:00:00"),
		Amount:       1250,
		Installments: 2,
		Description:  "usb cable",
	}

	creditCardTransationExpected, _ := domain.NewCreditCardTransaction(
		"",
		input.CardNumber,
		input.Amount,
		input.Installments,
		input.Description,
		[]domain.TransactionStatus{},
	)

	creditCardTransationGot, _ := application.CreateCreditCardTransactionFromInput(input)
	creditCardTransationExpected.TransactionID = creditCardTransationGot.TransactionID

	assert.Equal(t, creditCardTransationExpected, creditCardTransationGot)
}

func TestCreateAuthorizationRequestFromInput(t *testing.T) {
	input := application.CreateCreditCardTransctionInput{
		HolderName:   "John Doe",
		CardNumber:   "5130731304267489",
		CVV:          "456",
		ExpireDate:   *pkg.TimeFromCanonical("2025-05-01 00:00:00"),
		Amount:       1250,
		Installments: 2,
		Description:  "usb cable",
	}

	authorizationRequestExpected, _ := domain.NewAuthorizationRequest(
		input.HolderName,
		input.CardNumber,
		input.CVV,
		input.ExpireDate,
		input.Amount,
		input.Installments,
	)

	authorizationRequestGot, _ := application.CreateAuthorizationRequestFromInput(input)

	assert.Equal(t, authorizationRequestExpected, authorizationRequestGot)
}
