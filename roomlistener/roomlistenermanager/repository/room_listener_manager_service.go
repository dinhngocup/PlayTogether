package service

import (
	"PlayTogether/roomlistener/model"
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type RoomListenerManagerRepositoryHandler struct {
	roomListenerManager *model.RoomListenerManager
}

func (r *RoomListenerManagerRepositoryHandler) RunAllRoomListeners() {
	fmt.Println("ba")
	fmt.Println(len(r.roomListenerManager.RoomListeners))
	if len(r.roomListenerManager.RoomListeners) > 0 {
		listRoomListeners := make([]*model.RoomListener, len(r.roomListenerManager.RoomListeners))

		for _, value := range r.roomListenerManager.RoomListeners {
			listRoomListeners = append(listRoomListeners, value)
		}
		for _, roomlistener := range listRoomListeners {
			go Run(roomlistener)
		}
	}
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

func (r RoomListenerManagerRepositoryHandler) BroadcastData(roomListenerId string, client *model.Client) {
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
		// find room in room listener and broadcast
		r.roomListenerManager.RoomListeners[roomListenerId].Broadcast <- message
	}
}

func (r *RoomListenerManagerRepositoryHandler) RegisterConnection(roomListenerId string, client *model.Client) {
	// create room listen if not existed
	if r.roomListenerManager.RoomListeners[roomListenerId] == nil {
		roomListener := model.NewListener()
		r.roomListenerManager.RoomListeners[roomListenerId] = roomListener
		go Run(roomListener)
	}
	fmt.Println(r.roomListenerManager.RoomListeners[roomListenerId])
	// register connection
	r.roomListenerManager.RoomListeners[roomListenerId].Register <- client
}

//func (r RoomListenerManagerRepositoryHandler) UnregisterConnection(roomListenerId string, client *model.Client) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (r RoomListenerManagerRepositoryHandler) GetRoomListener(id string) (model.RoomListener, error) {
//	//TODO implement me
//	panic("implement me")
//}

func NewRoomListenerManagerRepository(roomListenerManager *model.RoomListenerManager) model.RoomListenerRepository {
	return &RoomListenerManagerRepositoryHandler{
		roomListenerManager: roomListenerManager,
	}
}

func (r RoomListenerManagerRepositoryHandler) ReadData(client *model.Client) {
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

func Run(h *model.RoomListener) {
	for {
		select {
		case client := <-h.Register:
			fmt.Println("****")
			h.Clients[client] = true
			fmt.Println(h.Clients)
		case message := <-h.Broadcast:
			fmt.Printf("mess %s \n", message)
			for client := range h.Clients {
				fmt.Println("send to ", client)
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
