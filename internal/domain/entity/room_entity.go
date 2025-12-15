package entity

import (
	"time"
)

type RoomStatus string

const (
	RoomStatusWaiting  RoomStatus = "WAITING"
	RoomStatusInPlay   RoomStatus = "IN_PLAY"
	RoomStatusFinished RoomStatus = "FINISHED"
)

type Room struct {
	ID           int
	Code         string
	HostPlayerID int
	Status       RoomStatus
	ConfigJSON   string
	CreatedAt    time.Time
}
