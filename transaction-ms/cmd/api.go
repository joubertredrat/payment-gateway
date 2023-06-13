package cmd

import (
	"fmt"
	"joubertredrat/transaction-ms/internal/infra"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

func getApiCommand() *cli.Command {
	return &cli.Command{
		Name:    "api",
		Aliases: []string{},
		Usage:   "Open HTTP api to listen",
		Action: func(c *cli.Context) error {
			config, err := infra.NewConfig()
			if err != nil {
				return err
			}

			r := gin.Default()
			if err := r.SetTrustedProxies(nil); err != nil {
				return err
			}

			db, err := infra.GetDatabaseConnection(infra.GetMysqlDSN(
				config.DatabaseHost,
				config.DatabasePort,
				config.DatabaseName,
				config.DatabaseUser,
				config.DatabasePassword,
			))
			if err != nil {
				return err
			}

			logrus := infra.GetLogrusStdout()
			logger := infra.GetLoggerStdout(logrus)

			authorizationService := infra.GetAuthorizationServiceMicroservice()
			dispatcher := infra.GetQueueDispatcher()

			creditCardTransactionRepository := infra.GetCreditCardTransactionRepository(logger, db)
			transactionStatusRepository := infra.GetTransactionStatusRepository(logger, db)

			usecaseCreateCreditCardTransaction := infra.GetUsecaseCreateCreditCardTransaction(
				logger,
				creditCardTransactionRepository,
				transactionStatusRepository,
				authorizationService,
				dispatcher,
			)
			usecaseEditCreditCardTransaction := infra.GetUsecaseEditCreditCardTransaction(
				logger,
				creditCardTransactionRepository,
				transactionStatusRepository,
				dispatcher,
			)
			usecaseDeleteCreditCardTransaction := infra.GetUsecaseDeleteCreditCardTransaction(
				logger,
				creditCardTransactionRepository,
				dispatcher,
			)
			usecaseGetCreditCardTransaction := infra.GetUsecaseGetCreditCardTransaction(
				logger,
				creditCardTransactionRepository,
				transactionStatusRepository,
				dispatcher,
			)
			usecaseListCreditCardTransaction := infra.GetUsecaseListCreditCardTransaction(
				logger,
				creditCardTransactionRepository,
				dispatcher,
			)

			apiBaseController := infra.NewApiBaseController()
			creditTransactionsController := infra.NewCreditTransactionsController()

			r.NoRoute(apiBaseController.HandleNotFound)

			ra := r.Group("/api")
			infra.RegisterCustomValidator()
			{
				ra.GET("/status", apiBaseController.HandleStatus)
				rt := ra.Group("/creditcard/transactions")
				{
					rt.GET("", creditTransactionsController.HandleList(usecaseListCreditCardTransaction))
					rt.POST(
						"",
						infra.JSONBodyMiddleware(),
						creditTransactionsController.HandleCreate(usecaseCreateCreditCardTransaction),
					)
					rt.GET("/:transactionid", creditTransactionsController.HandleGet(usecaseGetCreditCardTransaction))
					rt.PATCH(
						"/:transactionid",
						infra.JSONBodyMiddleware(),
						creditTransactionsController.HandleEdit(usecaseEditCreditCardTransaction),
					)
					rt.DELETE("/:transactionid", creditTransactionsController.HandleDelete(usecaseDeleteCreditCardTransaction))
				}
			}

			return r.Run(fmt.Sprintf("%s:%s", config.ApiHost, config.ApiPort))
		},
	}
}
