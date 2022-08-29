package http

import (
	"PlayTogether/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// RoomHandler  represent the http handler for room
type RoomHandler struct {
	roomService model.RoomService
}

func NewRoomDelivery(router *httprouter.Router, roomService model.RoomService) {
	handler := &RoomHandler{
		roomService: roomService,
	}
	log.Println("call room apis")

	router.GET("/rooms/:id", handler.GetByID)
	router.POST("/rooms", handler.CreateRoom)
	router.POST("/rooms/join", handler.JoinRoom)
	router.POST("/rooms/leave", handler.JoinRoom)
}

// GetByID will get room information by given id
func (roomHandler *RoomHandler) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	roomId := ps.ByName("id")
	fmt.Printf("Room ID: %s\n", roomId)

	roomInfo, err := roomHandler.roomService.GetByID(roomId)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(roomInfo)
}

// CreateRoom will create a new room if it not exists before
func (roomHandler *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// convert data from body request to struct Room by buffer
	// TODO: need to find another way to convert it
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()
	newRoom := model.Room{}
	json.Unmarshal([]byte(body), &newRoom)
	err := roomHandler.roomService.CreateRoom(newRoom)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// JoinRoom will add user into an existed room
func (roomHandler *RoomHandler) JoinRoom(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()
	joinRoomRequest := model.JoinRoomRequest{}
	json.Unmarshal([]byte(body), &joinRoomRequest)
	err := roomHandler.roomService.JoinRoom(joinRoomRequest)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
