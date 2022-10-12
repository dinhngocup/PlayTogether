package main

import (
	_roomHttpDelivery "PlayTogether/room/delivery/http"
	_roomRedisRepository "PlayTogether/room/repository/redis"
	_roomService "PlayTogether/room/service"

	_songHttpDelivery "PlayTogether/song/delivery/http"
	_songRedisRepository "PlayTogether/song/repository/redis"
	_songService "PlayTogether/song/service"

	_publisherService "PlayTogether/publishredis"
	_userHttpDelivery "PlayTogether/user/delivery/http"
	_userRedisRepository "PlayTogether/user/repository/redis"
	_userService "PlayTogether/user/service"
	"github.com/go-redis/redis"

	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	publisherService := _publisherService.NewPublisherService(redisClient)

	roomRedisRepo := _roomRedisRepository.NewRedisRoomRepository(redisClient)
	roomService := _roomService.NewRoomService(roomRedisRepo)
	_roomHttpDelivery.NewRoomDelivery(router, roomService, publisherService)

	userRedisRepo := _userRedisRepository.NewRedisUserRepository(redisClient)
	userService := _userService.NewUserService(userRedisRepo)
	_userHttpDelivery.NewUserDelivery(router, userService)

	songRedisRepo := _songRedisRepository.NewRedisSongRepository(redisClient)
	songService := _songService.NewSongService(songRedisRepo)
	_songHttpDelivery.NewSongDelivery(router, songService, publisherService)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
