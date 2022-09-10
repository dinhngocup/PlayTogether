package service

import (
	"PlayTogether/roomlistener/model"
	"fmt"
)

type RoomListenerManagerServiceHandler struct {
	roomListenManagerRepo model.RoomListenerRepository
}

func (r *RoomListenerManagerServiceHandler) RunAllRoomListeners() {
	fmt.Println("hi")
	r.roomListenManagerRepo.RunAllRoomListeners()
}

func (r RoomListenerManagerServiceHandler) ReadData(client *model.Client) {
	r.roomListenManagerRepo.ReadData(client)
}

func (r RoomListenerManagerServiceHandler) BroadcastData(roomListenerId string, client *model.Client) {
	r.roomListenManagerRepo.BroadcastData(roomListenerId, client)
}

func (r RoomListenerManagerServiceHandler) RegisterConnection(roomListenerId string, client *model.Client) {
	r.roomListenManagerRepo.RegisterConnection(roomListenerId, client)
}

//func (r RoomListenerManagerServiceHandler) UnregisterConnection(roomListenerId string, client *model.Client) {
//	r.roomListenManagerRepo.UnregisterConnection(roomListenerId, client)
//}
//
//func (r RoomListenerManagerServiceHandler) GetRoomListener(id string) (model.RoomListener, error) {
//	return r.roomListenManagerRepo.GetRoomListener(id)
//}

func NewRoomListenerManagerService(roomListenManagerRepo model.RoomListenerRepository) model.RoomListenerService {
	return &RoomListenerManagerServiceHandler{
		roomListenManagerRepo: roomListenManagerRepo,
	}
}
