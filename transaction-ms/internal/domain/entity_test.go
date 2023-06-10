package domain_test

import (
	"joubertredrat/transaction-ms/internal/domain"
	"joubertredrat/transaction-ms/pkg"
	"testing"
	"time"

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

func TestNewTransactionStatus(t *testing.T) {
	tests := []struct {
		name                      string
		creditCardTransctionID    uint
		status                    string
		transactionStatusExpected domain.TransactionStatus
		errExpected               error
	}{
		{
			name:                   "test new transaction status with valid data",
			creditCardTransctionID: 1,
			status:                 domain.TRANSACTION_STATUS_AUTHORIZED,
			transactionStatusExpected: domain.TransactionStatus{
				CreditCardTransctionID: 1,
				Status:                 domain.TRANSACTION_STATUS_AUTHORIZED,
			},
			errExpected: nil,
		},
		{
			name:                      "test new transaction status with invalid status",
			creditCardTransctionID:    1,
			status:                    "done",
			transactionStatusExpected: domain.TransactionStatus{},
			errExpected:               domain.NewErrTransctionStatusInvalid("done"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			transactionStatusGot, errGot := domain.NewTransactionStatus(
				test.creditCardTransctionID,
				test.status,
			)
			assert.Equal(t, test.transactionStatusExpected, transactionStatusGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}

func TestNewAuthorizationRequest(t *testing.T) {
	tests := []struct {
		name                         string
		holderName                   string
		cardNumber                   string
		cvv                          string
		expireDate                   time.Time
		amount                       uint
		installments                 uint
		authorizationRequestExpected domain.AuthorizationRequest
		errExpected                  error
	}{
		{
			name:         "test new authorization request with valid data",
			holderName:   "John Doe",
			cardNumber:   "5130731304267489",
			cvv:          "456",
			expireDate:   *pkg.TimeFromCanonical("2025-05-01 00:00:00"),
			amount:       1250,
			installments: 2,
			authorizationRequestExpected: domain.AuthorizationRequest{
				HolderName:   "John Doe",
				CardNumber:   "5130731304267489",
				CVV:          "456",
				ExpireDate:   *pkg.TimeFromCanonical("2025-05-01 00:00:00"),
				Amount:       1250,
				Installments: 2,
			},
			errExpected: nil,
		},
		{
			name:                         "test new authorization request with invalid card number",
			holderName:                   "John Doe",
			cardNumber:                   "513073130426",
			cvv:                          "456",
			expireDate:                   *pkg.TimeFromCanonical("2025-05-01 00:00:00"),
			amount:                       1250,
			installments:                 2,
			authorizationRequestExpected: domain.AuthorizationRequest{},
			errExpected:                  domain.NewErrInvalidCreditCardNumber("513073130426"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			authorizationRequestGot, errGot := domain.NewAuthorizationRequest(
				test.holderName,
				test.cardNumber,
				test.cvv,
				test.expireDate,
				test.amount,
				test.installments,
			)
			assert.Equal(t, test.authorizationRequestExpected, authorizationRequestGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}

func TestNewAuthorizationResponse(t *testing.T) {
	tests := []struct {
		name                          string
		status                        string
		authorizationResponseExpected domain.AuthorizationResponse
		errExpected                   error
	}{
		{
			name:   "test new authorization response with valid data",
			status: domain.AUTHORIZATION_STATUS_AUTHORIZED,
			authorizationResponseExpected: domain.AuthorizationResponse{
				Status: domain.AUTHORIZATION_STATUS_AUTHORIZED,
			},
			errExpected: nil,
		},
		{
			name:                          "test new authorization response with invalid status",
			status:                        "done",
			authorizationResponseExpected: domain.AuthorizationResponse{},
			errExpected:                   domain.NewErrAuthorizationResponseStatusInvalid("done"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			authorizationResponseGot, errGot := domain.NewAuthorizationResponse(
				test.status,
			)
			assert.Equal(t, test.authorizationResponseExpected, authorizationResponseGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}

func TestAuthorizationResponseGetTransactionStatus(t *testing.T) {
	tests := []struct {
		name                  string
		authorizationResponse domain.AuthorizationResponse
		statusExpected        string
	}{
		{
			name: "test authorization response get transaction status for authorized",
			authorizationResponse: func() domain.AuthorizationResponse {
				authorizationResponse, _ := domain.NewAuthorizationResponse(
					domain.AUTHORIZATION_STATUS_AUTHORIZED,
				)
				return authorizationResponse
			}(),
			statusExpected: domain.TRANSACTION_STATUS_AUTHORIZED,
		},
		{
			name: "test authorization response get transaction status for declined",
			authorizationResponse: func() domain.AuthorizationResponse {
				authorizationResponse, _ := domain.NewAuthorizationResponse(
					domain.AUTHORIZATION_STATUS_DECLINED,
				)
				return authorizationResponse
			}(),
			statusExpected: domain.TRANSACTION_STATUS_REFUSED,
		},
		{
			name: "test authorization response get transaction status for unknown",
			authorizationResponse: domain.AuthorizationResponse{
				Status: "unknown",
			},
			statusExpected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.statusExpected, test.authorizationResponse.GetTransactionStatus())
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
