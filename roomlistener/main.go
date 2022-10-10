package main

import (
	_client "PlayTogether/client"
	_connection_manager_service "PlayTogether/connectionmanager/service"
	_model "PlayTogether/model"
	"PlayTogether/model/manager"
	_postman_service "PlayTogether/postman/service"
	_subscriberService "PlayTogether/subscriberredis"
	"github.com/go-redis/redis/v7"

	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8088", "http service address")

func main() {
	flag.Parse()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	subcriberService := _subscriberService.NewSubscriberService(redisClient)

	featureManager := manager.NewFeatureManager()
	connectionManager := manager.NewConnectionManager(featureManager)
	connectionManagerService := _connection_manager_service.NewConnectionManagerService(connectionManager, subcriberService)

	postman := _model.NewPostman()
	postmanService := _postman_service.NewPostmanService(postman, connectionManagerService)

	clientService := _client.NewClientService()
	fmt.Println("hi")
	http.HandleFunc("/establish-connection", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET params were:", r.URL.Query())
		userId := r.URL.Query().Get("userId")
		client := clientService.CreateClient(w, r)
		connectionId := connectionManagerService.RegisterConnection(client)
		fmt.Println("Connection id:", connectionId)

		postmanService.MapUserConnection(userId, connectionId)

		go connectionManagerService.OnMessage(connectionId, postmanService, subcriberService)
		go connectionManagerService.SendToClient(connectionId)

	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
