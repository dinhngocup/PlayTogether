package manager

import (
	"PlayTogether/roomlistener/model"
	_subs "PlayTogether/roomlistener/model/redis"
)

// ConnectionManager store map <connectionId,Client>
type ConnectionManager struct {
	ConnectionMap  map[string]*model.Client
	FeatureManager *FeatureManager
}

func NewConnectionManager(featureManager *FeatureManager) *ConnectionManager {
	return &ConnectionManager{
		ConnectionMap:  make(map[string]*model.Client),
		FeatureManager: featureManager,
	}
}

type ConnectionManagerService interface {
	RegisterConnection(client *model.Client) string
	OnMessage(connectionId string, postmanService model.PostmanService, subsRedis _subs.SubscriberService)
	SendToClient(connectionId string)
	SendMessage(connectionIds []string, message string)
}
