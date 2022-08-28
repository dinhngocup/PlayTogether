package model

// Room Model
type Room struct {
	Id      string     `json:"id"`
	Name    string     `json:"name"`
	Manager string     `json:"manager"`
	Members []string   `json:"members"`
	Songs   SongInRoom `json:"songs"`
}

// RoomService represent the Room's services
type RoomService interface {
	GetByID(id string) (Room, error)
	AddMember(member *User) error
	RemoveMember(userId string) error
	CreateRoom(room Room) error
}

// RoomRepository represent the Room's repository contract
type RoomRepository interface {
	GetByID(id string) (Room, error)
	AddMember(member *User) error
	RemoveMember(userId string) error
	CreateRoom(room Room) error
}
