package entity

import (
	"database/sql"
	"time"
)

type RoomStatus string

const (
	RoomStatusWaiting  RoomStatus = "WAITING"
	RoomStatusInPlay   RoomStatus = "IN_PLAY"
	RoomStatusFinished RoomStatus = "FINISHED"
)

type Room struct {
	ID           int            `db:"id"`
	Code         string         `db:"code"`
	HostPlayerID int            `db:"host_player_id"`
	Status       RoomStatus     `db:"status"`
	ConfigJSON   sql.NullString `db:"config_json"`
	CreatedAt    time.Time      `db:"created_at"`
}
