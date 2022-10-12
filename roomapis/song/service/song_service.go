package service

import (
	"PlayTogether/model"
)

type SongServiceHandler struct {
	songRepo model.SongRepository
}

func (s *SongServiceHandler) AddSong(request model.AddSongRequest) error {
	return s.songRepo.AddSong(request)
}

func (s *SongServiceHandler) RemoveSong(song model.Song, roomId string) error {
	return s.songRepo.RemoveSong(song, roomId)
}

func NewSongService(songRepo model.SongRepository) model.SongService {
	return &SongServiceHandler{
		songRepo: songRepo,
	}
}
