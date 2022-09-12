package model

type RoomListenerManager struct {
	// manages list room listens
	RoomListeners map[string]*RoomListener
}

func NewRoomListenerManager() *RoomListenerManager {
	return &RoomListenerManager{
		RoomListeners: make(map[string]*RoomListener),
	}
}

type RoomListenerManagerService interface {
	GetRoomListener(roomListenerId string) *RoomListener
}
