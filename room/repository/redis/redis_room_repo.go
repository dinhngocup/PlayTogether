package redis

import (
	"PlayTogether/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
)

type RoomRepositoryHandler struct {
	client *redis.Client
}

func NewRedisRoomRepository() model.RoomRepository {
	redisclient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return RoomRepositoryHandler{
		client: redisclient,
	}
}

func (r RoomRepositoryHandler) CreateRoom(room model.Room) error {
	json, errJson := json.Marshal(room)
	if errJson != nil {
		fmt.Println(errJson)
	}
	// check room is existed
	_, err := r.GetByID(room.Id)
	if err == nil {
		fmt.Println("This room is existed.")
		return errors.New("this room is existed")
	}

	result := r.client.Set(string(room.Id), json, 0).Err()
	if result != nil {
		fmt.Println(err)
	}

	return result
}

func (r RoomRepositoryHandler) GetByID(id int32) (model.Room, error) {
	result, _ := r.client.Get(string(id)).Result()
	roomInfo := model.Room{}
	err := json.Unmarshal([]byte(result), &roomInfo)
	if err != nil {
		fmt.Println("This room id not exists")
		return model.Room{}, err
	}
	fmt.Printf("Room info: %s", result)
	return roomInfo, nil
}

func (r RoomRepositoryHandler) AddMember(member *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (r RoomRepositoryHandler) RemoveMember(userId string) error {
	//TODO implement me
	panic("implement me")
}
