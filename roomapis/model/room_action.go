package model

type JoinRoomRequest struct {
	UserId string `json:"userId"`
	RoomId string `json:"roomId"`
}

type LeaveRoomRequest struct {
	UserId string `json:"userId"`
	RoomId string `json:"roomId"`
}

type AddSongRequest struct {
	UserId string `json:"userId"`
	RoomId string `json:"roomId"`
	SongId string `json:"songId"`
}
