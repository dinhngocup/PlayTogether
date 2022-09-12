package main

import (
	_client "PlayTogether/roomlistener/client"
	_model "PlayTogether/roomlistener/model"
	_room_listener_manager_service "PlayTogether/roomlistener/roomlistenermanager/service"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	roomListenerManager := _model.NewRoomListenerManager()
	roomListenerManagerService := _room_listener_manager_service.NewRoomListenerManagerService(roomListenerManager)

	clientService := _client.NewClientService()

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET params were:", r.URL.Query())
		roomId := r.URL.Query().Get("roomId")
		client := clientService.CreateClient(w, r)
		roomListener := roomListenerManagerService.GetRoomListener(roomId)
		roomListener.RegisterConnection(client)
		go roomListener.OnMessage(client)
		go roomListener.SendDataToClient(client)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
