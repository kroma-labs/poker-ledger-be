package provider

import (
	"github.com/jmoiron/sqlx"
	"github.com/kroma-labs/poker-ledger-be/internal/domain/repository"
)

type Repositories struct {
	Transactor repository.Transactor
	Player     repository.PlayerRepository
	Room       repository.RoomRepository
}

func provideRepositories(db *sqlx.DB) *Repositories {
	t := repository.NewTransactor(db)
	return &Repositories{
		t,
		repository.NewPlayerRepository(t),
		repository.NewRoomRepository(t),
	}
}
