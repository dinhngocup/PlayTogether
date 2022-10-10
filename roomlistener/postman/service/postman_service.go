package service

import (
	"PlayTogether/model"
	"PlayTogether/model/manager"
)

type PostmanServiceHandler struct {
	postman *model.Postman
	// should not import connectionManagerService
	// TODO: need to refactor
	connectionManagerService manager.ConnectionManagerService
}

func (p *PostmanServiceHandler) DeliverMessage(userId string, message string) {
	connectionIds := p.postman.UserConnectionMap[userId]
	p.connectionManagerService.SendMessage(connectionIds, message)
}

func (p *PostmanServiceHandler) MapUserConnection(userId string, connectionId string) {
	listConnections := p.postman.UserConnectionMap[userId]
	listConnections = append(listConnections, connectionId)
	p.postman.UserConnectionMap[userId] = listConnections
}

func NewPostmanService(postman *model.Postman, connectionManagerService manager.ConnectionManagerService) model.PostmanService {
	return &PostmanServiceHandler{
		postman:                  postman,
		connectionManagerService: connectionManagerService,
	}
}
