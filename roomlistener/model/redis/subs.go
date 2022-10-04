package redis

import "github.com/go-redis/redis"

type SubscriberService interface {
	SubscribeTopic(chanel string) *redis.PubSub
}
