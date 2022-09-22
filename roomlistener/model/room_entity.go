package model

type RoomEntity struct {
	// list user id
	Members []string
}

func NewRoomEntity() *RoomEntity {
	return &RoomEntity{
		Members: make([]string, 0),
	}
}
