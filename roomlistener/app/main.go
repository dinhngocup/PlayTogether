package main

import (
	_client "PlayTogether/roomlistener/client"
	_connection_manager_service "PlayTogether/roomlistener/connectionmanager/service"
	_model "PlayTogether/roomlistener/model"
	"PlayTogether/roomlistener/model/manager"
	_postman_service "PlayTogether/roomlistener/postman/service"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()

	featureManager := manager.NewFeatureManager()
	connectionManager := manager.NewConnectionManager(featureManager)
	connectionManagerService := _connection_manager_service.NewConnectionManagerService(connectionManager)

	postman := _model.NewPostman()
	postmanService := _postman_service.NewPostmanService(postman, connectionManagerService)

	clientService := _client.NewClientService()

	http.HandleFunc("/establish-connection", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET params were:", r.URL.Query())
		userId := r.URL.Query().Get("userId")
		client := clientService.CreateClient(w, r)
		connectionId := connectionManagerService.RegisterConnection(client)
		fmt.Println("Connection id:", connectionId)

		postmanService.MapUserConnection(userId, connectionId)

		go connectionManagerService.OnMessage(connectionId, postmanService)
		go connectionManagerService.SendToClient(connectionId)

	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}