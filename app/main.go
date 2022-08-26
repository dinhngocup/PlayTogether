package main

import (
	_roomHttpDelivery "PlayTogether/room/delivery/http"
	_roomRedisRepository "PlayTogether/room/repository/redis"
	_roomService "PlayTogether/room/service"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Welcome!\n")
	})

	router.GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
	})
	//roomHttpRepo := _roomHttpRepository.NewHttpRoomRepository()
	roomRedisRepo := _roomRedisRepository.NewRedisRoomRepository()
	roomService := _roomService.NewRoomService(roomRedisRepo)
	_roomHttpDelivery.NewRoomDelivery(router, roomService)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
