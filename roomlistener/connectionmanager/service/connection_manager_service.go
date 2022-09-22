package service

import (
	"PlayTogether/roomlistener/model"
	"PlayTogether/roomlistener/model/manager"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type ConnectionManagerServiceHandler struct {
	connectionManager *manager.ConnectionManager
}

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

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func (c *ConnectionManagerServiceHandler) SendToClient(connectionId string) {
	client := c.connectionManager.ConnectionMap[connectionId]
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		client.Connection.Close()
	}()
	for {
		select {
		case message, ok := <-client.Send:
			fmt.Println("Read data from room")
			client.Connection.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				client.Connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.Connection.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(client.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-client.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			client.Connection.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.Connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *ConnectionManagerServiceHandler) SendMessage(connectionIds []string, message string) {
	for _, connectionId := range connectionIds {
		client := c.connectionManager.ConnectionMap[connectionId]
		client.Send <- []byte(message)
	}
}

func (c *ConnectionManagerServiceHandler) OnMessage(connectionId string, postmanService model.PostmanService) {
	client := c.connectionManager.ConnectionMap[connectionId]
	client.Connection.SetReadLimit(maxMessageSize)
	client.Connection.SetReadDeadline(time.Now().Add(pongWait))
	client.Connection.SetPongHandler(func(string) error { client.Connection.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		log.Println("Send data to room listener")
		_, message, err := client.Connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		payload := model.SocketData{}
		json.Unmarshal(message, &payload)
		c.connectionManager.FeatureManager.BroadcastMessage(payload, postmanService)
	}
}

func (c *ConnectionManagerServiceHandler) RegisterConnection(client *model.Client) string {
	connectionId := uuid.Must(uuid.NewRandom()).String()
	c.connectionManager.ConnectionMap[connectionId] = client

	return connectionId
}

func NewConnectionManagerService(connectionManager *manager.ConnectionManager) manager.ConnectionManagerService {
	return &ConnectionManagerServiceHandler{
		connectionManager: connectionManager,
	}
}
