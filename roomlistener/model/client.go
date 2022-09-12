package model

import (
	"github.com/gorilla/websocket"
	"net/http"
)

// upgrade http connection to websocket connection (do not close connect)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the room listener.
type Client struct {
	// The websocket Connection.
	Connection *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte
}

type ClientService interface {
	CreateClient(w http.ResponseWriter, r *http.Request) *Client
}
