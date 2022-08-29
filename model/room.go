package model

// Room Model
type Room struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Manager string   `json:"manager"`
	Members []string `json:"members"`
	Songs   []Song   `json:"songs"`
}

// RoomService represent the Room's services
type RoomService interface {
	GetByID(id string) (Room, error)
	CreateRoom(room Room) error
	JoinRoom(request JoinRoomRequest) error
	LeaveRoom(request LeaveRoomRequest) error
}

// RoomRepository represent the Room's repository contract
type RoomRepository interface {
	GetByID(id string) (Room, error)
	CreateRoom(room Room) error
	JoinRoom(request JoinRoomRequest) error
	LeaveRoom(request LeaveRoomRequest) error
}
