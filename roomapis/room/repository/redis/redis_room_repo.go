package redis

import (
	model "PlayTogether/model"
	_roomModelRedis "PlayTogether/model/redis"
	_redisValueGenerator "PlayTogether/utils/redis"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type RoomRepositoryHandler struct {
	client *redis.Client
}

func (r *RoomRepositoryHandler) AddSong(songs []model.Song, roomId string) error {
	existedRoom, err := r.GetByID(roomId)
	if existedRoom.Id == "" {
		return err
	}
	roomKey := _redisValueGenerator.GenPrefixKey("room", roomId, "")
	songsKey := _redisValueGenerator.GenPrefixKey("room", roomId, "songs")
	var listSongs []string

	for _, song := range songs {
		json, _ := json.Marshal(song)
		listSongs = append(listSongs, string(json))
	}

	r.client.SAdd(songsKey, listSongs)
	return r.client.HSet(roomKey, "members", songsKey).Err()
}

func NewRedisRoomRepository() model.RoomRepository {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &RoomRepositoryHandler{
		client: redisClient,
	}
}

func (r *RoomRepositoryHandler) LeaveRoom(request model.LeaveRoomRequest) error {
	existedRoom, err := r.GetByID(request.RoomId)
	if existedRoom.Id == "" {
		return err
	}
	//roomKey := _redisValueGenerator.GenPrefixKey("room", request.RoomId, "")
	membersKey := _redisValueGenerator.GenPrefixKey("room", request.RoomId, "members")
	return r.client.SRem(membersKey, request.UserId).Err()
}

func (r *RoomRepositoryHandler) JoinRoom(request model.JoinRoomRequest) error {
	existedRoom, err := r.GetByID(request.RoomId)
	if existedRoom.Id == "" {
		return err
	}
	roomKey := _redisValueGenerator.GenPrefixKey("room", request.RoomId, "")
	membersKey := _redisValueGenerator.GenPrefixKey("room", request.RoomId, "members")
	r.client.HSet(roomKey, "members", membersKey)
	return r.client.SAdd(membersKey, request.UserId).Err()
}

func (r *RoomRepositoryHandler) CreateRoom(room model.Room) error {
	roomId := uuid.Must(uuid.NewRandom()).String()
	room.Id = roomId

	// gen key for hashmap
	roomKey := _redisValueGenerator.GenPrefixKey("room", roomId, "")
	roomModelRedis := _roomModelRedis.ConvertRoomToModelRedis(room)

	// convert struct to map
	var roomMap map[string]interface{}
	inrecRoom, _ := json.Marshal(roomModelRedis)
	json.Unmarshal(inrecRoom, &roomMap)

	r.AddSong(room.Songs, room.Id)

	createRoomResult := r.client.HMSet(roomKey, roomMap).Err()

	if createRoomResult != nil {
		println("create room failed: " + createRoomResult.Error())
		return errors.New("create room failed")
	}
	return createRoomResult
}

func (r *RoomRepositoryHandler) GetByID(id string) (model.Room, error) {
	roomKey := _redisValueGenerator.GenPrefixKey("room", id, "")
	mapRoom, _ := r.client.HGetAll(roomKey).Result()

	songKey := _redisValueGenerator.GenPrefixKey("room", id, "songs")
	listSongs, _ := r.client.SMembers(songKey).Result()

	membersKey := _redisValueGenerator.GenPrefixKey("room", id, "members")
	listMembers, _ := r.client.SMembers(membersKey).Result()

	roomInfo := ConvertMapToRoom(mapRoom, listSongs, listMembers)
	if roomInfo.Id == "" {
		return model.Room{}, errors.New("this room id not exists")
	}
	return roomInfo, nil
}

func ConvertMapToRoom(mapRoom map[string]string, listSongs []string, listMembers []string) model.Room {
	var songs []model.Song
	for _, song := range listSongs {
		songInfo := model.Song{}
		json.Unmarshal([]byte(song), &songInfo)
		songs = append(songs, songInfo)
	}

	return model.Room{
		Id:      mapRoom["id"],
		Name:    mapRoom["name"],
		Manager: mapRoom["manager"],
		Songs:   songs,
		Members: listMembers,
	}
}
