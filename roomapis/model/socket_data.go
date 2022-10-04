package model

import "PlayTogether/roomlistener/utils"

type SocketData struct {
	Type         utils.TYPE   `json:"type"`
	Action       utils.ACTION `json:"action"`
	UserId       string       `json:"userId"`
	RoomId       string       `json:"roomId"`
	ConnectionId string       `json:"connectionId"`
	Data         string       `json:"data"`
}
