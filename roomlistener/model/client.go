package model

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// upgrade http connection to websocket connection (do not close connect)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the room listener.
type Client struct {
	// The websocket connection.
	Connection *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte
}

type ClientService interface {
	CreateClient(w http.ResponseWriter, r *http.Request) *Client
}
