package main

import (
	_client "PlayTogether/roomlistener/client"
	_model "PlayTogether/roomlistener/model"
	_room_listener_manager_repo "PlayTogether/roomlistener/roomlistenermanager/repository"
	_room_listener_manager_service "PlayTogether/roomlistener/roomlistenermanager/service"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	roomListenerManager := &_model.RoomListenerManager{
		RoomListeners: make(map[string]*_model.RoomListener),
	}
	roomListenerManagerRepo := _room_listener_manager_repo.NewRoomListenerManagerRepository(roomListenerManager)
	roomListenerManagerService := _room_listener_manager_service.NewRoomListenerManagerService(roomListenerManagerRepo)

	clientService := _client.NewClientService()

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET params were:", r.URL.Query())
		roomId := r.URL.Query().Get("roomId")
		client := clientService.CreateClient(w, r)
		roomListenerManagerService.RegisterConnection(roomId, client)
		roomListenerManagerService.BroadcastData(roomId, client)
		roomListenerManagerService.ReadData(client)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
