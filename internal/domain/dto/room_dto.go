package dto

type NewRoomRequest struct {
	HostName string `json:"hostName" binding:"required,min=2"`
}

type RoomResponse struct {
	Code       string         `json:"code"`
	Status     string         `json:"status"`
	HostPlayer PlayerResponse `json:"hostPlayer"`
}
