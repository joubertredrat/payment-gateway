package cmd

import (
	"context"
	"errors"
	"fmt"
	"joubertredrat/transaction-ms/internal/infra"
	"joubertredrat/transaction-ms/internal/infra/authorization"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
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

			grpcAuthorization, err := grpc.Dial(
				fmt.Sprintf("%s:%s", config.AuthorizationMsHost, config.AuthorizationMsPort),
				grpc.WithInsecure(),
			)
			if err != nil {
				return err
			}

			redisClient := redis.NewClient(&redis.Options{
				Addr: fmt.Sprintf("%s:%s", config.RedisQueueHost, config.RedisQueuePort),
			})

			logrus := infra.GetLogrusStdout()
			logger := infra.GetLoggerStdout(logrus)

			authorizationService := infra.GetAuthorizationServiceMicroservice(
				logger,
				authorization.NewAuthorizationClient(grpcAuthorization),
			)
			dispatcher := infra.GetQueueDispatcher(logger, redisClient, config.RedisQueueTransactionTopicName)

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

			server := &http.Server{
				Addr:    fmt.Sprintf("%s:%s", config.ApiHost, config.ApiPort),
				Handler: r,
			}

			go func() {
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Error(err)
				}
			}()

			done := make(chan os.Signal)
			signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
			<-done

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				return err
			}

			if err := grpcAuthorization.Close(); err != nil {
				logger.Error(err)
			}

			if err := redisClient.Close(); err != nil {
				logger.Error(err)
			}

			if err := db.Close(); err != nil {
				logger.Error(err)
			}

			logger.Info("Finishing api listener")
			return nil
		},
	}
}
