package publishredis

import (
	_model "PlayTogether/roomlistener/model/redis"
	"github.com/go-redis/redis"
)

type SubscriberServiceHandler struct {
	redisPubSub *redis.Client
}

func (s SubscriberServiceHandler) SubscribeTopic(chanel string) *redis.PubSub {
	return s.redisPubSub.Subscribe(chanel)
}

func NewSubscriberService(redisClient *redis.Client) _model.SubscriberService {
	return &SubscriberServiceHandler{
		redisPubSub: redisClient,
	}
}
