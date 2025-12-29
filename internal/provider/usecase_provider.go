package provider

import "github.com/kroma-labs/poker-ledger-be/internal/domain/usecase"

type Usecases struct {
	Room usecase.RoomUsecase
}

func provideUsecases(repos *Repositories) *Usecases {
	return &Usecases{
		usecase.NewRoomUsecase(repos.Transactor, repos.Room, repos.Player),
	}
}
