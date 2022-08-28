package redis

import (
	"PlayTogether/model"
	_roomModelRedis "PlayTogether/model/redis"
	_redisValueGenerator "PlayTogether/utils/redis"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type RoomRepositoryHandler struct {
	client *redis.Client
}

func NewRedisRoomRepository() model.RoomRepository {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return RoomRepositoryHandler{
		client: redisClient,
	}
}

func (r RoomRepositoryHandler) CreateRoom(room model.Room) error {
	roomId := uuid.Must(uuid.NewRandom()).String()
	room.Id = roomId

	// gen key for hashmap
	roomKey := _redisValueGenerator.GenPrefixKey("room", roomId, "")
	songKey := _redisValueGenerator.GenPrefixKey("room", roomId, "songs")
	roomModelRedis := _roomModelRedis.ConvertRoomToModelRedis(room)

	// convert struct to map
	var songMap map[string]interface{}
	inrecSong, _ := json.Marshal(room.Songs)
	json.Unmarshal(inrecSong, &songMap)

	var roomMap map[string]interface{}
	inrecRoom, _ := json.Marshal(roomModelRedis)
	json.Unmarshal(inrecRoom, &roomMap)

	r.client.HMSet(songKey, songMap)
	createRoomResult := r.client.HMSet(roomKey, roomMap).Err()

	if createRoomResult != nil {
		println("create room failed: " + createRoomResult.Error())
		return errors.New("create room failed")
	}
	return createRoomResult
}

func (r RoomRepositoryHandler) GetByID(id string) (model.Room, error) {
	roomKey := _redisValueGenerator.GenPrefixKey("room", id, "")
	mapRoom, _ := r.client.HGetAll(roomKey).Result()

	songKey := _redisValueGenerator.GenPrefixKey("room", id, "songs")
	mapSongs, _ := r.client.HGetAll(songKey).Result()

	roomInfo := ConvertMapToRoom(mapRoom, mapSongs)
	if roomInfo.Id == "" {
		return model.Room{}, errors.New("this room id not exists")
	}
	fmt.Printf("Room info: %s", roomInfo)
	return roomInfo, nil
}

func ConvertMapToRoom(mapRoom map[string]string, mapSongs map[string]string) model.Room {
	fmt.Println(mapSongs)
	fmt.Println(mapRoom)

	return model.Room{
		Id:      mapRoom["id"],
		Name:    mapRoom["name"],
		Manager: mapRoom["manager"],
		Songs: model.SongInRoom{
			Id:    mapSongs["id"],
			Owner: mapSongs["owner"],
		},
	}
}

func (r RoomRepositoryHandler) AddMember(member *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (r RoomRepositoryHandler) RemoveMember(userId string) error {
	//TODO implement me
	panic("implement me")
}
