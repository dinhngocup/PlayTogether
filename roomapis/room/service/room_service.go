package service

import (
	"PlayTogether/model"
)

type RoomServiceHandler struct {
	roomRepo model.RoomRepository
}

func (r *RoomServiceHandler) AddSong(song []model.Song, roomId string) error {
	//TODO implement me
	panic("implement me")
}

func (r *RoomServiceHandler) LeaveRoom(request model.LeaveRoomRequest) error {
	//TODO implement me
	panic("implement me")
}

func NewRoomService(roomRepo model.RoomRepository) model.RoomService {
	return &RoomServiceHandler{
		roomRepo: roomRepo,
	}
}

func (r *RoomServiceHandler) JoinRoom(request model.JoinRoomRequest) error {
	return r.roomRepo.JoinRoom(request)
}

func (r *RoomServiceHandler) CreateRoom(room model.Room) error {
	return r.roomRepo.CreateRoom(room)
}

func (r *RoomServiceHandler) GetByID(id string) (model.Room, error) {
	return r.roomRepo.GetByID(id)
}
