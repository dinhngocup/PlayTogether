package service

import (
	"PlayTogether/model"
)

type RoomServiceHandler struct {
	roomRepo model.RoomRepository
}

func NewRoomService(roomRepo model.RoomRepository) model.RoomService {
	return RoomServiceHandler{
		roomRepo: roomRepo,
	}
}

func (r RoomServiceHandler) GetByID(id int) (model.Room, error) {
	return r.roomRepo.GetByID(id)
}

func (r RoomServiceHandler) AddMember(member *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (r RoomServiceHandler) RemoveMember(userId string) error {
	//TODO implement me
	panic("implement me")
}
