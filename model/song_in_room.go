package model

// Song Redis Model
type SongInRoom struct {
	Id    string `json:"id"`
	Owner string `json:"owner"`
}
