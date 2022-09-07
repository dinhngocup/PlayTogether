package http

import (
	model2 "PlayTogether/roomapis/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RoomRepositoryHandler struct {
}

func (r *RoomRepositoryHandler) JoinRoom(request model2.JoinRoomRequest) error {
	//TODO implement me
	panic("implement me")
}

func (r *RoomRepositoryHandler) LeaveRoom(request model2.LeaveRoomRequest) error {
	//TODO implement me
	panic("implement me")
}

func NewHttpRoomRepository() model2.RoomRepository {
	return &RoomRepositoryHandler{}
}

func (r *RoomRepositoryHandler) CreateRoom(room model2.Room) error {
	//TODO implement me
	panic("implement me")
}

func (r *RoomRepositoryHandler) GetByID(id string) (model2.Room, error) {
	requestURL := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", id)
	res, err := http.Get(requestURL)
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		//log.Printf(string(bodyBytes))
		roomInfo := model2.Room{}
		json.Unmarshal(bodyBytes, &roomInfo)
		return roomInfo, err
	} else {
		return model2.Room{}, err
	}
}
