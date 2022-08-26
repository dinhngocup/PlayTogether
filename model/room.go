package model

// Room Model
type Room struct {
	Id      int32  `json:"id"`
	Name    string `json:"name"`
	Manager User   `json:"manager"`
	Members []User `json:"members"`
}

// RoomService represent the Room's services
type RoomService interface {
	GetByID(id int32) (Room, error)
	AddMember(member *User) error
	RemoveMember(userId string) error
	CreateRoom(room Room) error
}

// RoomRepository represent the Room's repository contract
type RoomRepository interface {
	GetByID(id int32) (Room, error)
	AddMember(member *User) error
	RemoveMember(userId string) error
	CreateRoom(room Room) error
}
