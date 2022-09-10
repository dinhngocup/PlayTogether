package model

type RoomListenerManager struct {
	// manages list room listens
	RoomListeners map[string]*RoomListener
}

type RoomListenerService interface {
	//GetRoomListener(id string) (RoomListener, error)
	RegisterConnection(roomListenerId string, client *Client)
	//UnregisterConnection(roomListenerId string, client *Client)
	BroadcastData(roomListenerId string, client *Client)
	ReadData(client *Client)
	RunAllRoomListeners()
}

type RoomListenerRepository interface {
	//GetRoomListener(id string) (RoomListener, error)
	RegisterConnection(roomListenerId string, client *Client)
	//UnregisterConnection(roomListenerId string, client *Client)
	BroadcastData(roomListenerId string, client *Client)
	ReadData(client *Client)
	RunAllRoomListeners()
}
