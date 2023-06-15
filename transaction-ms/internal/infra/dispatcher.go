package infra

import (
	"context"
	"fmt"
	"joubertredrat/transaction-ms/internal/application"
	"joubertredrat/transaction-ms/internal/domain"

	"github.com/redis/go-redis/v9"
)

type QueueDispatcher struct {
	logger      application.Logger
	redisClient *redis.Client
	topicName   string
}

func NewQueueDispatcher(logger application.Logger, redisClient *redis.Client, topicName string) QueueDispatcher {
	return QueueDispatcher{
		logger:      logger,
		redisClient: redisClient,
		topicName:   topicName,
	}
}

func (d QueueDispatcher) CreditCardTransactionCreated(c domain.CreditCardTransaction) error {
	d.dispatch(fmt.Sprintf("Transaction with TransactionID [ %s ] was created", c.TransactionID))
	return nil
}

func (d QueueDispatcher) CreditCardTransactionEdited(c domain.CreditCardTransaction) error {
	d.dispatch(fmt.Sprintf("Transaction with TransactionID [ %s ] was edited", c.TransactionID))
	return nil
}

func (d QueueDispatcher) CreditCardTransactionDeleted(TransactionID string) error {
	d.dispatch(fmt.Sprintf("Transaction with TransactionID [ %s ] was deleted", TransactionID))
	return nil
}

func (d QueueDispatcher) CreditCardTransactionGot(c domain.CreditCardTransaction) error {
	d.dispatch(fmt.Sprintf("Transaction with TransactionID [ %s ] was got", c.TransactionID))
	return nil
}

func (d QueueDispatcher) CreditCardTransactionListed(p domain.PaginationCriteria) error {
	d.dispatch(fmt.Sprintf(
		"Transactions was listed by pagination criteria, page [ %d ] and items per page [ %d ]",
		p.Page,
		p.ItemsPerPage,
	))
	return nil
}
func (d QueueDispatcher) dispatch(payload string) {
	if err := d.redisClient.Publish(context.Background(), d.topicName, payload).Err(); err != nil {
		d.logger.Error(err)
	}
}
