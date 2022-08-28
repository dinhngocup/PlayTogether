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
type UserHandler struct {
	userService model.UserService
}

func NewUserDelivery(router *httprouter.Router, userService model.UserService) {
	handler := UserHandler{
		userService: userService,
	}
	log.Println("call user apis")

	router.GET("/users/:id", handler.GetByID)
	router.POST("/users", handler.CreateUser)
}

// GetByID will get room information by given id
func (userHandler *UserHandler) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userId := ps.ByName("id")
	fmt.Printf("User ID: %s\n", userId)

	userInfo, err := userHandler.userService.GetByID(userId)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userInfo)
}

// CreateRoom will create a new room if it not exists before
func (userHandler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()
	newUser := model.User{}
	json.Unmarshal([]byte(body), &newUser)
	err := userHandler.userService.CreateUser(newUser)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
