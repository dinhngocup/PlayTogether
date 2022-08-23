package http

import (
	"PlayTogether/model"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

// RoomHandler  represent the http handler for room
type RoomHandler struct {
	roomService model.RoomService
}

func NewRoomDelivery(router *httprouter.Router, roomService model.RoomService) {
	handler := RoomHandler{
		roomService: roomService,
	}
	log.Println("call room apis")

	router.GET("/rooms/:id", handler.GetByID)
}

// GetByID will get room information by given id
func (roomHandler *RoomHandler) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	roomId, _ := strconv.Atoi(ps.ByName("id"))
	fmt.Printf("Room ID: %d\n", roomId)

	roomInfo, err := roomHandler.roomService.GetByID(roomId)

	if err != nil {
		http.Error(w, model.ErrInternalServerError.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(roomInfo)
}
