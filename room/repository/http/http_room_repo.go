package http

import (
	"PlayTogether/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RoomRepositoryHandler struct {
}

func (r *RoomRepositoryHandler) JoinRoom(request model.JoinRoomRequest) error {
	//TODO implement me
	panic("implement me")
}

func (r *RoomRepositoryHandler) LeaveRoom(request model.LeaveRoomRequest) error {
	//TODO implement me
	panic("implement me")
}

func NewHttpRoomRepository() model.RoomRepository {
	return &RoomRepositoryHandler{}
}

func (r *RoomRepositoryHandler) CreateRoom(room model.Room) error {
	//TODO implement me
	panic("implement me")
}

func (r *RoomRepositoryHandler) GetByID(id string) (model.Room, error) {
	requestURL := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", id)
	res, err := http.Get(requestURL)
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		//log.Printf(string(bodyBytes))
		roomInfo := model.Room{}
		json.Unmarshal(bodyBytes, &roomInfo)
		return roomInfo, err
	} else {
		return model.Room{}, err
	}
}
