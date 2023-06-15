package main

import (
	"context"
	"fmt"
	"joubertredrat/notification-ms/pkg"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	config, err := pkg.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	redisHost := fmt.Sprintf("%s:%s", config.RedisQueueHost, config.RedisQueuePort)

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisHost,
	})

	subscriber := redisClient.Subscribe(ctx, config.RedisQueueTransactionTopicName)
	fmt.Printf(
		"Subscribe to topic [ %s ] at [ %s ]\n",
		config.RedisQueueTransactionTopicName,
		redisHost,
	)
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf(
			"Topic [ %s ] Message received [ %s ]\n",
			config.RedisQueueTransactionTopicName,
			msg.Payload,
		)
	}
}
