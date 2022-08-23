package model

// Room Model
type Room struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Manager User
	Members []User
}

// RoomService represent the Room's services
type RoomService interface {
	GetByID(id int) (Room, error)
	AddMember(member *User) error
	RemoveMember(userId string) error
}

// RoomRepository represent the Room's repository contract
type RoomRepository interface {
	GetByID(id int) (Room, error)
	AddMember(member *User) error
	RemoveMember(userId string) error
}
