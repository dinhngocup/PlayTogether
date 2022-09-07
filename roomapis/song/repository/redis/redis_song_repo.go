package redis

import (
	"PlayTogether/roomapis/model"
	"github.com/go-redis/redis"
)

type SongRepositoryHandler struct {
	client *redis.Client
}

func (s *SongRepositoryHandler) AddSong(song model.Song, roomId string) error {
	//TODO implement me
	panic("implement me")
}

func (s *SongRepositoryHandler) RemoveSong(song model.Song, roomId string) error {
	//TODO implement me
	panic("implement me")
}

func NewRedisSongRepository() model.SongRepository {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &SongRepositoryHandler{
		client: redisClient,
	}
}
