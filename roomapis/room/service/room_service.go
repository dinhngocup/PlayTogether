package service

import (
	model2 "PlayTogether/roomapis/model"
)

type RoomServiceHandler struct {
	roomRepo model2.RoomRepository
}

func (r *RoomServiceHandler) AddSong(song []model2.Song, roomId string) error {
	//TODO implement me
	panic("implement me")
}

func (r *RoomServiceHandler) LeaveRoom(request model2.LeaveRoomRequest) error {
	//TODO implement me
	panic("implement me")
}

func NewRoomService(roomRepo model2.RoomRepository) model2.RoomService {
	return &RoomServiceHandler{
		roomRepo: roomRepo,
	}
}

func (r *RoomServiceHandler) JoinRoom(request model2.JoinRoomRequest) error {
	return r.roomRepo.JoinRoom(request)
}

func (r *RoomServiceHandler) CreateRoom(room model2.Room) error {
	return r.roomRepo.CreateRoom(room)
}

func (r *RoomServiceHandler) GetByID(id string) (model2.Room, error) {
	return r.roomRepo.GetByID(id)
}
