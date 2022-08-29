package service

import (
	"PlayTogether/model"
)

type SongServiceHandler struct {
	songRepo model.SongRepository
}

func (s *SongServiceHandler) AddSong(song model.Song, roomId string) error {
	return s.songRepo.AddSong(song, roomId)
}

func (s *SongServiceHandler) RemoveSong(song model.Song, roomId string) error {
	return s.songRepo.RemoveSong(song, roomId)
}

func NewRoomService(songRepo model.SongRepository) model.SongService {
	return &SongServiceHandler{
		songRepo: songRepo,
	}
}
