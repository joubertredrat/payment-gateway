package domain_test

import (
	"joubertredrat/transaction-ms/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAmountBRL(t *testing.T) {
	tests := []struct {
		name        string
		amountCents uint
		BRLExpected string
	}{
		{
			name:        "Test amount with R$ 0,57",
			amountCents: 57,
			BRLExpected: "R$0,57",
		},
		{
			name:        "Test amount with R$ 12,50",
			amountCents: 1250,
			BRLExpected: "R$12,50",
		},
		{
			name:        "Test amount with R$ 13.175,26",
			amountCents: 1317526,
			BRLExpected: "R$13.175,26",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			amount := domain.Amount{
				Value: test.amountCents,
			}
			BRLGot := amount.BRL()
			assert.Equal(t, test.BRLExpected, BRLGot)
		})
	}
}

func TestNewCreditCardTransaction(t *testing.T) {
	tests := []struct {
		name                         string
		transactionID                string
		cardNumber                   string
		amount                       uint
		installments                 uint
		description                  string
		transctionStatus             []domain.TransactionStatus
		creditCardTransctionExpected domain.CreditCardTransaction
		errExpected                  error
	}{
		{
			name:             "test new credit card transction with valid data",
			transactionID:    "01H2K6VSVKJA8GDC13MK28P03M",
			cardNumber:       "5130731304267489",
			amount:           1250,
			installments:     2,
			description:      "usb cable",
			transctionStatus: []domain.TransactionStatus{},
			creditCardTransctionExpected: domain.CreditCardTransaction{
				TransactionID: "01H2K6VSVKJA8GDC13MK28P03M",
				CardNumber:    "513073XXXXXX7489",
				Amount: domain.Amount{
					Value: 1250,
				},
				Installments:      2,
				Description:       "usb cable",
				TransactionStatus: []domain.TransactionStatus{},
			},
			errExpected: nil,
		},
		{
			name:                         "test new credit card transction with invalid installments",
			transactionID:                "01H2K6VSVKJA8GDC13MK28P03M",
			cardNumber:                   "5130731304267489",
			amount:                       1250,
			installments:                 15,
			description:                  "usb cable",
			transctionStatus:             []domain.TransactionStatus{},
			creditCardTransctionExpected: domain.CreditCardTransaction{},
			errExpected:                  domain.NewErrCreditCardTransctionInstallments(15),
		},
		{
			name:                         "test new credit card transction with invalid card number",
			transactionID:                "01H2K6VSVKJA8GDC13MK28P03M",
			cardNumber:                   "513073130426",
			amount:                       1250,
			installments:                 2,
			description:                  "usb cable",
			transctionStatus:             []domain.TransactionStatus{},
			creditCardTransctionExpected: domain.CreditCardTransaction{},
			errExpected:                  domain.NewErrInvalidCreditCardNumber("513073130426"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			creditCardTransctionGot, errGot := domain.NewCreditCardTransaction(
				test.transactionID,
				test.cardNumber,
				test.amount,
				test.installments,
				test.description,
				test.transctionStatus,
			)
			assert.Equal(t, test.creditCardTransctionExpected, creditCardTransctionGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}

func TestIsValidInstallments(t *testing.T) {
	tests := []struct {
		name           string
		installment    uint
		returnExpected bool
	}{
		{
			name:           "Test is valid installments with valid installment",
			installment:    4,
			returnExpected: true,
		},
		{
			name:           "Test is valid installments with less than minimum expected",
			installment:    0,
			returnExpected: false,
		},
		{
			name:           "Test is valid installments with more than maximum expected",
			installment:    15,
			returnExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			returnGot := domain.IsValidInstallments(test.installment)
			assert.Equal(t, test.returnExpected, returnGot)
		})
	}
}

func TestIsValidTransactionStatus(t *testing.T) {
	tests := []struct {
		name           string
		status         string
		returnExpected bool
	}{
		{
			name:           "Test is valid transaction status with valid status",
			status:         domain.TRANSACTION_STATUS_CREATED,
			returnExpected: true,
		},
		{
			name:           "Test is valid transaction status with invalid status",
			status:         "done",
			returnExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			returnGot := domain.IsValidTransactionStatus(test.status)
			assert.Equal(t, test.returnExpected, returnGot)
		})
	}
}

func TestIsValidAuthorizationStatus(t *testing.T) {
	tests := []struct {
		name           string
		status         string
		returnExpected bool
	}{
		{
			name:           "Test is valid authorization status with valid status",
			status:         domain.AUTHORIZATION_STATUS_AUTHORIZED,
			returnExpected: true,
		},
		{
			name:           "Test is valid authorization status with invalid status",
			status:         "done",
			returnExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			returnGot := domain.IsValidAuthorizationStatus(test.status)
			assert.Equal(t, test.returnExpected, returnGot)
		})
	}
}
