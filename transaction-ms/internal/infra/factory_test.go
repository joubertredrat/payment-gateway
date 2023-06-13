package infra_test

import (
	"joubertredrat/transaction-ms/internal/domain"
	"joubertredrat/transaction-ms/internal/infra"
	"joubertredrat/transaction-ms/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCreditCardTransactionResponseFromUsecase(t *testing.T) {
	creditCardTransaction := domain.CreditCardTransaction{
		ID:            1,
		TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
		CardNumber:    "513073XXXXXX7489",
		Amount: domain.Amount{
			Value: 1250,
		},
		Installments: 2,
		Description:  "usb cable",
		TransactionStatus: []domain.TransactionStatus{
			{
				ID:                      10,
				CreditCardTransactionID: 1,
				Status:                  domain.TRANSACTION_STATUS_CREATED,
				CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:01:00"),
			},
		},
		CreatedAt: pkg.TimeFromCanonical("2023-06-10 17:00:00"),
	}

	dt := "2023-06-10 14:01:00"
	responseExpected := infra.CreditCardTransactionResponse{
		TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
		CardNumber:    "513073XXXXXX7489",
		Amount: infra.AmountResponse{
			Cents: 1250,
			BRL:   "R$12,50",
		},
		Installments: 2,
		Description:  "usb cable",
		TransactionStatus: []infra.TransactionStatusResponse{
			{
				Status:    domain.TRANSACTION_STATUS_CREATED,
				CreatedAt: &dt,
			},
		},
		CreatedAt: infra.DatetimeCanonical(creditCardTransaction.CreatedAt),
		UpdatedAt: infra.DatetimeCanonical(creditCardTransaction.UpdatedAt),
	}

	responseGot := infra.CreateCreditCardTransactionResponseFromUsecase(creditCardTransaction)
	assert.Equal(t, responseExpected, responseGot)
}

func TestCreateTransactionStatusResponseFromUsecase(t *testing.T) {
	transactionStatus := domain.TransactionStatus{
		ID:        10,
		Status:    domain.TRANSACTION_STATUS_AUTHORIZED,
		CreatedAt: pkg.TimeFromCanonical("2023-06-10 17:00:00"),
	}

	dt := "2023-06-10 14:00:00"
	responseExpected := infra.TransactionStatusResponse{
		Status:    domain.TRANSACTION_STATUS_AUTHORIZED,
		CreatedAt: &dt,
	}

	responseGot := infra.CreateTransactionStatusResponseFromUsecase(transactionStatus)
	assert.Equal(t, responseExpected, responseGot)
}

func TestCreateCreditCardTransactionListResponseFromUsecase(t *testing.T) {
	list := []domain.CreditCardTransaction{
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
	}

	dt := "2023-06-10 14:00:00"
	responseExpected := []infra.CreditCardTransactionListResponse{
		{
			TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
			CardNumber:    "513073XXXXXX7489",
			Amount: infra.AmountResponse{
				Cents: 1250,
				BRL:   "R$12,50",
			},
			Installments: 2,
			Description:  "usb cable",
			CreatedAt:    &dt,
		},
	}

	responseGot := infra.CreateCreditCardTransactionListResponseFromUsecase(list)
	assert.Equal(t, responseExpected, responseGot)
}
