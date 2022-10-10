package redis

import "github.com/go-redis/redis/v7"

type SubscriberService interface {
	SubscribeTopic(chanel string) *redis.PubSub
}
