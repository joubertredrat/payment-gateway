package infra_test

import (
	"context"
	"joubertredrat/transaction-ms/internal/application"
	"joubertredrat/transaction-ms/internal/domain"
	"joubertredrat/transaction-ms/internal/infra"
	"joubertredrat/transaction-ms/pkg"
	"testing"

	redisClient "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/redis"
	"go.uber.org/mock/gomock"
)

const (
	CONTAINER_IMAGE = "redis:6"
	PORT            = ""
	TOPIC_NAME      = "transactions_test"
)

type RedisIntegrationTestSuite struct {
	suite.Suite
	ctx           context.Context
	testContainer TestContainer
	redisClient   *redisClient.Client
	logger        application.Logger
	ctrl          *gomock.Controller
}

func TestRedisIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(RedisIntegrationTestSuite))
}

func (s *RedisIntegrationTestSuite) SetupSuite() {
	s.ctx = context.Background()
	s.testContainer = setupRedis(s.ctx, s.T())
	opts, err := redisClient.ParseURL(s.testContainer.URI)
	require.Nil(s.T(), err)
	s.redisClient = redisClient.NewClient(opts)
	s.ctrl = gomock.NewController(s.T())
	s.logger = pkg.NewMockLogger(s.ctrl)
}

func (s *RedisIntegrationTestSuite) TearDownSuite() {
	s.ctrl.Finish()
	if err := s.testContainer.Terminate(s.ctx); err != nil {
		s.T().Errorf("Error on terminate test container: %s", err)
	}
}

func (s *RedisIntegrationTestSuite) TestQueueDispatcher() {
	dispatcher := infra.NewQueueDispatcher(s.logger, s.redisClient, TOPIC_NAME)

	assert.Nil(s.T(), dispatcher.CreditCardTransactionCreated(domain.CreditCardTransaction{
		TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
	}))

	assert.Nil(s.T(), dispatcher.CreditCardTransactionEdited(domain.CreditCardTransaction{
		TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
	}))

	assert.Nil(s.T(), dispatcher.CreditCardTransactionDeleted("01H2KDJMHCTVTN0YDY10S5SNWB"))

	assert.Nil(s.T(), dispatcher.CreditCardTransactionGot(domain.CreditCardTransaction{
		TransactionID: "01H2KDJMHCTVTN0YDY10S5SNWB",
	}))

	assert.Nil(s.T(), dispatcher.CreditCardTransactionListed(domain.PaginationCriteria{
		Page:         1,
		ItemsPerPage: 50,
	}))
}

type TestContainer struct {
	testcontainers.Container
	URI string
}

func setupRedis(ctx context.Context, t testing.TB) TestContainer {
	redisContainer, err := redis.RunContainer(
		ctx,
		testcontainers.WithImage(CONTAINER_IMAGE),
		redis.WithLogLevel(redis.LogLevelVerbose),
	)
	if err != nil {
		t.Errorf("Error on create test container: %s", err)
	}

	URI, err := redisContainer.ConnectionString(ctx)
	if err != nil {
		t.Errorf("Error on get test container uri: %s", err)
	}

	return TestContainer{
		Container: redisContainer,
		URI:       URI,
	}
}
