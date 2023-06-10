package application_test

import (
	"joubertredrat/transaction-ms/internal/application"
	"joubertredrat/transaction-ms/internal/domain"
	"joubertredrat/transaction-ms/pkg"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecaseCreateCreditCardTransction(t *testing.T) {
	tests := []struct {
		name                                      string
		loggerDependency                          func(ctrl *gomock.Controller) application.Logger
		creditCardTransactionRepositoryDependency func(ctrl *gomock.Controller) domain.CreditCardTransactionRepository
		transactionStatusRepositoryDependency     func(ctrl *gomock.Controller) domain.TransactionStatusRepository
		authorizationServiceDependency            func(ctrl *gomock.Controller) domain.AuthorizationService
		input                                     application.CreateCreditCardTransctionInput
		creditCardTransactionExpected             domain.CreditCardTransaction
		errExpected                               error
	}{
		{
			name: "Test usecase create credit card transction with valid data",
			loggerDependency: func(ctrl *gomock.Controller) application.Logger {
				return pkg.NewMockLogger(ctrl)
			},
			creditCardTransactionRepositoryDependency: func(ctrl *gomock.Controller) domain.CreditCardTransactionRepository {
				repository := pkg.NewMockCreditCardTransactionRepository(ctrl)
				repository.
					EXPECT().
					Create(gomock.AssignableToTypeOf(domain.CreditCardTransaction{})).
					Return(domain.CreditCardTransaction{
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
					}, nil).
					Times(1)

				return repository
			},
			transactionStatusRepositoryDependency: func(ctrl *gomock.Controller) domain.TransactionStatusRepository {
				repository := pkg.NewMockTransactionStatusRepository(ctrl)

				repository.
					EXPECT().
					Create(gomock.AssignableToTypeOf(domain.TransactionStatus{})).
					Return(domain.TransactionStatus{
						ID:                     10,
						CreditCardTransctionID: 1,
						Status:                 domain.TRANSACTION_STATUS_CREATED,
						CreatedAt:              pkg.TimeFromCanonical("2023-06-10 17:01:00"),
					}, nil).
					Times(1)

				repository.
					EXPECT().
					Create(gomock.AssignableToTypeOf(domain.TransactionStatus{})).
					Return(domain.TransactionStatus{
						ID:                     11,
						CreditCardTransctionID: 1,
						Status:                 domain.TRANSACTION_STATUS_AUTHORIZED,
						CreatedAt:              pkg.TimeFromCanonical("2023-06-10 17:02:00"),
					}, nil).
					Times(1)

				return repository
			},
			authorizationServiceDependency: func(ctrl *gomock.Controller) domain.AuthorizationService {
				service := pkg.NewMockAuthorizationService(ctrl)
				service.
					EXPECT().
					Handle(gomock.AssignableToTypeOf(domain.AuthorizationRequest{})).
					Return(domain.AuthorizationResponse{
						Status: domain.AUTHORIZATION_STATUS_AUTHORIZED,
					}, nil).
					Times(1)

				return service
			},
			input: application.CreateCreditCardTransctionInput{
				HolderName:   "John Doe",
				CardNumber:   "5130731304267489",
				CVV:          "456",
				ExpireDate:   *pkg.TimeFromCanonical("2025-05-01 00:00:00"),
				Amount:       1250,
				Installments: 2,
				Description:  "usb cable",
			},
			creditCardTransactionExpected: domain.CreditCardTransaction{
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
						ID:                     10,
						CreditCardTransctionID: 1,
						Status:                 domain.TRANSACTION_STATUS_CREATED,
						CreatedAt:              pkg.TimeFromCanonical("2023-06-10 17:01:00"),
					},
					{
						ID:                     11,
						CreditCardTransctionID: 1,
						Status:                 domain.TRANSACTION_STATUS_AUTHORIZED,
						CreatedAt:              pkg.TimeFromCanonical("2023-06-10 17:02:00"),
					},
				},
				CreatedAt: pkg.TimeFromCanonical("2023-06-10 17:00:00"),
			},
			errExpected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			usecase := application.NewUsecaseCreateCreditCardTransction(
				test.loggerDependency(ctrl),
				test.creditCardTransactionRepositoryDependency(ctrl),
				test.transactionStatusRepositoryDependency(ctrl),
				test.authorizationServiceDependency(ctrl),
			)

			creditCardTransactionGot, errGot := usecase.Execute(test.input)

			assert.Equal(t, test.creditCardTransactionExpected, creditCardTransactionGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}
