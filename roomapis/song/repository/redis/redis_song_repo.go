package redis

import (
	"PlayTogether/roomapis/model"
	_redisValueGenerator "PlayTogether/roomapis/utils/redis"
	"encoding/json"
	"github.com/go-redis/redis"
)

type SongRepositoryHandler struct {
	client *redis.Client
}

func (s *SongRepositoryHandler) AddSong(request model.AddSongRequest) error {
	songsKey := _redisValueGenerator.GenPrefixKey("room", request.RoomId, "songs")

	songKey := _redisValueGenerator.GenPrefixKey("room", request.RoomId, "songs")
	listSongs, _ := s.client.SMembers(songKey).Result()
	newSong := model.Song{
		Id:    request.SongId,
		Owner: request.UserId,
	}
	json, _ := json.Marshal(newSong)
	listSongs = append(listSongs, string(json))
	return s.client.SAdd(songsKey, listSongs).Err()
}

func (s *SongRepositoryHandler) RemoveSong(song model.Song, roomId string) error {
	//TODO implement me
	panic("implement me")
}

func NewRedisSongRepository(redisClient *redis.Client) model.SongRepository {
	return &SongRepositoryHandler{
		client: redisClient,
	}
}
