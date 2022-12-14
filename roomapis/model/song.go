package model

// Song Redis Model
type Song struct {
	Id    string `json:"id"`
	Owner string `json:"owner"`
}

// SongService represent the Song's services
type SongService interface {
	AddSong(request AddSongRequest) error
	RemoveSong(song Song, roomId string) error
}

// SongRepository represent the Song's service contract
type SongRepository interface {
	AddSong(request AddSongRequest) error
	RemoveSong(song Song, roomId string) error
}
