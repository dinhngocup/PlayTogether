package service

import (
	"PlayTogether/roomlistener/model"
)

type RoomListenerManagerServiceHandler struct {
	roomListenerManager *model.RoomListenerManager
}

func (r *RoomListenerManagerServiceHandler) GetRoomListener(roomListenerId string) *model.RoomListener {
	// create room listen if not existed
	if r.roomListenerManager.RoomListeners[roomListenerId] == nil {
		roomListener := model.NewListener()
		r.roomListenerManager.RoomListeners[roomListenerId] = roomListener
		go roomListener.Run()
	}
	return r.roomListenerManager.RoomListeners[roomListenerId]
	//fmt.Println(r.roomListenerManager.RoomListeners[roomListenerId])
	// register connection
	//r.roomListenerManager.RoomListeners[roomListenerId].Register <- client
}

func NewRoomListenerManagerService(roomListenerManager *model.RoomListenerManager) model.RoomListenerManagerService {
	return &RoomListenerManagerServiceHandler{
		roomListenerManager: roomListenerManager,
	}
}
