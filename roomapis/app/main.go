package main

import (
	_roomHttpDelivery "PlayTogether/roomapis/room/delivery/http"
	_roomRedisRepository "PlayTogether/roomapis/room/repository/redis"
	_roomService "PlayTogether/roomapis/room/service"
	_userHttpDelivery "PlayTogether/roomapis/user/delivery/http"
	_userRedisRepository "PlayTogether/roomapis/user/repository/redis"
	_userService "PlayTogether/roomapis/user/service"

	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()

	roomRedisRepo := _roomRedisRepository.NewRedisRoomRepository()
	roomService := _roomService.NewRoomService(roomRedisRepo)
	_roomHttpDelivery.NewRoomDelivery(router, roomService)

	userRedisRepo := _userRedisRepository.NewRedisUserRepository()
	userService := _userService.NewUserService(userRedisRepo)
	_userHttpDelivery.NewUserDelivery(router, userService)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
