package http

import (
	"PlayTogether/model"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// SongHandler  represent the http handler for song
type SongHandler struct {
	songService model.SongService
}

func NewSongDelivery(router *httprouter.Router, songService model.SongService) {
	handler := &SongHandler{
		songService: songService,
	}
	log.Println("call song apis")

	router.POST("/songs", handler.AddSong)
	router.DELETE("/songs", handler.RemoveSong)
}

// add song
func (songHandler *SongHandler) AddSong(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// remove song
func (songHandler *SongHandler) RemoveSong(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
