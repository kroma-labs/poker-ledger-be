package entity

import "time"

type Player struct {
	ID        int
	Name      string
	CreatedAt time.Time
}
