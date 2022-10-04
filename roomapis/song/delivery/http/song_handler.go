package http

import (
	"PlayTogether/roomapis/model"
	"PlayTogether/roomapis/model/redis"
	"PlayTogether/roomapis/utils"
	"bytes"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// SongHandler  represent the http handler for song
type SongHandler struct {
	songService model.SongService
	publisher   redis.PublisherService
}

func NewSongDelivery(router *httprouter.Router, songService model.SongService, publisher redis.PublisherService) {
	handler := &SongHandler{
		songService: songService,
		publisher:   publisher,
	}
	log.Println("call song apis")

	router.POST("/songs", handler.AddSong)
	router.DELETE("/songs", handler.RemoveSong)
}

// add song
func (songHandler *SongHandler) AddSong(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()
	addSongRequest := model.AddSongRequest{}
	json.Unmarshal([]byte(body), &addSongRequest)
	err := songHandler.songService.AddSong(addSongRequest)
	if err != nil {
		return
	}
	pubsubPayload := model.SocketData{
		Type:   utils.ROOM,
		Action: utils.BROADCAST,
		UserId: addSongRequest.UserId,
		RoomId: addSongRequest.RoomId,
		Data:   body,
	}
	json, err := json.Marshal(pubsubPayload)
	songHandler.publisher.PublishMessage("mychannel1", string(json))
}

// remove song
func (songHandler *SongHandler) RemoveSong(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
