package application_test

import (
	"joubertredrat/transaction-ms/internal/application"
	"joubertredrat/transaction-ms/internal/domain"
	"joubertredrat/transaction-ms/pkg"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecaseCreateCreditCardTransaction(t *testing.T) {
	tests := []struct {
		name                                      string
		loggerDependency                          func(ctrl *gomock.Controller) application.Logger
		creditCardTransactionRepositoryDependency func(ctrl *gomock.Controller) domain.CreditCardTransactionRepository
		transactionStatusRepositoryDependency     func(ctrl *gomock.Controller) domain.TransactionStatusRepository
		authorizationServiceDependency            func(ctrl *gomock.Controller) domain.AuthorizationService
		dispatcher                                func(ctrl *gomock.Controller) application.Dispatcher
		input                                     application.CreateCreditCardTransactionInput
		creditCardTransactionExpected             domain.CreditCardTransaction
		errExpected                               error
	}{
		{
			name: "test usecase create credit card Transaction with valid data",
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
						ID:                      10,
						CreditCardTransactionID: 1,
						Status:                  domain.TRANSACTION_STATUS_CREATED,
						CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:01:00"),
					}, nil).
					Times(1)

				repository.
					EXPECT().
					Create(gomock.AssignableToTypeOf(domain.TransactionStatus{})).
					Return(domain.TransactionStatus{
						ID:                      11,
						CreditCardTransactionID: 1,
						Status:                  domain.TRANSACTION_STATUS_AUTHORIZED,
						CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:02:00"),
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
			dispatcher: func(ctrl *gomock.Controller) application.Dispatcher {
				dispatcher := pkg.NewMockDispatcher(ctrl)
				dispatcher.
					EXPECT().
					CreditCardTransactionCreated(gomock.AssignableToTypeOf(domain.CreditCardTransaction{})).
					Return(nil).
					Times(1)

				return dispatcher
			},
			input: application.CreateCreditCardTransactionInput{
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
						ID:                      10,
						CreditCardTransactionID: 1,
						Status:                  domain.TRANSACTION_STATUS_CREATED,
						CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:01:00"),
					},
					{
						ID:                      11,
						CreditCardTransactionID: 1,
						Status:                  domain.TRANSACTION_STATUS_AUTHORIZED,
						CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:02:00"),
					},
				},
				CreatedAt: pkg.TimeFromCanonical("2023-06-10 17:00:00"),
			},
			errExpected: nil,
		},
		{
			name: "test usecase create credit card Transaction with invalid credit card number",
			loggerDependency: func(ctrl *gomock.Controller) application.Logger {
				logger := pkg.NewMockLogger(ctrl)
				logger.EXPECT().Error(gomock.Any()).Times(1)
				return logger
			},
			creditCardTransactionRepositoryDependency: func(ctrl *gomock.Controller) domain.CreditCardTransactionRepository {
				return pkg.NewMockCreditCardTransactionRepository(ctrl)
			},
			transactionStatusRepositoryDependency: func(ctrl *gomock.Controller) domain.TransactionStatusRepository {
				return pkg.NewMockTransactionStatusRepository(ctrl)
			},
			authorizationServiceDependency: func(ctrl *gomock.Controller) domain.AuthorizationService {
				return pkg.NewMockAuthorizationService(ctrl)
			},
			dispatcher: func(ctrl *gomock.Controller) application.Dispatcher {
				return pkg.NewMockDispatcher(ctrl)
			},
			input: application.CreateCreditCardTransactionInput{
				HolderName:   "John Doe",
				CardNumber:   "513073130426",
				CVV:          "456",
				ExpireDate:   *pkg.TimeFromCanonical("2025-05-01 00:00:00"),
				Amount:       1250,
				Installments: 2,
				Description:  "usb cable",
			},
			creditCardTransactionExpected: domain.CreditCardTransaction{},
			errExpected:                   domain.NewErrInvalidCreditCardNumber("513073130426"),
		},
		{
			name: "test usecase create credit card Transaction with houston error from credit card transaction repository",
			loggerDependency: func(ctrl *gomock.Controller) application.Logger {
				logger := pkg.NewMockLogger(ctrl)
				logger.EXPECT().Error(gomock.Any()).Times(1)
				return logger
			},
			creditCardTransactionRepositoryDependency: func(ctrl *gomock.Controller) domain.CreditCardTransactionRepository {
				repository := pkg.NewMockCreditCardTransactionRepository(ctrl)
				repository.
					EXPECT().
					Create(gomock.AssignableToTypeOf(domain.CreditCardTransaction{})).
					Return(domain.CreditCardTransaction{}, domain.ErrCreditCardTransactionRepositoryHouston).
					Times(1)

				return repository
			},
			transactionStatusRepositoryDependency: func(ctrl *gomock.Controller) domain.TransactionStatusRepository {
				return pkg.NewMockTransactionStatusRepository(ctrl)
			},
			authorizationServiceDependency: func(ctrl *gomock.Controller) domain.AuthorizationService {
				return pkg.NewMockAuthorizationService(ctrl)
			},
			dispatcher: func(ctrl *gomock.Controller) application.Dispatcher {
				return pkg.NewMockDispatcher(ctrl)
			},
			input: application.CreateCreditCardTransactionInput{
				HolderName:   "John Doe",
				CardNumber:   "5130731304267489",
				CVV:          "456",
				ExpireDate:   *pkg.TimeFromCanonical("2025-05-01 00:00:00"),
				Amount:       1250,
				Installments: 2,
				Description:  "usb cable",
			},
			creditCardTransactionExpected: domain.CreditCardTransaction{},
			errExpected:                   application.ErrUsecaseCreateCreditCardTransactionHouston,
		},
		{
			name: "test usecase create credit card Transaction with houston error from first create on transaction status repository",
			loggerDependency: func(ctrl *gomock.Controller) application.Logger {
				logger := pkg.NewMockLogger(ctrl)
				logger.EXPECT().Error(gomock.Any()).Times(1)
				return logger
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
					Return(domain.TransactionStatus{}, domain.ErrTransactionStatusRepositoryHouston).
					Times(1)

				return repository
			},
			authorizationServiceDependency: func(ctrl *gomock.Controller) domain.AuthorizationService {
				return pkg.NewMockAuthorizationService(ctrl)
			},
			dispatcher: func(ctrl *gomock.Controller) application.Dispatcher {
				return pkg.NewMockDispatcher(ctrl)
			},
			input: application.CreateCreditCardTransactionInput{
				HolderName:   "John Doe",
				CardNumber:   "5130731304267489",
				CVV:          "456",
				ExpireDate:   *pkg.TimeFromCanonical("2025-05-01 00:00:00"),
				Amount:       1250,
				Installments: 2,
				Description:  "usb cable",
			},
			creditCardTransactionExpected: domain.CreditCardTransaction{},
			errExpected:                   application.ErrUsecaseCreateCreditCardTransactionHouston,
		},
		{
			name: "test usecase create credit card Transaction with houston error from authorization service",
			loggerDependency: func(ctrl *gomock.Controller) application.Logger {
				logger := pkg.NewMockLogger(ctrl)
				logger.EXPECT().Error(gomock.Any()).Times(1)
				return logger
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
						ID:                      10,
						CreditCardTransactionID: 1,
						Status:                  domain.TRANSACTION_STATUS_CREATED,
						CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:01:00"),
					}, nil).
					Times(1)

				repository.
					EXPECT().
					Create(gomock.AssignableToTypeOf(domain.TransactionStatus{})).
					Return(domain.TransactionStatus{
						ID:                      11,
						CreditCardTransactionID: 1,
						Status:                  domain.TRANSACTION_STATUS_ERROR_AUTHORIZATION,
						CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:02:00"),
					}, nil).
					Times(1)

				return repository
			},
			authorizationServiceDependency: func(ctrl *gomock.Controller) domain.AuthorizationService {
				service := pkg.NewMockAuthorizationService(ctrl)
				service.
					EXPECT().
					Handle(gomock.AssignableToTypeOf(domain.AuthorizationRequest{})).
					Return(domain.AuthorizationResponse{}, domain.ErrAuthorizationServiceHouston).
					Times(1)

				return service
			},
			dispatcher: func(ctrl *gomock.Controller) application.Dispatcher {
				return pkg.NewMockDispatcher(ctrl)
			},
			input: application.CreateCreditCardTransactionInput{
				HolderName:   "John Doe",
				CardNumber:   "5130731304267489",
				CVV:          "456",
				ExpireDate:   *pkg.TimeFromCanonical("2025-05-01 00:00:00"),
				Amount:       1250,
				Installments: 2,
				Description:  "usb cable",
			},
			creditCardTransactionExpected: domain.CreditCardTransaction{},
			errExpected:                   application.ErrUsecaseCreateCreditCardTransactionHouston,
		},
		{
			name: "test usecase create credit card Transaction with houston error from create error status on transaction status repository",
			loggerDependency: func(ctrl *gomock.Controller) application.Logger {
				logger := pkg.NewMockLogger(ctrl)
				logger.EXPECT().Error(gomock.Any()).Times(2)
				return logger
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
						ID:                      10,
						CreditCardTransactionID: 1,
						Status:                  domain.TRANSACTION_STATUS_CREATED,
						CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:01:00"),
					}, nil).
					Times(1)

				repository.
					EXPECT().
					Create(gomock.AssignableToTypeOf(domain.TransactionStatus{})).
					Return(domain.TransactionStatus{}, domain.ErrTransactionStatusRepositoryHouston).
					Times(1)

				return repository
			},
			authorizationServiceDependency: func(ctrl *gomock.Controller) domain.AuthorizationService {
				service := pkg.NewMockAuthorizationService(ctrl)
				service.
					EXPECT().
					Handle(gomock.AssignableToTypeOf(domain.AuthorizationRequest{})).
					Return(domain.AuthorizationResponse{}, domain.ErrAuthorizationServiceHouston).
					Times(1)

				return service
			},
			dispatcher: func(ctrl *gomock.Controller) application.Dispatcher {
				return pkg.NewMockDispatcher(ctrl)
			},
			input: application.CreateCreditCardTransactionInput{
				HolderName:   "John Doe",
				CardNumber:   "5130731304267489",
				CVV:          "456",
				ExpireDate:   *pkg.TimeFromCanonical("2025-05-01 00:00:00"),
				Amount:       1250,
				Installments: 2,
				Description:  "usb cable",
			},
			creditCardTransactionExpected: domain.CreditCardTransaction{},
			errExpected:                   application.ErrUsecaseCreateCreditCardTransactionHouston,
		},
		{
			name: "test usecase create credit card Transaction with houston error from second create on transaction status repository",
			loggerDependency: func(ctrl *gomock.Controller) application.Logger {
				logger := pkg.NewMockLogger(ctrl)
				logger.EXPECT().Error(gomock.Any()).Times(1)
				return logger
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
						ID:                      10,
						CreditCardTransactionID: 1,
						Status:                  domain.TRANSACTION_STATUS_CREATED,
						CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:01:00"),
					}, nil).
					Times(1)

				repository.
					EXPECT().
					Create(gomock.AssignableToTypeOf(domain.TransactionStatus{})).
					Return(domain.TransactionStatus{}, domain.ErrTransactionStatusRepositoryHouston).
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
			dispatcher: func(ctrl *gomock.Controller) application.Dispatcher {
				return pkg.NewMockDispatcher(ctrl)
			},
			input: application.CreateCreditCardTransactionInput{
				HolderName:   "John Doe",
				CardNumber:   "5130731304267489",
				CVV:          "456",
				ExpireDate:   *pkg.TimeFromCanonical("2025-05-01 00:00:00"),
				Amount:       1250,
				Installments: 2,
				Description:  "usb cable",
			},
			creditCardTransactionExpected: domain.CreditCardTransaction{},
			errExpected:                   application.ErrUsecaseCreateCreditCardTransactionHouston,
		},
		{
			name: "test usecase create credit card Transaction with houston error on dispatch created transaction",
			loggerDependency: func(ctrl *gomock.Controller) application.Logger {
				logger := pkg.NewMockLogger(ctrl)
				logger.EXPECT().Error(gomock.Any()).Times(1)
				return logger
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
						ID:                      10,
						CreditCardTransactionID: 1,
						Status:                  domain.TRANSACTION_STATUS_CREATED,
						CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:01:00"),
					}, nil).
					Times(1)

				repository.
					EXPECT().
					Create(gomock.AssignableToTypeOf(domain.TransactionStatus{})).
					Return(domain.TransactionStatus{
						ID:                      11,
						CreditCardTransactionID: 1,
						Status:                  domain.TRANSACTION_STATUS_AUTHORIZED,
						CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:02:00"),
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
			dispatcher: func(ctrl *gomock.Controller) application.Dispatcher {
				dispatcher := pkg.NewMockDispatcher(ctrl)
				dispatcher.
					EXPECT().
					CreditCardTransactionCreated(gomock.AssignableToTypeOf(domain.CreditCardTransaction{})).
					Return(application.ErrDispatcherHouston).
					Times(1)

				return dispatcher
			},
			input: application.CreateCreditCardTransactionInput{
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
						ID:                      10,
						CreditCardTransactionID: 1,
						Status:                  domain.TRANSACTION_STATUS_CREATED,
						CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:01:00"),
					},
					{
						ID:                      11,
						CreditCardTransactionID: 1,
						Status:                  domain.TRANSACTION_STATUS_AUTHORIZED,
						CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:02:00"),
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

			usecase := application.NewUsecaseCreateCreditCardTransaction(
				test.loggerDependency(ctrl),
				test.creditCardTransactionRepositoryDependency(ctrl),
				test.transactionStatusRepositoryDependency(ctrl),
				test.authorizationServiceDependency(ctrl),
				test.dispatcher(ctrl),
			)

			creditCardTransactionGot, errGot := usecase.Execute(test.input)

			assert.Equal(t, test.creditCardTransactionExpected, creditCardTransactionGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}

func TestUsecaseEditCreditCardTransaction(t *testing.T) {
	tests := []struct {
		name                                      string
		loggerDependency                          func(ctrl *gomock.Controller) application.Logger
		creditCardTransactionRepositoryDependency func(ctrl *gomock.Controller) domain.CreditCardTransactionRepository
		transactionStatusRepositoryDependency     func(ctrl *gomock.Controller) domain.TransactionStatusRepository
		dispatcher                                func(ctrl *gomock.Controller) application.Dispatcher
		input                                     application.EditCreditCardTransactionInput
		creditCardTransactionExpected             domain.CreditCardTransaction
		errExpected                               error
	}{
		{
			name: "test usecase edit credit card Transaction with valid data",
			loggerDependency: func(ctrl *gomock.Controller) application.Logger {
				return pkg.NewMockLogger(ctrl)
			},
			creditCardTransactionRepositoryDependency: func(ctrl *gomock.Controller) domain.CreditCardTransactionRepository {
				repository := pkg.NewMockCreditCardTransactionRepository(ctrl)

				repository.
					EXPECT().
					GetByTransactionID(gomock.Eq("01H2KDJMHCTVTN0YDY10S5SNWB")).
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
				repository.
					EXPECT().
					Update(gomock.AssignableToTypeOf(domain.CreditCardTransaction{})).
					Return(domain.CreditCardTransaction{
						ID:            1,
						TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
						CardNumber:    "513073XXXXXX7489",
						Amount: domain.Amount{
							Value: 1250,
						},
						Installments:      2,
						Description:       "super usb cable",
						TransactionStatus: []domain.TransactionStatus{},
						CreatedAt:         pkg.TimeFromCanonical("2023-06-10 17:00:00"),
						UpdatedAt:         pkg.TimeFromCanonical("2023-06-10 21:30:00"),
					}, nil).
					Times(1)

				return repository
			},
			transactionStatusRepositoryDependency: func(ctrl *gomock.Controller) domain.TransactionStatusRepository {
				repository := pkg.NewMockTransactionStatusRepository(ctrl)
				repository.
					EXPECT().
					GetByCreditCardTransactionID(gomock.Eq("01H2KDJMHCTVTN0YDY10S5SNWB")).
					Return([]domain.TransactionStatus{
						{
							ID:                      10,
							CreditCardTransactionID: 1,
							Status:                  domain.TRANSACTION_STATUS_CREATED,
							CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:01:00"),
						},
						{
							ID:                      11,
							CreditCardTransactionID: 1,
							Status:                  domain.TRANSACTION_STATUS_AUTHORIZED,
							CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:02:00"),
						},
					}, nil).
					Times(1)

				return repository
			},
			dispatcher: func(ctrl *gomock.Controller) application.Dispatcher {
				dispatcher := pkg.NewMockDispatcher(ctrl)
				dispatcher.
					EXPECT().
					CreditCardTransactionEdited(gomock.AssignableToTypeOf(domain.CreditCardTransaction{})).
					Return(nil).
					Times(1)

				return dispatcher
			},
			input: application.EditCreditCardTransactionInput{
				TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
				Description:   "super usb cable",
			},
			creditCardTransactionExpected: domain.CreditCardTransaction{
				ID:            1,
				TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
				CardNumber:    "513073XXXXXX7489",
				Amount: domain.Amount{
					Value: 1250,
				},
				Installments: 2,
				Description:  "super usb cable",
				TransactionStatus: []domain.TransactionStatus{
					{
						ID:                      10,
						CreditCardTransactionID: 1,
						Status:                  domain.TRANSACTION_STATUS_CREATED,
						CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:01:00"),
					},
					{
						ID:                      11,
						CreditCardTransactionID: 1,
						Status:                  domain.TRANSACTION_STATUS_AUTHORIZED,
						CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:02:00"),
					},
				},
				CreatedAt: pkg.TimeFromCanonical("2023-06-10 17:00:00"),
				UpdatedAt: pkg.TimeFromCanonical("2023-06-10 21:30:00"),
			},
			errExpected: nil,
		},
		{
			name: "test usecase edit credit card Transaction with not found Transaction",
			loggerDependency: func(ctrl *gomock.Controller) application.Logger {
				logger := pkg.NewMockLogger(ctrl)
				logger.EXPECT().Error(gomock.Any()).Times(1)
				return logger
			},
			creditCardTransactionRepositoryDependency: func(ctrl *gomock.Controller) domain.CreditCardTransactionRepository {
				repository := pkg.NewMockCreditCardTransactionRepository(ctrl)
				repository.
					EXPECT().
					GetByTransactionID(gomock.Eq("01H2KDJMHCTVTN0YDY10S5SNWB")).
					Return(domain.CreditCardTransaction{}, domain.NewErrCreditCardTransactionNotFound("TransactionID", "01H2KDJMHCTVTN0YDY10S5SNWB")).
					Times(1)

				return repository
			},
			transactionStatusRepositoryDependency: func(ctrl *gomock.Controller) domain.TransactionStatusRepository {
				return pkg.NewMockTransactionStatusRepository(ctrl)
			},
			dispatcher: func(ctrl *gomock.Controller) application.Dispatcher {
				return pkg.NewMockDispatcher(ctrl)
			},
			input: application.EditCreditCardTransactionInput{
				TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
				Description:   "super usb cable",
			},
			creditCardTransactionExpected: domain.CreditCardTransaction{},
			errExpected:                   domain.NewErrCreditCardTransactionNotFound("TransactionID", "01H2KDJMHCTVTN0YDY10S5SNWB"),
		},
		{
			name: "test usecase edit credit card Transaction with houston error from credit card transaction repository on get transaction",
			loggerDependency: func(ctrl *gomock.Controller) application.Logger {
				logger := pkg.NewMockLogger(ctrl)
				logger.EXPECT().Error(gomock.Any()).Times(1)
				return logger
			},
			creditCardTransactionRepositoryDependency: func(ctrl *gomock.Controller) domain.CreditCardTransactionRepository {
				repository := pkg.NewMockCreditCardTransactionRepository(ctrl)
				repository.
					EXPECT().
					GetByTransactionID(gomock.Eq("01H2KDJMHCTVTN0YDY10S5SNWB")).
					Return(domain.CreditCardTransaction{}, domain.ErrCreditCardTransactionRepositoryHouston).
					Times(1)

				return repository
			},
			transactionStatusRepositoryDependency: func(ctrl *gomock.Controller) domain.TransactionStatusRepository {
				return pkg.NewMockTransactionStatusRepository(ctrl)
			},
			dispatcher: func(ctrl *gomock.Controller) application.Dispatcher {
				return pkg.NewMockDispatcher(ctrl)
			},
			input: application.EditCreditCardTransactionInput{
				TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
				Description:   "super usb cable",
			},
			creditCardTransactionExpected: domain.CreditCardTransaction{},
			errExpected:                   application.ErrUsecaseEditCreditCardTransactionHouston,
		},
		{
			name: "test usecase edit credit card Transaction with houston error from credit card transaction repository on update transaction",
			loggerDependency: func(ctrl *gomock.Controller) application.Logger {
				logger := pkg.NewMockLogger(ctrl)
				logger.EXPECT().Error(gomock.Any()).Times(1)
				return logger
			},
			creditCardTransactionRepositoryDependency: func(ctrl *gomock.Controller) domain.CreditCardTransactionRepository {
				repository := pkg.NewMockCreditCardTransactionRepository(ctrl)

				repository.
					EXPECT().
					GetByTransactionID(gomock.Eq("01H2KDJMHCTVTN0YDY10S5SNWB")).
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
				repository.
					EXPECT().
					Update(gomock.AssignableToTypeOf(domain.CreditCardTransaction{})).
					Return(domain.CreditCardTransaction{}, domain.ErrCreditCardTransactionRepositoryHouston).
					Times(1)

				return repository
			},
			transactionStatusRepositoryDependency: func(ctrl *gomock.Controller) domain.TransactionStatusRepository {
				return pkg.NewMockTransactionStatusRepository(ctrl)
			},
			dispatcher: func(ctrl *gomock.Controller) application.Dispatcher {
				return pkg.NewMockDispatcher(ctrl)
			},
			input: application.EditCreditCardTransactionInput{
				TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
				Description:   "super usb cable",
			},
			creditCardTransactionExpected: domain.CreditCardTransaction{},
			errExpected:                   application.ErrUsecaseEditCreditCardTransactionHouston,
		},
		{
			name: "test usecase edit credit card Transaction with houston error from transaction status repository on get status",
			loggerDependency: func(ctrl *gomock.Controller) application.Logger {
				logger := pkg.NewMockLogger(ctrl)
				logger.EXPECT().Error(gomock.Any()).Times(1)
				return logger
			},
			creditCardTransactionRepositoryDependency: func(ctrl *gomock.Controller) domain.CreditCardTransactionRepository {
				repository := pkg.NewMockCreditCardTransactionRepository(ctrl)

				repository.
					EXPECT().
					GetByTransactionID(gomock.Eq("01H2KDJMHCTVTN0YDY10S5SNWB")).
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
				repository.
					EXPECT().
					Update(gomock.AssignableToTypeOf(domain.CreditCardTransaction{})).
					Return(domain.CreditCardTransaction{
						ID:            1,
						TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
						CardNumber:    "513073XXXXXX7489",
						Amount: domain.Amount{
							Value: 1250,
						},
						Installments:      2,
						Description:       "super usb cable",
						TransactionStatus: []domain.TransactionStatus{},
						CreatedAt:         pkg.TimeFromCanonical("2023-06-10 17:00:00"),
						UpdatedAt:         pkg.TimeFromCanonical("2023-06-10 21:30:00"),
					}, nil).
					Times(1)

				return repository
			},
			transactionStatusRepositoryDependency: func(ctrl *gomock.Controller) domain.TransactionStatusRepository {
				repository := pkg.NewMockTransactionStatusRepository(ctrl)
				repository.
					EXPECT().
					GetByCreditCardTransactionID(gomock.Eq("01H2KDJMHCTVTN0YDY10S5SNWB")).
					Return([]domain.TransactionStatus{}, domain.ErrTransactionStatusRepositoryHouston).
					Times(1)

				return repository
			},
			dispatcher: func(ctrl *gomock.Controller) application.Dispatcher {
				return pkg.NewMockDispatcher(ctrl)
			},
			input: application.EditCreditCardTransactionInput{
				TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
				Description:   "super usb cable",
			},
			creditCardTransactionExpected: domain.CreditCardTransaction{},
			errExpected:                   application.ErrUsecaseEditCreditCardTransactionHouston,
		},
		{
			name: "test usecase edit credit card Transaction with houston error on dispatch created transaction",
			loggerDependency: func(ctrl *gomock.Controller) application.Logger {
				logger := pkg.NewMockLogger(ctrl)
				logger.EXPECT().Error(gomock.Any()).Times(1)
				return logger
			},
			creditCardTransactionRepositoryDependency: func(ctrl *gomock.Controller) domain.CreditCardTransactionRepository {
				repository := pkg.NewMockCreditCardTransactionRepository(ctrl)

				repository.
					EXPECT().
					GetByTransactionID(gomock.Eq("01H2KDJMHCTVTN0YDY10S5SNWB")).
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
				repository.
					EXPECT().
					Update(gomock.AssignableToTypeOf(domain.CreditCardTransaction{})).
					Return(domain.CreditCardTransaction{
						ID:            1,
						TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
						CardNumber:    "513073XXXXXX7489",
						Amount: domain.Amount{
							Value: 1250,
						},
						Installments:      2,
						Description:       "super usb cable",
						TransactionStatus: []domain.TransactionStatus{},
						CreatedAt:         pkg.TimeFromCanonical("2023-06-10 17:00:00"),
						UpdatedAt:         pkg.TimeFromCanonical("2023-06-10 21:30:00"),
					}, nil).
					Times(1)

				return repository
			},
			transactionStatusRepositoryDependency: func(ctrl *gomock.Controller) domain.TransactionStatusRepository {
				repository := pkg.NewMockTransactionStatusRepository(ctrl)
				repository.
					EXPECT().
					GetByCreditCardTransactionID(gomock.Eq("01H2KDJMHCTVTN0YDY10S5SNWB")).
					Return([]domain.TransactionStatus{
						{
							ID:                      10,
							CreditCardTransactionID: 1,
							Status:                  domain.TRANSACTION_STATUS_CREATED,
							CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:01:00"),
						},
						{
							ID:                      11,
							CreditCardTransactionID: 1,
							Status:                  domain.TRANSACTION_STATUS_AUTHORIZED,
							CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:02:00"),
						},
					}, nil).
					Times(1)

				return repository
			},
			dispatcher: func(ctrl *gomock.Controller) application.Dispatcher {
				dispatcher := pkg.NewMockDispatcher(ctrl)
				dispatcher.
					EXPECT().
					CreditCardTransactionEdited(gomock.AssignableToTypeOf(domain.CreditCardTransaction{})).
					Return(application.ErrDispatcherHouston).
					Times(1)

				return dispatcher
			},
			input: application.EditCreditCardTransactionInput{
				TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
				Description:   "super usb cable",
			},
			creditCardTransactionExpected: domain.CreditCardTransaction{
				ID:            1,
				TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
				CardNumber:    "513073XXXXXX7489",
				Amount: domain.Amount{
					Value: 1250,
				},
				Installments: 2,
				Description:  "super usb cable",
				TransactionStatus: []domain.TransactionStatus{
					{
						ID:                      10,
						CreditCardTransactionID: 1,
						Status:                  domain.TRANSACTION_STATUS_CREATED,
						CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:01:00"),
					},
					{
						ID:                      11,
						CreditCardTransactionID: 1,
						Status:                  domain.TRANSACTION_STATUS_AUTHORIZED,
						CreatedAt:               pkg.TimeFromCanonical("2023-06-10 17:02:00"),
					},
				},
				CreatedAt: pkg.TimeFromCanonical("2023-06-10 17:00:00"),
				UpdatedAt: pkg.TimeFromCanonical("2023-06-10 21:30:00"),
			},
			errExpected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			usecase := application.NewUsecaseEditCreditCardTransaction(
				test.loggerDependency(ctrl),
				test.creditCardTransactionRepositoryDependency(ctrl),
				test.transactionStatusRepositoryDependency(ctrl),
				test.dispatcher(ctrl),
			)

			creditCardTransactionGot, errGot := usecase.Execute(test.input)

			assert.Equal(t, test.creditCardTransactionExpected, creditCardTransactionGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}
