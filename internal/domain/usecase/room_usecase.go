package usecase

import (
	"context"
	"time"

	"github.com/kroma-labs/poker-ledger-be/internal/domain/dto"
	"github.com/kroma-labs/poker-ledger-be/internal/domain/entity"
	"github.com/kroma-labs/poker-ledger-be/internal/domain/repository"
	"github.com/kroma-labs/poker-ledger-be/internal/pkg/util"
)

type RoomUsecase interface {
	Create(ctx context.Context, request dto.NewRoomRequest) (dto.RoomResponse, error)
}

type roomUsecase struct {
	transactor repository.Transactor
	roomRepo   repository.RoomRepository
	playerRepo repository.PlayerRepository
}

func NewRoomUsecase(
	transactor repository.Transactor,
	roomRepo repository.RoomRepository,
	playerRepo repository.PlayerRepository,
) *roomUsecase {
	return &roomUsecase{
		transactor,
		roomRepo,
		playerRepo,
	}
}

func (ru *roomUsecase) Create(ctx context.Context, request dto.NewRoomRequest) (dto.RoomResponse, error) {
	var response dto.RoomResponse
	err := ru.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		newHost := entity.Player{
			Name:      request.HostName,
			CreatedAt: time.Now(),
		}

		host, err := ru.playerRepo.Insert(ctx, newHost)
		if err != nil {
			return err
		}

		code, err := util.GenerateRandomString(6)
		if err != nil {
			return err
		}

		newRoom := entity.Room{
			Code:         code,
			HostPlayerID: host.ID,
			Status:       entity.RoomStatusWaiting,
			CreatedAt:    time.Now(),
		}

		createdRoom, err := ru.roomRepo.Insert(ctx, newRoom)
		if err != nil {
			return err
		}

		response = dto.RoomResponse{
			Code:   createdRoom.Code,
			Status: string(createdRoom.Status),
			HostPlayer: dto.PlayerResponse{
				ID:   host.ID,
				Name: host.Name,
			},
		}

		return nil
	})
	return response, err
}
