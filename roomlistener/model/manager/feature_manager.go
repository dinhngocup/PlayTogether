package manager

import (
	"PlayTogether/model"
	"PlayTogether/utils"
	"fmt"
)

// FeatureManager store map <connectionId,Client>
type FeatureManager struct {
	roomManager *RoomManager
}

func NewFeatureManager() *FeatureManager {
	return &FeatureManager{
		roomManager: NewRoomManager(),
	}
}

func (featureManager *FeatureManager) BroadcastMessage(payload model.SocketData, postmanService model.PostmanService) {
	fmt.Println(payload.Type)
	switch payload.Type {
	case utils.ROOM:
		featureManager.roomManager.onAction(payload, postmanService)
	}
}
