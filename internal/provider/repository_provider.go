package provider

import (
	"database/sql"

	"github.com/kroma-labs/poker-ledger-be/internal/domain/repository"
)

type Repositories struct {
	Transactor repository.Transactor
	Player     repository.PlayerRepository
	Room       repository.RoomRepository
}

func provideRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		repository.NewTransactor(db),
		repository.NewPlayerRepository(db),
		repository.NewRoomRepository(db),
	}
}
