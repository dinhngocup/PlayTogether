package publishredis

import (
	_model "PlayTogether/roomapis/model/redis"
	"github.com/go-redis/redis"
)

type PublisherServiceHandler struct {
	redisPublisher *redis.Client
}

func (p *PublisherServiceHandler) PublishMessage(chanel string, message string) error {
	return p.redisPublisher.Publish(chanel, message).Err()
}

func NewPublisherService(redisClient *redis.Client) _model.PublisherService {
	return &PublisherServiceHandler{
		redisPublisher: redisClient,
	}
}
