package redis

import (
	"PlayTogether/roomapis/model"
	_redisValueGenerator "PlayTogether/roomapis/utils/redis"
)

// Room Redis Model
type RoomRedis struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Manager string   `json:"manager"`
	Members []string `json:"members"`
	Songs   string   `json:"songs"`
}

func ConvertRoomToModelRedis(room model.Room) RoomRedis {
	return RoomRedis{
		Id:      room.Id,
		Name:    room.Name,
		Manager: room.Manager,
		Members: room.Members,
		Songs:   _redisValueGenerator.GenPrefixKey("room", room.Id, "songs"),
	}
}
