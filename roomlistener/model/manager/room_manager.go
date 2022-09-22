package manager

import (
	"PlayTogether/roomlistener/model"
	"PlayTogether/roomlistener/utils"
	"fmt"
)

type RoomManager struct {
	// manages list room with schema <roomId, RoomEntity>
	Rooms map[string]*model.RoomEntity
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		Rooms: make(map[string]*model.RoomEntity),
	}
}

func (roomManager *RoomManager) onAction(payload model.SocketData, postmanService model.PostmanService) {
	switch payload.Action {
	case utils.BROADCAST:
		{
			for _, userId := range roomManager.Rooms[payload.RoomId].Members {
				// broadcast message to userId
				fmt.Println(userId)
				postmanService.DeliverMessage(userId, payload.Data)
			}
		}
	case utils.JOIN:
		{
			if roomManager.Rooms[payload.RoomId] == nil {
				roomManager.Rooms[payload.RoomId] = model.NewRoomEntity()
			}
			if !utils.Contains(roomManager.Rooms[payload.RoomId].Members, payload.UserId) {
				roomManager.Rooms[payload.RoomId].Members = append(roomManager.Rooms[payload.RoomId].Members, payload.UserId)
			}
			postmanService.AddConnection(payload.UserId, payload.ConnectionId)
		}

	}
}
